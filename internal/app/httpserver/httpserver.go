package httpserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// HTTP server object
type httpServer struct {
	config *ServerConfig
	logger *logrus.Logger
	router *mux.Router
}

// Return server instanse
func New(log *logrus.Logger, config *ServerConfig) *httpServer {
	return &httpServer{
		config: config,
		logger: log,
		router: mux.NewRouter(),
	}
}

// Startup func
func (s *httpServer) Start() error {

	s.configureRouter()

	s.logger.Info("Server is listening, URL: ", s.config.BindAddr)

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// Configure router
func (s *httpServer) configureRouter() {
	s.router.HandleFunc(s.config.InfoRoute, s.handleInfo())
}

// Declare handle function
func (s *httpServer) handleInfo() http.HandlerFunc {
	s.logger.Info("Bind handle of ", s.config.InfoRoute)
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "DB Info:")
	}
}
