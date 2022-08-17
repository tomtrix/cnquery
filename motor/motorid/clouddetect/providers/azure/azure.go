package azure

import (
	"strings"

	"go.mondoo.io/mondoo/motor/providers/os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"go.mondoo.io/mondoo/lumi/resources/smbios"
	"go.mondoo.io/mondoo/motor/motorid/azcompute"
	"go.mondoo.io/mondoo/motor/platform"
)

const (
	azureIdentifierFileLinux = "/sys/class/dmi/id/sys_vendor"
)

func Detect(provider os.OperatingSystemProvider, pf *platform.Platform) string {
	sysVendor := ""
	if pf.IsFamily("linux") {
		// Fetching the product version from the smbios manager is slow
		// because it iterates through files we don't need to check. This
		// is an optimzation for our sshfs. Also, be aware that on linux,
		// you may not have access to all the smbios things under /sys, so
		// you want to make sure to only check the product_version file
		content, err := afero.ReadFile(provider.FS(), azureIdentifierFileLinux)
		if err != nil {
			log.Debug().Err(err).Msgf("unable to read %s", azureIdentifierFileLinux)
			return ""
		}
		sysVendor = string(content)
	} else {
		mgr, err := smbios.ResolveManager(provider, pf)
		if err != nil {
			return ""
		}
		info, err := mgr.Info()
		if err != nil {
			log.Debug().Err(err).Msg("failed to query smbios")
			return ""
		}
		sysVendor = info.SysInfo.Vendor
	}

	if strings.Contains(sysVendor, "Microsoft Corporation") {
		mdsvc, err := azcompute.Resolve(provider, pf)
		if err != nil {
			log.Debug().Err(err).Msg("failed to get metadata resolver")
			return ""
		}
		id, err := mdsvc.InstanceID()
		if err != nil {
			log.Debug().Err(err).
				Str("transport", provider.Kind().String()).
				Strs("platform", pf.GetFamily()).
				Msg("failed to get azure platform id")
			return ""
		}
		return id
	}

	return ""
}
