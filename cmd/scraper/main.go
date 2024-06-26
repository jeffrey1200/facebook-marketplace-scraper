package main

import (
	"log"

	"github.com/jeffrey1200/facebook-marketplace-scraper/config"
	"github.com/jeffrey1200/facebook-marketplace-scraper/internals/browser"
	"github.com/jeffrey1200/facebook-marketplace-scraper/internals/cli"
	"github.com/jeffrey1200/facebook-marketplace-scraper/internals/scraper"
	"github.com/jeffrey1200/facebook-marketplace-scraper/util"
	"go.uber.org/zap"
)

func main() {
	// fmt.Printf("does pageUrlInfo has data?:%v", pageUrlInfo)
	// pageUrlInfo["maxPrice"]
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	cfg := config.LoadConfig()

	browserInstance, err := browser.NewBrowser(cfg.BrowserConfig, logger)
	if err != nil {
		logger.Fatal("Failed to launch browser", zap.Error(err))

	}
	defer browserInstance.Close()
	// var pageUrlInfo cli.Repl()

	// completeURL := fmt.Sprintf(baseURl+"/search?minPrice=%d&maxPrice=%d&minMileage=%d&maxMileage=%d&daysSinceListed=%d&query=%s&exact=false", pageUrlInfo.minPrice, pageUrlInfo.maxPrice, pageUrlInfo, pageUrlInfo.maxMileage, pageUrlInfo.daysSinceListed, qpageUrlInfo.ueryNameInput)

	pageUrlInfo := cli.Repl()
	// scraper.ScrapeMarketplaceIndividualCar(browserInstance, cfg, logger)
	carData, err := scraper.ScrapeMarketplace(browserInstance, cfg, pageUrlInfo, logger)
	if err != nil {
		logger.Fatal("Failed to scrape marketplace", zap.Error(err))
	}
	err = util.SaveScrapedDataToJSON(carData, pageUrlInfo)
	if err != nil {
		logger.Fatal("Problems with saveScrapeDataToJSON", zap.Error(err))
	}
	log.Println("finished scraping!")
}
