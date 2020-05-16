package httpserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type httpServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// Return server instanse
func New(config *Config) *httpServer {
	return &httpServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Startup func
func (s *httpServer) Start() error {

	err := s.configureLogger()
	if err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("Server startup")

	return http.ListenAndServe(s.config.Address.BindAddr, s.router)
}

// Configure logger level
func (s *httpServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.Logger.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

// Configure router
func (s *httpServer) configureRouter() {
	s.router.HandleFunc("/info", s.handleInfo())
}

func (s *httpServer) handleInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "DB Info:")
	}
}
