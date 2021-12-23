package config

// AppContext holds application level dependencies and config
type AppContext struct {
	*AppConfig
}

// NewContext creates a new context from environment variables
func NewContext() (*AppContext, error) {
	config, err := newAppConfig()
	if err != nil {
		return nil, err
	}
	ctx := AppContext{config}
	return &ctx, nil
}
