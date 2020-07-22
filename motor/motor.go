package motor

import (
	"github.com/rs/zerolog/log"
	"go.mondoo.io/mondoo/motor/platform"
	"go.mondoo.io/mondoo/motor/transports"
	"go.mondoo.io/mondoo/motor/transports/events"
	"go.mondoo.io/mondoo/motor/transports/local"
	"go.mondoo.io/mondoo/motor/transports/mock"
)

func New(trans transports.Transport) (*Motor, error) {
	c := &Motor{
		Transport: trans,
		detector:  &platform.Detector{Transport: trans},
	}
	return c, nil
}

type Motor struct {
	Transport transports.Transport
	detector  *platform.Detector
	watcher   transports.Watcher
	Meta      MetaInfo
	recording bool
}

type MetaInfo struct {
	Name       string
	Identifier []string
	Labels     map[string]string
}

func (m *Motor) Platform() (*platform.Platform, error) {
	return m.detector.Platform()
}

func (m *Motor) Watcher() transports.Watcher {
	// create watcher once
	if m.watcher == nil {
		m.watcher = events.NewWatcher(m.Transport)
	}
	return m.watcher
}

func (m *Motor) ActivateRecorder() {
	if m.recording {
		return
	}

	mockT, _ := mock.NewRecordTransport(m.Transport)
	m.Transport = mockT
	m.recording = true
}

func (m *Motor) IsRecording() bool {
	return m.recording
}

// returns marshaled toml stucture
func (m *Motor) Recording() []byte {

	if m.recording {
		rt := m.Transport.(*mock.RecordTransport)
		data, err := rt.ExportData()
		if err != nil {
			log.Error().Err(err).Msg("could not export data")
			return nil
		}
		return data
	}
	return nil
}

func (m *Motor) Close() {
	if m == nil {
		return
	}
	if m.Transport != nil {
		m.Transport.Close()
	}
}

func (m *Motor) HasCapability(capability transports.Capability) bool {
	list := m.Transport.Capabilities()
	for i := range list {
		if list[i] == capability {
			return true
		}
	}
	return false
}

func (m *Motor) IsLocalTransport() bool {
	_, ok := m.Transport.(*local.LocalTransport)
	if !ok {
		return false
	}
	return true
}
