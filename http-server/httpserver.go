package httpserver

import (
	"io"
	"net/http"
	"strconv"

	"github.com/StukaNya/SteamREST/internal/app/controller"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// HTTP server object
type httpServer struct {
	config *ServerConfig
	logger *logrus.Logger
	ctrl   *controller.Controller
	router *mux.Router
}

// Return server instanse
func New(log *logrus.Logger, config *ServerConfig, ctrl *controller.Controller) *httpServer {
	return &httpServer{
		config: config,
		logger: log,
		ctrl:   ctrl,
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
		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}
		// Get id from mux router
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		// Call controller func to get app data from DB
		info, err := s.ctrl.AppInfo(id)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		io.WriteString(w, info)
	}
}
