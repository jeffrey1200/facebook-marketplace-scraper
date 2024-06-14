package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"go.uber.org/zap"
)

type carInformation struct {
	Price         int    `json:"price"`
	PriceWasAt    int    `json:"price_was_at"`
	CarName       string `json:"car_name"`
	Location      string `json:"location"`
	Mileage       string `json:"mileage"`
	CarDetailsURL string `json:"car_details_url"`
	ImageURL      string `json:"image_url"`
}

func handleError(err error, msg string, logger *zap.Logger) {
	if err != nil {
		logger.Error(msg, zap.Error(err))
	}
}

func retry(attempts int, initialSleep time.Duration, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		if i == 0 {
			time.Sleep(initialSleep)
		}
		err = fn()
		if err == nil {
			return nil
		}

	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

func main() {
	// const (
	// 	email                    = "jeffreymartes@hotmail.com"
	// 	pass                     = "jeffreymartes"
	// 	marketplaceIconClassName = "x1i10hfl x1qjc9v5 xjbqb8w xjqpnuy xa49m3k xqeqjp1 x2hbi6w x13fuv20 xu3j5b3 x1q0q8m5 x26u7qi x972fbf xcfux6l x1qhh985 xm0m39n x9f619 x1ypdohk xdl72j9 x2lah0s xe8uvvx xdj266r x11i5rnm xat24cr x1mh8g0r x2lwn1j xeuugli xexx8yu x4uap5 x18d9i69 xkhd6sd x1n2onr6 x16tdsg8 x1hl2dhg xggy1nq x1ja2u2z x1t137rt x1q0g3np x87ps6o x1lku1pv x1a2a7pz x1lq5wgf xgqcy7u x30kzoy x9jhf4c x1lliihq"
	// )
	selectLocationClasses := "x1i10hfl x1qjc9v5 xjbqb8w xjqpnuy xa49m3k xqeqjp1 x2hbi6w x13fuv20 xu3j5b3 x1q0q8m5 x26u7qi x972fbf xcfux6l x1qhh985 xm0m39n x9f619 x1ypdohk xdl72j9 x2lah0s xe8uvvx x11i5rnm xat24cr x1mh8g0r x2lwn1j xeuugli xexx8yu x4uap5 x18d9i69 xkhd6sd x1n2onr6 x16tdsg8 x1hl2dhg xggy1nq x1ja2u2z x1t137rt x1o1ewxj x3x9cwd x1e5q0jg x13rtm0m x1q0g3np x87ps6o x1lku1pv x78zum5 x1a2a7pz x1xmf6yo"
	closeButtonClasses := "x1i10hfl x1ejq31n xd10rxx x1sy0etr x17r0tee x1ypdohk xe8uvvx xdj266r x11i5rnm xat24cr x1mh8g0r x16tdsg8 x1hl2dhg xggy1nq x87ps6o x1lku1pv x1a2a7pz x6s0dn4 xzolkzo x12go9s9 x1rnf11y xprq8jg x972fbf xcfux6l x1qhh985 xm0m39n x9f619 x78zum5 xl56j7k xexx8yu x4uap5 x18d9i69 xkhd6sd x1n2onr6 xc9qbxq x14qfxbe x1qhmfi1"
	// selectRadiusClasses := "xjhjgkd x1epquy7 xsnmfus x1562eck xcymrrh x1268tai x1mxuytg x14hpm34 xqvykr2 x13fuv20 xu3j5b3 x1q0q8m5 x26u7qi xq2ru2l x17qb25w xjmv2fv x1b4qsv2 x78zum5 xdt5ytf x6ikm8r x10wlt62 x1n2onr6 x1ja2u2z x1egnk41 x1ypdohk x1a2a7pz"
	// https://www.facebook.com/marketplace/108530542504412/search?minPrice=100&maxPrice=10000&daysSinceListed=1&maxMileage=100000&minMileage=50000&query=honda%20civic&exact=false
	// itemsImgAndInformationClassName := "xt7dq6l xl1xv1r x6ikm8r x10wlt62 xh8yej3"
	// itemMileageClassName := "x1lliihq x6ikm8r x10wlt62 x1n2onr6 xlyipyv xuxw1ft"

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	launcher := launcher.New().Headless(false).Set("disable-notifications")
	//log in and sign up frame class name
	//x78zum5 xdt5ytf x2lah0s x193iq5w x2bj2ny x1ey2m1c xayqjjm x9f619 xds687c x1xy6bms xn6708d x1s14bel x1ye3gou xixxii4 x17qophe x1u8a7rm
	url, err := launcher.Launch()
	if err != nil {
		log.Fatalf("Could not launch the browser. err: %v", err)
	}
	browser := rod.New().ControlURL(url)
	defer browser.Close()

	err = browser.Connect()
	if err != nil {
		logger.Fatal("Could not connect to the browser.", zap.Error(err))
	}
	defer browser.Close()
	// // .SlowMotion(2 * time.Second)

	page, err := browser.Page(proto.TargetCreateTarget{URL: "https://www.facebook.com/marketplace/search?minPrice=100&maxPrice=10000&daysSinceListed=1&maxMileage=170000&minMileage=0&query=honda civic&exact=false"})
	if err != nil {
		logger.Fatal("Error creating page", zap.Error(err))
	}

	viewportOpts := proto.EmulationSetDeviceMetricsOverride{Width: 1680, Height: 1080, DeviceScaleFactor: 1, Mobile: false}
	err = page.SetViewport(&viewportOpts)
	if err != nil {
		logger.Fatal("Error setting viewport", zap.Error(err))
	}

	closeButtonJoinedClasses := fmt.Sprintf("div.%s", joinClassNames(closeButtonClasses))
	selectLocationJoinedClasses := fmt.Sprintf("div.%s", joinClassNames(selectLocationClasses))
	// itemIMG := joinClassNames(itemsImgAndInformationClassName)
	// itemMileage := joinClassNames(itemMileageClassName)
	// joinedCloseButtonClasses := fmt.Sprintf("div.%s", closeButton)
	// joinedSelectLocationClasses := fmt.Sprintf("div.%s", selectLocation)
	signUpCloseButton, err := page.Element(closeButtonJoinedClasses)
	handleError(err, "Error finding sign up close button", logger)
	handleError(signUpCloseButton.Click(proto.InputMouseButtonLeft, 1), "Error clicking sign up close button", logger)
	// page.MustElement("div[aria-label='Close']").MustClick()

	locationButton, err := page.Element(selectLocationJoinedClasses)
	// if err != nil {
	// 	log.Fatalf("Could not locate the location button text. err: %v", err)
	// }
	handleError(err, "Error finding location text button", logger)
	handleError(locationButton.Click(proto.InputMouseButtonLeft, 1), "Error clicking location text button", logger)

	time.Sleep(1 * time.Second)

	// joinedRadiusClasses := fmt.Sprintf("div.%s", joinClassNames(selectRadiusClasses))
	locationInputBox, err := page.Element("label[aria-label='Location']")
	handleError(err, "Error finding location input box inside location modal", logger)
	handleError(locationInputBox.Click(proto.InputMouseButtonLeft, 1), "Error selecting location input box inside location model", logger)
	handleError(locationInputBox.Input("Newark, Delaware"), "Error inputting text into location input box inside location model", logger)

	time.Sleep(2 * time.Second)
	handleError(page.Keyboard.Press(input.ArrowDown), "Error pressing down arrow keyboard to select suggested city, state", logger)
	time.Sleep(500 * time.Millisecond)
	handleError(page.Keyboard.Press(input.Enter), "Error while selecting suggested city, state first option", logger)
	// if err != nil {
	// 	log.Fatalf("Could not locate the location input box. err: %v", err)

	// }
	// page.Keyboard.Press(input.ArrowDown)
	// time.Sleep(500 * time.Millisecond)
	// page.Keyboard.Press(input.Enter)
	// time.Sleep(500 * time.Millisecond)

	radiusDropdown, err := page.Element("label[aria-label='Radius']")
	handleError(err, "Error finding select radius in miles button", logger)
	handleError(radiusDropdown.Click(proto.InputMouseButtonLeft, 1), "Error clicking the miles selection dropdown", logger)
	time.Sleep(1 * time.Second)
	err = selectMiles(page, 80, logger)
	handleError(err, "Error while selecting the given miles by selectMiles function", logger)
	// handleError(page.Keyboard.Press(input.Enter), "Error while pressing enter to select the desired miles", logger)
	// handleError()
	// var arrowDown input.Key. = "ArrowDown"
	// page.Keyboard.Press(input.ArrowDown)
	// time.Sleep(500 * time.Millisecond)
	// page.Keyboard.Press(input.Enter)

	// if err != nil {
	// 	log.Fatalf("could not push down key on miles dropdown menu. err: %s", err)
	// }
	// handleError(err,"Error selecting given miles from the select miles dropdown",logger)
	locationModalApplyButton, err := page.Element("div[aria-label='Apply']")
	handleError(err, "Error finding apply button", logger)
	handleError(locationModalApplyButton.Click(proto.InputMouseButtonLeft, 1), "Error while clicking the Apply button inside location modal", logger)
	// page.MustElement("div[aria-label='Apply']").MustClick()

	time.Sleep(1 * time.Second)

	// data := page.MustElements(fmt.Sprintf("div.%s", joinClassNames("x9f619 x78zum5 x1r8uery xdt5ytf x1iyjqo2 xs83m0k x1e558r4 x150jy0e x1iorvi4 xjkvuk6 xnpuxes x291uyu x1uepa24")))
	// te := fmt.Sprintf("a.%s", joinClassNames("x1i10hfl xjbqb8w x1ejq31n xd10rxx x1sy0etr x17r0tee x972fbf xcfux6l x1qhh985 xm0m39n x9f619 x1ypdohk xt0psk2 xe8uvvx xdj266r x11i5rnm xat24cr x1mh8g0r xexx8yu x4uap5 x18d9i69 xkhd6sd x16tdsg8 x1hl2dhg xggy1nq x1a2a7pz x1heor9g x1sur9pj xkrqix3 x1lku1pv"))
	// mileage := page.MustElements(fmt.Sprintf("div.%s", itemMileage))
	imgTagClassName := fmt.Sprintf("img.%s", joinClassNames("xt7dq6l xl1xv1r x6ikm8r x10wlt62 xh8yej3"))
	beforeMileageSpanClassName := fmt.Sprintf("span.%s", joinClassNames("x193iq5w xeuugli x13faqbe x1vvkbs x10flsy6 x1lliihq x1s928wv xhkezso x1gmr53x x1cpjm7i x1fgarty x1943h6x x1tu3fi x3x7a5m x1nxh6w3 x1sibtaa xo1l8bm xi81zsa"))
	spanMileage := fmt.Sprintf("span.%s", joinClassNames("x1lliihq x6ikm8r x10wlt62 x1n2onr6 xlyipyv xuxw1ft"))
	actualCostSpanClassName := fmt.Sprintf("span.%s", joinClassNames("x193iq5w xeuugli x13faqbe x1vvkbs xlh3980 xvmahel x1n0sxbx x1lliihq x1s928wv xhkezso x1gmr53x x1cpjm7i x1fgarty x1943h6x x4zkp8e x3x7a5m x1lkfr7t x1lbecb7 x1s688f xzsf02u"))
	priceWasAtSpanClassName := fmt.Sprintf("span.%s", joinClassNames("x193iq5w xeuugli x13faqbe x1vvkbs xlh3980 xvmahel x1n0sxbx x1lliihq x1s928wv xhkezso x1gmr53x x1cpjm7i x1fgarty x1943h6x x4zkp8e x3x7a5m x1lkfr7t x1lbecb7 xk50ysn xi81zsa"))
	// data1 := page.MustElements(imgTagClassName)
	// data2 := page.MustElements(spanMileage)
	log.Println(actualCostSpanClassName)
	time.Sleep(1 * time.Second)

	scrollDuration := 10 * time.Second
	scrollInterval := 500 * time.Millisecond
	scrollEndTime := time.Now().Add(scrollDuration)
	for time.Now().Before(scrollEndTime) {

		handleError(page.Mouse.Scroll(0, 500, 10), "Error while trying to scroll through the page", logger) // Scroll vertically by 500 pixels in 10 steps
		time.Sleep(scrollInterval)
	}

	parentDIV := fmt.Sprintf("div.%s", "xjp7ctv")
	data, err := page.Elements(parentDIV)
	handleError(err, "Error finding all divs with class name: xjp7ctv", logger)

	// anchorClassName := fmt.Sprintf("div.x3ct3a4 a.%s", joinClassNames("x1i10hfl xjbqb8w x1ejq31n xd10rxx x1sy0etr x17r0tee x972fbf xcfux6l x1qhh985 xm0m39n x9f619 x1ypdohk xt0psk2 xe8uvvx xdj266r x11i5rnm xat24cr x1mh8g0r xexx8yu x4uap5 x18d9i69 xkhd6sd x16tdsg8 x1hl2dhg xggy1nq x1a2a7pz x1heor9g x1sur9pj xkrqix3 x1lku1pv"))
	fmt.Printf("%s %s", beforeMileageSpanClassName, spanMileage)
	var cars []carInformation
	time.Sleep(2 * time.Second)
	for _, divs := range data {
		var car carInformation
		// fmt.Println(divs.Text())
		links, err := divs.Element(fmt.Sprintf("a.%s", joinClassNames("x1i10hfl xjbqb8w x1ejq31n xd10rxx x1sy0etr x17r0tee x972fbf xcfux6l x1qhh985 xm0m39n x9f619 x1ypdohk xt0psk2 xe8uvvx xdj266r x11i5rnm xat24cr x1mh8g0r xexx8yu x4uap5 x18d9i69 xkhd6sd x16tdsg8 x1hl2dhg xggy1nq x1a2a7pz x1heor9g x1sur9pj xkrqix3 x1lku1pv")))
		if err != nil {
			fmt.Printf("error finding <a> tags with item's detailed information urls. Err:%v", err)
			continue
		}
		href, err := links.Attribute("href")
		if err != nil {
			fmt.Printf("error with item's <a> tags href attribute *string.href:%v Err:%v", *href, err)
		}
		car.CarDetailsURL = *href

		// links, err := divs.Element(anchorClassName)
		// if err != nil {
		// 	fmt.Println("err while getting anchor element", err)
		// 	continue
		// }
		// href, err := links.Attribute("href")
		// if err != nil {
		// 	fmt.Println("err while getting href element", err)
		// 	continue
		// }

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
			fmt.Printf("error while converting price string to int. Err:%v", err)
		}
		car.Price = costInt

		// actualCostSpan, err := divs.Element(actualCostSpanClassName)
		// if err != nil {
		// 	fmt.Println("err while getting price span", err)
		// 	continue
		// }
		// actualCostText, err := actualCostSpan.Text()
		// if err != nil {
		// 	fmt.Println("err while getting price", err)
		// 	continue
		// }
		// costInt, err := formatAndConvertPriceToInt(actualCostText)
		// if err != nil {
		// 	fmt.Println("err while converting price to integer", err)
		// 	continue
		// }
		// actualCost := strings.Replace(actualCostText, "$", "", 1)
		// actualCost = strings.Replace(actualCost, ",", "", 1)
		// costInt, err := strconv.Atoi(actualCost)
		// var priceWasAt int

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

		// if err == nil {
		// 	priceWasAtText, _ := priceWasAtSpan.Text()

		// 	priceWasAtInt, _ := formatAndConvertPriceToInt(priceWasAtText)
		// 	priceWasAt = priceWasAtInt
		// 	// fmt.Println("err while getting price was at span", err)
		// 	// continue
		// }
		// if err != nil {
		// 	fmt.Println("err while getting price was at text", err)
		// 	continue
		// }
		// if err != nil {
		// 	fmt.Println("err while converting price was at to int", err)
		// }

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

		// imgTagInfo, err := divs.Element(imgTagClassName)
		// if err != nil {
		// 	fmt.Println("err while getting src element", err)
		// 	continue
		// }
		// imgSourceURL, err := imgTagInfo.Attribute("src")
		// if err != nil {
		// 	fmt.Println("err while getting image source url property", err)
		// 	continue
		// }

		mileageElement, err := divs.Elements("div.x1iorvi4.x4uap5.xjkvuk6.xkhd6sd span span.x1lliihq.x6ikm8r.x10wlt62.x1n2onr6.xlyipyv.xuxw1ft")
		if err != nil {
			fmt.Printf("error finding car's mileage <span> tags. Err:%v", err)
		}
		// spanWithText, err := mileageElement
		if len(mileageElement) > 0 {

			t, _ := mileageElement[0].Text()
			tt, _ := mileageElement[1].Text()
			fmt.Println(t, tt)
			car.Location = t
			car.Mileage = tt

		}
		// tt, _ := mileageElement[1].Text()
		// for i, v := range mileageElement {
		// 	t, _ := v.Text()
		// 	fmt.Printf("index:%v, and the item:%v\n", i, t)
		// }

		// cityAndState, err := mileageElement[0].Text()
		// if err != nil {
		// 	return fmt.Errorf("error getting car's city, state. Err:%v", err)
		// }
		// car.Location = cityAndState

		// mileage, err := mileageElement[1].Text()
		// if err != nil {
		// 	return fmt.Errorf("error getting car's mileage. Err:%v", err)
		// }
		// car.Mileage = mileage

		// mileageElement, err := divs.Elements(spanMileage)
		// if err != nil {
		// 	fmt.Println("err while getting mileage element", err)
		// 	continue
		// }
		// log.Println(mileageElement[0],mileageElement[])
		// for i, v := range mileageElement {
		// 	log.Println(i, v)
		// }

		itemNameSpan, err := divs.Element(fmt.Sprintf("span.%s", joinClassNames("x1lliihq x6ikm8r x10wlt62 x1n2onr6")))
		if err != nil {
			fmt.Printf("error finding car's name <span>. Err:%v", err)
		}
		itemNameText, err := itemNameSpan.Text()
		if err != nil {
			fmt.Printf("error getting car's name text. Err:%v", err)
		}
		car.CarName = itemNameText

		// itemNameSpan, err := divs.Element(fmt.Sprintf("span.%s", joinClassNames("x1lliihq x6ikm8r x10wlt62 x1n2onr6")))
		// if err != nil {
		// 	fmt.Println("err while getting item name span", err)
		// }
		// itemName, err := itemNameSpan.Text()
		// if err != nil {
		// 	fmt.Println("err while getting item name text", err)
		// }
		// cityAndState, err := mileageElement[0].Text()
		// if err != nil {
		// 	fmt.Println("err while getting location text", err)
		// }
		// mileage, err := mileageElement[1].Text()
		// if err != nil {
		// 	fmt.Println("err while getting mileage element text", err)
		// 	continue
		// }
		// carNameAndLocation := strings.Split(*alt, "in")
		// car := carInformation{Price: costInt, PriceWasAt: priceWasAt, CarName: itemName, Location: cityAndState, Mileage: mileage, CarDetailsURL: *href, ImageURL: *imgSourceURL}
		cars = append(cars, car)

	}
	saveScrapedDataToJSON(cars)
	// fmt.Println(cars)

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

func saveScrapedDataToJSON(cars []carInformation) {

	file, err := os.Create("cars.json")
	if err != nil {
		log.Fatalf("Failed to create JSON file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(cars); err != nil {
		log.Fatalf("Failed to encode cars to JSON: %v", err)
	}

}

func joinClassNames(classes string) string {
	return strings.Join(strings.Split(classes, " "), ".")
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

func selectMiles(page *rod.Page, mile int, logger *zap.Logger) error {

	var selectedMileageIndex int
	for i, v := range milesValues {
		if mile == v {
			selectedMileageIndex = i
			break
		}
	}

	for i := 0; i < selectedMileageIndex+1; i++ {
		handleError(page.Keyboard.Press(input.ArrowDown), "Error while trying to press keyboard down arrow", logger)
		// err := page.Keyboard.Press(input.ArrowDown)
		// if err != nil {
		// return err
		// }
		time.Sleep(300 * time.Millisecond)
	}
	handleError(page.Keyboard.Press(input.Enter), "Error while pressing enter to select the desired miles", logger)
	return nil

}

// func addParamatersToURL(t string) string {
// 	te := t || ""
// }
