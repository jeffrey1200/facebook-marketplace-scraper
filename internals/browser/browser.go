package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/jeffrey1200/facebook-marketplace-scraper/config"
	"go.uber.org/zap"
)

func NewBrowser(config config.BrowserConfig, logger *zap.Logger) (*rod.Browser, error) {
	launcher := launcher.New().Headless(false).Set("disable-notifications")

	// if config.Headless {
	// 	launcher = launcher.Headless(false)
	// }
	// if config.DisableNotifications {
	// 	launcher = launcher.Set("disable-notifications")
	// }
	// if config.WindowSize != "" {
	// 	launcher = launcher.Set("window-size", config.WindowSize)
	// }
	launch, err := launcher.Launch()
	if err != nil {
		logger.Error("Failed to launch browser", zap.Error(err))
		return nil, err
	}
	browser := rod.New().ControlURL(launch)
	// defer browser.Close()

	err = browser.Connect()
	if err != nil {
		return nil, err
	}

	logger.Info("Browser launched",
		zap.Bool("Headless", config.Headless),
		zap.Bool("DisableNotifications", config.DisableNotifications),
		zap.String("WindowSize", config.WindowSize),
		zap.String("ControlURL", launch),
	)

	return browser, nil
}
