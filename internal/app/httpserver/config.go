package httpserver

// HTTP config
// TODO: logger
type Config struct {
	BindAddr string `toml:"bind_addr"`
}

// Return config instance
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
