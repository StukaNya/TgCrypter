package httpserver

// HTTP config
// TODO: logger

type Address struct {
	BindAddr string
}

type Logger struct {
	LogLevel string
}

type Config struct {
	Address Address
	Logger  Logger
}

// Return config instance
func NewConfig() *Config {
	return &Config{
		Address: Address{
			BindAddr: ":8080",
		},
		Logger: Logger{
			LogLevel: "debug",
		},
	}
}
