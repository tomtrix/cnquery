// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: BUSL-1.1

package device

import (
	"errors"
	"runtime"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"go.mondoo.com/cnquery/v11/providers-sdk/v1/inventory"
	"go.mondoo.com/cnquery/v11/providers-sdk/v1/plugin"
	"go.mondoo.com/cnquery/v11/providers/os/connection/device/linux"
	"go.mondoo.com/cnquery/v11/providers/os/connection/device/windows"
	"go.mondoo.com/cnquery/v11/providers/os/connection/snapshot"
	"go.mondoo.com/cnquery/v11/utils/stringx"

	"go.mondoo.com/cnquery/v11/providers/os/connection/fs"
	"go.mondoo.com/cnquery/v11/providers/os/connection/shared"
	"go.mondoo.com/cnquery/v11/providers/os/detector"
	"go.mondoo.com/cnquery/v11/providers/os/id"
	"go.mondoo.com/cnquery/v11/providers/os/id/ids"
)

const PlatformIdInject = "inject-platform-ids"

type DeviceConnection struct {
	*fs.FileSystemConnection
	plugin.Connection
	asset         *inventory.Asset
	deviceManager DeviceManager

	MountedDirs []string
	// map of mountpoints to partition infos
	partitions map[string]*snapshot.PartitionInfo
}

func getDeviceManager(conf *inventory.Config) (DeviceManager, error) {
	shell := []string{"sh", "-c"}
	if runtime.GOOS == "darwin" {
		return nil, errors.New("device manager not implemented for darwin")
	}
	if runtime.GOOS == "windows" {
		shell = []string{"powershell", "-c"}
		return windows.NewWindowsDeviceManager(shell, conf.Options)
	}
	return linux.NewLinuxDeviceManager(shell, conf.Options)
}

func NewDeviceConnection(connId uint32, conf *inventory.Config, asset *inventory.Asset) (*DeviceConnection, error) {
	manager, err := getDeviceManager(conf)
	if err != nil {
		return nil, err
	}
	log.Debug().Str("manager", manager.Name()).Msg("device manager created")

	blocks, err := manager.IdentifyMountTargets(conf.Options)
	if err != nil {
		return nil, err
	}
	if len(blocks) == 0 {
		return nil, errors.New("device connection> internal: blocks found")
	}

	res := &DeviceConnection{
		Connection:    plugin.NewConnection(connId, asset),
		deviceManager: manager,
		asset:         asset,
	}

	if conf.Options == nil {
		conf.Options = make(map[string]string)
	}

	if len(asset.IdDetector) == 0 {
		asset.IdDetector = []string{ids.IdDetector_Hostname, ids.IdDetector_SshHostkey}
	}

	res.partitions = make(map[string]*snapshot.PartitionInfo)

	// we iterate over all the blocks and try to run OS detection on each one of them
	// we only return one asset, if we find the right block (e.g. the one with the root FS)
	for _, block := range blocks {
		fsConn, scanDir, err := tryDetectAsset(connId, block, manager, conf, asset)
		if scanDir != "" {
			res.MountedDirs = append(res.MountedDirs, scanDir)
			res.partitions[scanDir] = block
		}
		if err != nil {
			log.Error().Err(err).Msg("partition did not return an asset, continuing")
		} else {
			res.FileSystemConnection = fsConn
		}
	}

	// if none of the blocks returned a platform that we could detect, we return an error
	if asset.Platform == nil {
		res.Close()
		return nil, errors.New("device connection> no platform detected")
	}

	// allow injecting platform ids into the device connection. we cannot always know the asset that's being scanned, e.g.
	// if we can scan an azure VM's disk we should be able to inject the platform ids of the VM
	if platformIDs, ok := conf.Options[PlatformIdInject]; ok {
		platformIds := strings.Split(platformIDs, ",")
		for _, id := range platformIds {
			if !stringx.Contains(asset.PlatformIds, id) {
				log.Debug().Str("platform-id", id).Msg("device connection> injecting platform id")
				asset.PlatformIds = append(asset.PlatformIds, id)
			}
		}
	}

	return res, nil
}

func (c *DeviceConnection) Close() {
	log.Debug().Msg("closing device connection")
	if c == nil {
		return
	}

	if c.deviceManager != nil {
		c.deviceManager.UnmountAndClose()
	}
}

func (p *DeviceConnection) Name() string {
	return "device"
}

func (p *DeviceConnection) Type() shared.ConnectionType {
	return shared.Type_Device
}

func (p *DeviceConnection) Asset() *inventory.Asset {
	return p.asset
}

func (p *DeviceConnection) UpdateAsset(asset *inventory.Asset) {
	p.asset = asset
}

func (p *DeviceConnection) Capabilities() shared.Capabilities {
	return p.FileSystemConnection.Capabilities()
}

func (p *DeviceConnection) RunCommand(command string) (*shared.Command, error) {
	return nil, plugin.ErrRunCommandNotImplemented
}

func (p *DeviceConnection) FileSystem() afero.Fs {
	return p.FileSystemConnection.FileSystem()
}

func (p *DeviceConnection) FileInfo(path string) (shared.FileInfoDetails, error) {
	return p.FileSystemConnection.FileInfo(path)
}

func (p *DeviceConnection) Conf() *inventory.Config {
	return p.FileSystemConnection.Conf
}

func (p *DeviceConnection) Partitions() map[string]*snapshot.PartitionInfo {
	if p.partitions == nil {
		p.partitions = make(map[string]*snapshot.PartitionInfo)
	}

	return p.partitions
}

// tryDetectAsset tries to detect the OS on a given block device
func tryDetectAsset(connId uint32, partition *snapshot.PartitionInfo, manager DeviceManager, conf *inventory.Config, asset *inventory.Asset) (*fs.FileSystemConnection, string, error) {
	log.Debug().Str("name", partition.Name).Str("type", partition.FsType).Msg("mounting partition")
	scanDir, err := manager.Mount(partition)
	if err != nil {
		log.Error().Err(err).Msg("unable to complete mount step")
		return nil, "", err
	}

	// create and initialize fs provider
	conf.Options["path"] = scanDir
	fsConn, err := fs.NewConnection(connId, &inventory.Config{
		Path:       scanDir,
		PlatformId: conf.PlatformId,
		Options:    conf.Options,
		Type:       "fs",
		Record:     conf.Record,
	}, asset)
	if err != nil {
		return nil, scanDir, err
	}

	p, ok := detector.DetectOS(fsConn)
	if !ok {
		log.Debug().
			Str("partition", partition.Name).
			Msg("device connection> cannot detect os")
		return nil, scanDir, errors.New("cannot detect os")
	}

	fingerprint, p, err := id.IdentifyPlatform(fsConn, &plugin.ConnectReq{}, p, asset.IdDetector)
	if err != nil {
		log.Debug().Err(err).Msg("device connection> failed to identify platform from device")
		return nil, scanDir, err
	}
	log.Debug().Str("scan_dir", scanDir).Msg("device connection> detected platform from device")
	asset.Platform = p
	if asset.Name == "" {
		asset.Name = fingerprint.Name
	}
	asset.PlatformIds = append(asset.PlatformIds, fingerprint.PlatformIDs...)
	asset.IdDetector = fingerprint.ActiveIdDetectors
	asset.Platform = p
	asset.Id = conf.Type

	return fsConn, scanDir, nil
}
