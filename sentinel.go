package sentinel

import (
	"time"

	"github.com/dlbarduzzi/sentinel/apis"
	"github.com/dlbarduzzi/sentinel/core"
)

// Ensures that the Sentinel implements the App interface.
var _ core.App = (*Sentinel)(nil)

type Sentinel struct {
	core.App
}

// Config is the Sentinel initialization config struct.
type Config struct{}

func New() *Sentinel {
	return NewWithConfig(Config{})
}

func NewWithConfig(_ Config) *Sentinel {
	s := &Sentinel{}
	s.App = core.NewBaseApp(core.BaseAppConfig{})
	return s
}

func (s *Sentinel) Start() error {
	if err := s.Bootstrap(); err != nil {
		return err
	}

	return apis.Serve(s.App, apis.ServeConfig{
		Port:         9000,
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	})
}
