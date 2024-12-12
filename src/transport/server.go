package transport

import (
	"context"
	"net/http"

	"github.com/sagikazarmark/slog-shim"
	"github.com/startup_krasnodar_test/src/pkg/config"
)

type Server struct {
	Server *http.Server
	logger *slog.Logger
}

// создание нового сервера
func NewServer(cfg *config.SrvConfig, handler http.Handler, logger *slog.Logger) (*Server, error) {
	server := &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1MB
		ReadTimeout:    cfg.Timeout,
		WriteTimeout:   cfg.Timeout,
	}

	return &Server{
		Server: server,
		logger: logger,
	}, nil
}

func (s *Server) MustRun() {
	fi := "transport.Server.MustRun"

	s.logger.Info(fi + ":" + "starting server...")
	if err := s.Server.ListenAndServe(); err != nil {
		panic(fi + ":" + "cannot run server: " + err.Error())
	}
}

func (s *Server) Stop(ctx context.Context) {
	fi := "transport.Server.Stop"

	if err := s.Server.Shutdown(ctx); err != nil {
		s.logger.Error(fi + ":" + "cannot stop server: " + err.Error())
	}
}
