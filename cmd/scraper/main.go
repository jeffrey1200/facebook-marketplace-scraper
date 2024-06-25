package main

import (
	"log"

	"github.com/jeffrey1200/facebook-marketplace-scraper/config"
	"github.com/jeffrey1200/facebook-marketplace-scraper/internals/browser"
	"github.com/jeffrey1200/facebook-marketplace-scraper/internals/scraper"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	cfg := config.LoadConfig()

	browserInstance, err := browser.NewBrowser(cfg.BrowserConfig, logger)
	if err != nil {
		logger.Fatal("Failed to launch browser", zap.Error(err))

	}
	defer browserInstance.Close()
	scraper.ScrapeMarketplaceIndividualCar(browserInstance, cfg, logger)
	// carData, err := scraper.ScrapeMarketplace(browserInstance, cfg, logger)
	if err != nil {
		logger.Fatal("Failed to scrape marketplace", zap.Error(err))
	}
	// util.SaveScrapedDataToJSON(carData)
	log.Println("finished scraping!")
}
