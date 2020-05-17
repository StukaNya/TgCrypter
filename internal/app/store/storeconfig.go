package store

// DataBase config
type StoreConfig struct {
	DatabaseURL string
}

// Return config instance
func NewConfig() *StoreConfig {
	return &StoreConfig{
		DatabaseURL: "",
	}
}
