package httpserver

import "github.com/sirupsen/logrus"

type httpServer struct {
	config *Config
	logger *logrus.Logger
}

// Return server instanse
func New(config *Config) *httpServer {
	return &httpServer{
		config: config,
		logger: logrus.New(),
	}
}

// Startup func
func (s *httpServer) Start() error {

	err := s.configureLogger()
	if err != nil {
		return err
	}

	s.logger.Info("Server startup")

	return nil
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
