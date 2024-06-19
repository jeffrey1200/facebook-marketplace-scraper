package config

type BrowserConfig struct {
	Headless             bool
	DisableNotifications bool
	WindowSize           string
}

type Config struct {
	BrowserConfig BrowserConfig
}

func LoadConfig() Config {

	return Config{
		BrowserConfig: BrowserConfig{
			Headless:             false,
			DisableNotifications: true,
			WindowSize:           "1920,1080",
		},
	}
}
