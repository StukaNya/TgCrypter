package api

// API server config
type ServerConfig struct {
	BindAddr  string
	InfoRoute string
}

// Return config instance
func NewConfig() *ServerConfig {
	return &ServerConfig{
		BindAddr:  ":8080",
		InfoRoute: "/",
	}
}
