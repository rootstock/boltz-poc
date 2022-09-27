package http

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"context"

	"github.com/gorilla/handlers"
	"github.com/patogallaiov/boltz-poc/connectors"
	"github.com/patogallaiov/boltz-poc/services"
	"github.com/patogallaiov/boltz-poc/storage"

	log "github.com/sirupsen/logrus"
)

const (
	svcStatusOk          = "ok"
	svcStatusDegraded    = "degraded"
	svcStatusUnreachable = "unreachable"
)

type servicesType struct {
	Db  string `json:"db"`
	Rsk string `json:"rsk"`
}

type healthRes struct {
	Status   string       `json:"status"`
	Services servicesType `json:"services"`
}

type Server struct {
	srv      http.Server
	boltz    connectors.BoltzConnector
	rsk      connectors.RSKConnector
	db       storage.DBConnector
	checkout *services.CheckoutService
	config   *services.ConfigService
	now      func() time.Time
}

func New(boltz connectors.BoltzConnector, rsk connectors.RSKConnector, db storage.DBConnector, checkout *services.CheckoutService, config *services.ConfigService) Server {
	return newServer(boltz, rsk, db, checkout, config, time.Now)
}

func newServer(boltz connectors.BoltzConnector, rsk connectors.RSKConnector, db storage.DBConnector, checkout *services.CheckoutService, config *services.ConfigService, now func() time.Time) Server {
	return Server{
		boltz:    boltz,
		rsk:      rsk,
		db:       db,
		checkout: checkout,
		config:   config,
		now:      now,
	}
}

func (s *Server) Start(port uint) error {
	r := Controller(s)

	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	c := handlers.CORS(credentials, methods, origins, headersOk)(r)

	w := log.StandardLogger().WriterLevel(log.DebugLevel)
	h := handlers.LoggingHandler(w, c)
	defer func(w *io.PipeWriter) {
		_ = w.Close()
	}(w)
	s.srv = http.Server{
		Addr:    ":" + fmt.Sprint(port),
		Handler: h,
	}
	log.Info("server started at localhost:", s.srv.Addr)

	err := s.srv.ListenAndServe()
	if err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Shutdown() {
	log.Info("stopping server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown failed: ", err)
	}
	log.Info("server stopped")
}
