package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/base-media-cloud/base-test/config"
	"github.com/base-media-cloud/base-test/internal/domain/ports"
)

type App struct {
	cfg     *config.Config
	handler ports.Handler
	l       zerolog.Logger
}

func New(cfg *config.Config, h ports.Handler, l zerolog.Logger) *App {
	return &App{
		cfg:     cfg,
		handler: h,
		l:       l,
	}
}

func (a *App) Start() error {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/products", a.handler.Get)

	svr := &http.Server{
		Addr:              a.cfg.Host,
		Handler:           r,
		ReadHeaderTimeout: a.cfg.ReadHeaderTimeout,
	}

	a.l.Info().Str("address", a.cfg.Host).Msg("starting server")

	if err := svr.ListenAndServe(); err != nil {
		return err
	}

	a.l.Info().Msg("server shutting down")

	return nil
}
