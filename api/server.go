package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/babylonchain/staking-indexer/api/handlers"
	"github.com/babylonchain/staking-indexer/api/middlewares"
	"github.com/babylonchain/staking-indexer/api/services"
	"github.com/babylonchain/staking-indexer/config"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Server struct {
	httpServer *http.Server
	handlers   *handlers.Handler
}

func New(
	ctx context.Context, cfg *config.Config, services *services.Services,
) (*Server, error) {
	r := chi.NewRouter()

	logLevel, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatal().Err(err).Msg("error while parsing log level")
	}
	zerolog.SetGlobalLevel(logLevel)

	r.Use(middlewares.CorsMiddleware())
	r.Use(middlewares.SecurityHeadersMiddleware())

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		WriteTimeout: cfg.Server.WriteTimeout,
		ReadTimeout:  cfg.Server.ReadTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
		Handler:      r,
	}

	handlers, err := handlers.New(ctx, cfg, services)
	if err != nil {
		log.Fatal().Err(err).Msg("error while setting up handlers")
	}

	server := &Server{
		httpServer: srv,
		handlers:   handlers,
	}
	server.SetupRoutes(r)
	return server, nil
}

func (a *Server) Start() error {
	log.Info().Msgf("Starting server on %s", a.httpServer.Addr)
	return a.httpServer.ListenAndServe()
}

func RunBackground(ctx context.Context, cfg *config.Config, services *services.Services) error {
	srv, err := New(ctx, cfg, services)
	if err != nil {
		return err
	}

	go srv.Start()
	return nil
}
