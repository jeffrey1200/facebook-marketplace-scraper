package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	// "path/filepath"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/jeffrey1200/facebook-marketplace-scraper/config"
	"github.com/jeffrey1200/facebook-marketplace-scraper/internals/models"
	"github.com/jeffrey1200/facebook-marketplace-scraper/util"
	"go.uber.org/zap"
)

func ScrapeMarketplaceIndividualCar(browser *rod.Browser, cfg config.Config, logger *zap.Logger) {
	// outputPath := filepath.Join("facebook-marketplace-scraper", "output")
	// t := fmt.Sprintf("%s/2024-06-19T13:26:46-04:00_cars.json", outputPath)
	// fmt.Println(t)
	file, err := os.ReadFile("../../output/2024-06-19T13:26:46-04:00_cars.json")
	util.HandleError(err, "Error opening file", logger)

	carData := models.CarData{}

	data := json.Unmarshal(file, &carData)
	util.HandleError(data, "Error unmarshalling json data", logger)

	for _, car := range carData.Cars {
		// fmt.Println(baseURl + car.CarDetailsURL)
		page := browser.MustPage("https://www.facebook.com/" + car.CarDetailsURL)
		// pageURL := page.MustInfo().URL
		// page = page.Timeout(10 * time.Second)
		// if pageURL == "https://www.facebook.com/login/?next=%2Fmarketplace%2F" {
		// 	fmt.Printf("skipped page: %s", pageURL)
		// 	continue
		// }
		// page.MustNavigate(baseURl + car.CarDetailsURL)
		viewportOpts := proto.EmulationSetDeviceMetricsOverride{Width: 1680, Height: 1080, DeviceScaleFactor: 1, Mobile: false}
		err = page.SetViewport(&viewportOpts)
		page.MustWaitLoad()
		if err != nil {
			logger.Fatal("Error setting viewport", zap.Error(err))
		}
		// if "https://www.facebook.com/"+car.CarDetailsURL != pageURL {
		// 	fmt.Printf("the skipped url page :%s", "https://www.facebook.com/"+car.CarDetailsURL)
		// 	continue
		// }
		// x78zum5 xdj266r x1emribx xat24cr x1i64zmx x1y1aw1k x1sxyh0 xwib8y2 xurb0ha
		page = page.Timeout(10 * time.Second)
		signInSignUpCloseButton := page.MustElement("div[aria-label='Close']")
		signInSignUpCloseButton.MustClick()
		fmt.Println("Before looking for elements .x1gslohp")
		infoElements, err := page.Elements(".x1gslohp")
		fmt.Println("Before handling err for elements .x1gslohp")
		if err != nil {
			log.Printf("did not find el, next please. %v", page.MustInfo().URL)
			continue
		}
		fmt.Println("after error handling of elements .x1gslohp")
		// proto.NetworkResourceTypeImage
		// exist, el, err := page.HasR()
		seeMoreDescriptionButton, err := page.ElementR("div[role='button'][tabindex='0']", "See more")
		if err != nil {
			fmt.Println("oops see more button was not found")
			continue
		}
		fmt.Println("after looking for see more button text")
		fmt.Println(seeMoreDescriptionButton.MustText())
		seeMoreDescriptionButton.MustClick()
		var carDetails []models.CarExtendedInformation
		for i, el := range infoElements {
			var carDetail models.CarExtendedInformation
			className := el.MustAttribute("class")
			if *className == "x1gslohp" {
				if i == 2 {
					carDetail.AboutThisVehicle = el.MustText()
				} else if i == 3 {
					carDetail.SellerDescription = el.MustText()
				}
				carDetails = append(carDetails, carDetail)
				fmt.Printf("index:%d, element:%v\n", i, el.MustText())

			}

		}
		fmt.Println("saved info", carDetails)

		// time.Sleep(200000 * time.Second)
		page.MustClose()
	}
}
