package core

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/dlbarduzzi/sentinel/tools/logging"
)

const (
	defaultLogLevel  = "info"
	defaultLogFormat = "json"
)

// BaseAppConfig defines a BaseApp configuration option.
type BaseAppConfig struct {
	LogLevel    string
	LogFormat   string
	LogDisabled bool
}

// Ensures that the BaseApp implements the App interface.
var _ App = (*BaseApp)(nil)

// BaseApp implements core.App and defines the base Sentinel app structure.
type BaseApp struct {
	logger *slog.Logger
	config *BaseAppConfig
}

func NewBaseApp(config BaseAppConfig) *BaseApp {
	app := &BaseApp{
		config: &config,
	}

	if app.config.LogLevel == "" {
		app.config.LogLevel = defaultLogLevel
	}

	if app.config.LogFormat == "" {
		app.config.LogFormat = defaultLogFormat
	}

	return app
}

// Logger returns the default app logger.
func (app *BaseApp) Logger() *slog.Logger {
	if app.logger == nil {
		return slog.Default()
	}
	return app.logger
}

// Bootstrap initializes the application.
func (app *BaseApp) Bootstrap() error {
	if err := app.initLogger(); err != nil {
		return err
	}

	return nil
}

// OnShutdown run jobs before the application shuts down.
func (app *BaseApp) OnShutdown() {
	func() {
		fmt.Println("...running first shutdown func()...")
		time.Sleep(time.Millisecond * 10)
	}()
	func() {
		fmt.Println("...running second shutdown func()...")
		time.Sleep(time.Millisecond * 10)
	}()
}

func (app *BaseApp) initLogger() error {
	app.logger = logging.NewLoggerWithConfig(logging.Config{
		Level:    logging.LogLevel(app.config.LogLevel),
		Format:   logging.LogFormat(app.config.LogFormat),
		Disabled: app.config.LogDisabled,
	}).With(slog.String("app", "sentinel"))

	if app.logger == nil {
		return errors.New("logger not initialized")
	}

	return nil
}
