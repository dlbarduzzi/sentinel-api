package apis

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dlbarduzzi/sentinel/core"
)

const (
	DefaultServerPort         = 8090
	DefaultServerIdleTimeout  = time.Second * 5
	DefaultServerReadTimeout  = time.Second * 5
	DefaultServerWriteTimeout = time.Second * 5
)

type ServeConfig struct {
	Port         int
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func Serve(app core.App, config ServeConfig) error {
	if config.Port < 1 {
		config.Port = DefaultServerPort
	}

	if config.IdleTimeout < 1 {
		config.IdleTimeout = DefaultServerIdleTimeout
	}

	if config.ReadTimeout < 1 {
		config.ReadTimeout = DefaultServerReadTimeout
	}

	if config.WriteTimeout < 1 {
		config.WriteTimeout = DefaultServerWriteTimeout
	}

	router := newRouter(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		Handler:      router.buildMux(),
		IdleTimeout:  config.IdleTimeout,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}

	shutdownErr := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		inSignal := <-quit
		app.Logger().Info("server received shutdown signal",
			slog.String("signal", inSignal.String()),
		)

		// Finish running jobs before the server shuts down.
		app.OnShutdown()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			shutdownErr <- err
		}

		shutdownErr <- nil
	}()

	app.Logger().Info("server starting", slog.Int("port", config.Port))

	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownErr
	if err != nil {
		return err
	}

	app.Logger().Info("server stopped")

	return nil
}
