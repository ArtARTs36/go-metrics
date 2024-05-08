package metrics

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	server *http.Server
	cfg    Config
}

func NewServer(cfg Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Serve() error {
	slog.Info(fmt.Sprintf("[metrics-server] listen on %s", s.cfg.Server.Addr))

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	hServer := &http.Server{
		Addr:        s.cfg.Server.Addr,
		Handler:     mux,
		ReadTimeout: s.cfg.Server.Timeout,
	}

	s.server = hServer

	return hServer.ListenAndServe()
}

func (s *Server) ShutdownCtx(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.Server.Timeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}

func (s *Server) SilentShutdown() {
	err := s.Shutdown()
	if err != nil {
		slog.
			With(slog.String("err", err.Error())).
			Error("[metrics-server] failed to shutdown")
	}
}
