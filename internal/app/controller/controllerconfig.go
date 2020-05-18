package controller

// Controller settings
type ControllerConfig struct {
	Format     string
	AppList    string
	AppDetails string
}

// Return config instance
func NewConfig() *ControllerConfig {
	return &ControllerConfig{
		Format:     "json",
		AppList:    "/",
		AppDetails: "/",
	}
}
