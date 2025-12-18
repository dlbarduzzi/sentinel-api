package core

import (
	"fmt"
	"log/slog"
	"time"
)

// BaseAppConfig defines a BaseApp configuration option.
type BaseAppConfig struct{}

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
	return nil
}

// IsBootstrapped checks if the application was initialized.
func (app *BaseApp) IsBootstrapped() bool {
	// TODO: add validation...
	return false
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
