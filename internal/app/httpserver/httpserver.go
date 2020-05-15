package httpserver

type httpServer struct {
	config *Config
}

// Return server instanse
func New(config *Config) *httpServer {
	return &httpServer{
		config: config,
	}
}

// Startup func
func (s *httpServer) Start() error {
	return nil
}
