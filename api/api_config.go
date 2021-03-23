package api

// API server config
type ServerConfig struct {
	BindAddr    string `yaml:"bind_addr"`
	SchemaRoute string `yaml:"schema_route"`
}

// Return config instance
func NewConfig() *ServerConfig {
	return &ServerConfig{
		BindAddr:    ":8080",
		SchemaRoute: "/",
	}
}
