package sentinel

import (
	"fmt"
	"os"
	"time"

	"github.com/dlbarduzzi/sentinel/apis"
	"github.com/dlbarduzzi/sentinel/core"
	"github.com/dlbarduzzi/sentinel/tools/registry"
)

// Ensures that the Sentinel implements the App interface.
var _ core.App = (*Sentinel)(nil)

type Sentinel struct {
	core.App

	// Logger configs
	logLevel  string
	logFormat string

	// Server configs.
	serverPort         int
	serverIdleTimeout  time.Duration
	serverReadTimeout  time.Duration
	serverWriteTimeout time.Duration
}

// Config is the Sentinel initialization config struct.
type Config struct {
	// Logger configs
	LogLevel  string
	LogFormat string

	// Server configs.
	ServerPort         int
	ServerIdleTimeout  time.Duration
	ServerReadTimeout  time.Duration
	ServerWriteTimeout time.Duration
}

func New() *Sentinel {
	return NewWithConfig(Config{
		LogLevel:           "info",
		LogFormat:          "json",
		ServerPort:         8091,
		ServerIdleTimeout:  5,
		ServerReadTimeout:  5,
		ServerWriteTimeout: 5,
	})
}

func NewWithConfig(config Config) *Sentinel {
	s := &Sentinel{
		logLevel:           config.LogLevel,
		logFormat:          config.LogFormat,
		serverPort:         config.ServerPort,
		serverIdleTimeout:  config.ServerIdleTimeout,
		serverReadTimeout:  config.ServerReadTimeout,
		serverWriteTimeout: config.ServerWriteTimeout,
	}

	s.parseConfig(&config)

	s.App = core.NewBaseApp(core.BaseAppConfig{
		LogLevel:  s.logLevel,
		LogFormat: s.logFormat,
	})

	return s
}

func (s *Sentinel) Start() error {
	if err := s.Bootstrap(); err != nil {
		return err
	}

	return apis.Serve(s.App, apis.ServeConfig{
		Port:         s.serverPort,
		IdleTimeout:  time.Second * s.serverIdleTimeout,
		ReadTimeout:  time.Second * s.serverReadTimeout,
		WriteTimeout: time.Second * s.serverWriteTimeout,
	})
}

func (s *Sentinel) parseConfig(config *Config) {
	// Set logger config defaults.
	s.logLevel = config.LogLevel
	s.logFormat = config.LogFormat

	// Set server config defaults.
	s.serverPort = config.ServerPort
	s.serverIdleTimeout = config.ServerIdleTimeout
	s.serverReadTimeout = config.ServerReadTimeout
	s.serverWriteTimeout = config.ServerWriteTimeout

	r, err := registry.NewWithConfig(registry.Config{
		EnvPrefix: "SENTINEL",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[error] fail to load registry - %s\n", err)
		return
	}

	// Read logger env variables.
	s.logLevel = r.GetString("LOG_LEVEL")
	s.logFormat = r.GetString("LOG_FORMAT")

	// Read server env variables.
	s.serverPort = r.GetInt("SERVER_PORT")
	s.serverIdleTimeout = r.GetDuration("SERVER_IDLE_TIMEOUT_SECS")
	s.serverReadTimeout = r.GetDuration("SERVER_READ_TIMEOUT_SECS")
	s.serverWriteTimeout = r.GetDuration("SERVER_WRITE_TIMEOUT_SECS")
}
