package scraper

import (
	"fmt"
	// "log"
	"strconv"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/proto"
	"github.com/jeffrey1200/facebook-marketplace-scraper/config"
	"github.com/jeffrey1200/facebook-marketplace-scraper/internals/cli"
	"github.com/jeffrey1200/facebook-marketplace-scraper/internals/models"
	"github.com/jeffrey1200/facebook-marketplace-scraper/util"
	"go.uber.org/zap"
)

func ScrapeMarketplace(browser *rod.Browser, cfg config.Config, searchParameters cli.UrlParameters, logger *zap.Logger) (models.CarData, error) {

	selectLocationClasses := "x1i10hfl x1qjc9v5 xjbqb8w xjqpnuy xa49m3k xqeqjp1 x2hbi6w x13fuv20 xu3j5b3 x1q0q8m5 x26u7qi x972fbf xcfux6l x1qhh985 xm0m39n x9f619 x1ypdohk xdl72j9 x2lah0s xe8uvvx x11i5rnm xat24cr x1mh8g0r x2lwn1j xeuugli xexx8yu x4uap5 x18d9i69 xkhd6sd x1n2onr6 x16tdsg8 x1hl2dhg xggy1nq x1ja2u2z x1t137rt x1o1ewxj x3x9cwd x1e5q0jg x13rtm0m x1q0g3np x87ps6o x1lku1pv x78zum5 x1a2a7pz x1xmf6yo"
	selectLocationJoinedClasses := fmt.Sprintf("div.%s", util.JoinClassNames(selectLocationClasses))
	closeButtonClasses := "x1i10hfl x1ejq31n xd10rxx x1sy0etr x17r0tee x1ypdohk xe8uvvx xdj266r x11i5rnm xat24cr x1mh8g0r x16tdsg8 x1hl2dhg xggy1nq x87ps6o x1lku1pv x1a2a7pz x6s0dn4 xzolkzo x12go9s9 x1rnf11y xprq8jg x972fbf xcfux6l x1qhh985 xm0m39n x9f619 x78zum5 xl56j7k xexx8yu x4uap5 x18d9i69 xkhd6sd x1n2onr6 xc9qbxq x14qfxbe x1qhmfi1"
	selectCityFirstOption := fmt.Sprintf("div.%s", util.JoinClassNames("x1i10hfl x1qjc9v5 xjbqb8w xjqpnuy xa49m3k xqeqjp1 x2hbi6w x13fuv20 xu3j5b3 x1q0q8m5 x26u7qi x972fbf xcfux6l x1qhh985 xm0m39n x9f619 x1ypdohk xdl72j9 x2lah0s xe8uvvx xdj266r x11i5rnm xat24cr x1mh8g0r x2lwn1j xeuugli xexx8yu x4uap5 x18d9i69 xkhd6sd x1n2onr6 x16tdsg8 x1hl2dhg xggy1nq x1ja2u2z x1t137rt x1o1ewxj x3x9cwd x1e5q0jg x13rtm0m x1q0g3np x87ps6o x1lku1pv x78zum5 x1a2a7pz xh8yej3"))
	closeButtonJoinedClasses := fmt.Sprintf("div.%s", util.JoinClassNames(closeButtonClasses))
	imgTagClassName := fmt.Sprintf("img.%s", util.JoinClassNames("xt7dq6l xl1xv1r x6ikm8r x10wlt62 xh8yej3"))
	// beforeMileageSpanClassName := fmt.Sprintf("span.%s", util.JoinClassNames("x193iq5w xeuugli x13faqbe x1vvkbs x10flsy6 x1lliihq x1s928wv xhkezso x1gmr53x x1cpjm7i x1fgarty x1943h6x x1tu3fi x3x7a5m x1nxh6w3 x1sibtaa xo1l8bm xi81zsa"))
	// spanMileage := fmt.Sprintf("span.%s", util.JoinClassNames("x1lliihq x6ikm8r x10wlt62 x1n2onr6 xlyipyv xuxw1ft"))
	actualCostSpanClassName := fmt.Sprintf("span.%s", util.JoinClassNames("x193iq5w xeuugli x13faqbe x1vvkbs xlh3980 xvmahel x1n0sxbx x1lliihq x1s928wv xhkezso x1gmr53x x1cpjm7i x1fgarty x1943h6x x4zkp8e x3x7a5m x1lkfr7t x1lbecb7 x1s688f xzsf02u"))
	priceWasAtSpanClassName := fmt.Sprintf("span.%s", util.JoinClassNames("x193iq5w xeuugli x13faqbe x1vvkbs xlh3980 xvmahel x1n0sxbx x1lliihq x1s928wv xhkezso x1gmr53x x1cpjm7i x1fgarty x1943h6x x4zkp8e x3x7a5m x1lkfr7t x1lbecb7 xk50ysn xi81zsa"))

	pageURL := fmt.Sprintf(baseURl+"/search?minPrice=%d&maxPrice=%d&minMileage=%d&maxMileage=%d&daysSinceListed=%d&query=%s&exact=false", searchParameters.MinPrice, searchParameters.MaxPrice, searchParameters.MinMileage, searchParameters.MaxMileage, searchParameters.DaysSinceListed, searchParameters.QueryName)

	page, err := browser.Page(proto.TargetCreateTarget{URL: pageURL})
	if err != nil {
		logger.Fatal("Error creating page", zap.Error(err))
	}

	viewportOpts := proto.EmulationSetDeviceMetricsOverride{Width: 1680, Height: 1080, DeviceScaleFactor: 1, Mobile: false}
	err = page.SetViewport(&viewportOpts)
	if err != nil {
		logger.Fatal("Error setting viewport", zap.Error(err))
	}

	signUpCloseButton, err := page.Element(closeButtonJoinedClasses)
	util.HandleError(err, "Error finding sign up close button", logger)
	util.HandleError(signUpCloseButton.Click(proto.InputMouseButtonLeft, 1), "Error clicking sign up close button", logger)

	locationButton, err := page.Element(selectLocationJoinedClasses)
	util.HandleError(err, "Error finding location text button", logger)
	util.HandleError(locationButton.Click(proto.InputMouseButtonLeft, 1), "Error clicking location text button", logger)

	locationInputBox, err := page.Element("label[aria-label='Location']")
	util.HandleError(err, "Error finding location input box inside location modal", logger)
	util.HandleError(locationInputBox.Click(proto.InputMouseButtonLeft, 1), "Error selecting location input box inside location model", logger)
	util.HandleError(locationInputBox.Input(fmt.Sprintf("%s, %s", searchParameters.City, searchParameters.State)), "Error inputting text into location input box inside location model", logger)

	selectCityRaceHandler := page.Race().Element(selectCityFirstOption).Handle(func(e *rod.Element) error {
		CityAndStateElement, err := page.Element("#\\:r1p\\:")
		util.HandleError(err, "Error finding city and state dropdown", logger)

		isCityAndStateElementVisible, err := CityAndStateElement.Visible()
		util.HandleError(err, "Error finding city and state modal. Not visible", logger)

		if isCityAndStateElementVisible {
			util.HandleError(e.Click(proto.InputMouseButtonLeft, 1), "Error clicking the first location suggested option", logger)
			return nil
		}

		return err
	})
	_, err = selectCityRaceHandler.Do()
	util.HandleError(err, "Error doing the selectCityRaceHandler", logger)

	radiusDropdown, err := page.Element("label[aria-label='Radius']")
	util.HandleError(err, "Error finding select radius in miles button", logger)
	util.HandleError(radiusDropdown.Click(proto.InputMouseButtonLeft, 1), "Error clicking the miles selection dropdown", logger)

	selectMilesDropdownRaceHandler := page.Race().Element("#\\:ra\\:").Handle(func(e *rod.Element) error {
		selectMiles(page, 80, logger)

		return nil
	})
	_, err = selectMilesDropdownRaceHandler.Do()
	util.HandleError(err, "Error doing the selectMilesDropdownRaceHandler", logger)

	locationModalApplyButton, err := page.Element("div[aria-label='Apply']")
	util.HandleError(err, "Error finding apply button", logger)
	util.HandleError(locationModalApplyButton.Click(proto.InputMouseButtonLeft, 1), "Error while clicking the Apply button inside location modal", logger)

	// log.Println(actualCostSpanClassName)
	// time.Sleep(1 * time.Second)

	scrollDuration := 10 * time.Second
	scrollInterval := 300 * time.Millisecond
	scrollEndTime := time.Now().Add(scrollDuration)
	for time.Now().Before(scrollEndTime) {

		util.HandleError(page.Mouse.Scroll(0, 500, 10), "Error while trying to scroll through the page", logger) // Scroll vertically by 500 pixels in 10 steps
		time.Sleep(scrollInterval)
	}

	parentDIV := fmt.Sprintf("div.%s", "xjp7ctv")
	data, err := page.Elements(parentDIV)
	util.HandleError(err, "Error finding all divs with class name: xjp7ctv", logger)

	// fmt.Printf("%s %s", beforeMileageSpanClassName, spanMileage)
	var cars []models.CarInformation
	// time.Sleep(2 * time.Second)
	page.MustWaitIdle()
	for _, divs := range data {
		var car models.CarInformation
		// fmt.Println(divs.Text())
		links, err := divs.Element(fmt.Sprintf("a.%s", util.JoinClassNames("x1i10hfl xjbqb8w x1ejq31n xd10rxx x1sy0etr x17r0tee x972fbf xcfux6l x1qhh985 xm0m39n x9f619 x1ypdohk xt0psk2 xe8uvvx xdj266r x11i5rnm xat24cr x1mh8g0r xexx8yu x4uap5 x18d9i69 xkhd6sd x16tdsg8 x1hl2dhg xggy1nq x1a2a7pz x1heor9g x1sur9pj xkrqix3 x1lku1pv")))
		if err != nil {
			fmt.Printf("error finding <a> tags with item's detailed information urls. Err:%v", err)
			continue
		}
		href, err := links.Attribute("href")
		if err != nil {
			fmt.Printf("error with item's <a> tags href attribute *string.href:%v Err:%v", *href, err)
		}
		car.CarDetailsURL = "www.facebook.com" + *href

		actualCostSpan, err := divs.Element(actualCostSpanClassName)
		if err != nil {
			fmt.Printf("error finding price spans. Err:%v", err)
		}
		actualCostText, err := actualCostSpan.Text()
		if err != nil {
			fmt.Printf("error while getting price span text. Err:%v", err)
		}
		costInt, err := formatAndConvertPriceToInt(actualCostText)
		if err != nil {
			costInt = 0
			fmt.Printf("error while converting price string to int. Err:%v", err)
		}
		car.Price = costInt

		priceWasAtSpan, err := divs.Element(priceWasAtSpanClassName)
		if err == nil {
			// fmt.Printf("error finding price was at spans. Err:%v", err)

			priceWasAtText, err := priceWasAtSpan.Text()
			if err != nil {
				fmt.Printf("error getting price was at price. Err:%v", err)
			}
			priceWasAtInt, err := formatAndConvertPriceToInt(priceWasAtText)
			if err != nil {
				fmt.Printf("error converting price was at price from string to int. Err:%v", err)
			}
			car.PriceWasAt = priceWasAtInt
		}

		imgTagInfo, err := divs.Element(imgTagClassName)
		if err != nil {
			fmt.Printf("error finding item's <img> tag. Err:%v", err)
			continue
		}
		imgSouceURL, err := imgTagInfo.Attribute("src")
		if err != nil {
			fmt.Printf("error getting item's image <src> attribute pointer *string. Err:%v", err)
		}

		car.ImageURL = *imgSouceURL

		mileageElement, err := divs.Elements("div.x1iorvi4.x4uap5.xjkvuk6.xkhd6sd span span.x1lliihq.x6ikm8r.x10wlt62.x1n2onr6.xlyipyv.xuxw1ft")
		if err != nil {
			fmt.Printf("error finding car's mileage <span> tags. Err:%v", err)
		}
		// spanWithText, err := mileageElement
		if len(mileageElement) > 0 {

			location, _ := mileageElement[0].Text()
			mileage, _ := mileageElement[1].Text()
			fmt.Println(location, mileage)
			car.Location = location
			car.Mileage = mileage

		}

		itemNameSpan, err := divs.Element(fmt.Sprintf("span.%s", util.JoinClassNames("x1lliihq x6ikm8r x10wlt62 x1n2onr6")))
		if err != nil {
			fmt.Printf("error finding car's name <span>. Err:%v", err)
		}
		itemNameText, err := itemNameSpan.Text()
		if err != nil {
			fmt.Printf("error getting car's name text. Err:%v", err)
		}
		car.CarName = itemNameText

		cars = append(cars, car)

	}
	currentTime := time.Now()
	formattedTime := fmt.Sprintf("%d-%d-%d-%d:%d:%d",
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Hour(),
		currentTime.Second())

	carData := models.CarData{
		CreationDate:         formattedTime,
		ScrapedUrlParameters: searchParameters,
		AmountOfCars:         len(cars),
		Cars:                 cars,
	}
	// util.SaveScrapedDataToJSON(carData)
	// fmt.Println(cars)

	return carData, nil

}

func formatAndConvertPriceToInt(price string) (int, error) {

	actualCost := strings.Replace(price, "$", "", 1)
	actualCost = strings.Replace(actualCost, ",", "", 1)
	costInt, err := strconv.Atoi(actualCost)
	if err != nil {
		return 0, err
	}
	return costInt, nil

}

var milesValues = map[int]int{
	0:  1,
	1:  2,
	2:  5,
	3:  10,
	4:  20,
	5:  40,
	6:  60,
	7:  80,
	8:  100,
	9:  250,
	10: 500,
}

// type milesValues map[int]string

func selectMiles(page *rod.Page, mile int, logger *zap.Logger) {

	var selectedMileageIndex int
	for i, v := range milesValues {
		if mile == v {
			selectedMileageIndex = i
			break
		}
	}

	for i := 0; i < selectedMileageIndex+1; i++ {
		util.HandleError(page.Keyboard.Press(input.ArrowDown), "Error while trying to press keyboard down arrow", logger)
		// err := page.Keyboard.Press(input.ArrowDown)
		// if err != nil {
		// return err
		// }
		time.Sleep(300 * time.Millisecond)
	}
	util.HandleError(page.Keyboard.Press(input.Enter), "Error while pressing enter to select the desired miles", logger)

}
