package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

type carInformation struct {
	CarName       string `json:"car_name"`
	Location      string `json:"location"`
	Mileage       string `json:"mileage"`
	CarDetailsURL string `json:"car_details_url"`
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
	launcher := launcher.New().Headless(false).Set("disable-notifications")
	//log in and sign up frame class name
	//x78zum5 xdt5ytf x2lah0s x193iq5w x2bj2ny x1ey2m1c xayqjjm x9f619 xds687c x1xy6bms xn6708d x1s14bel x1ye3gou xixxii4 x17qophe x1u8a7rm
	url, err := launcher.Launch()
	if err != nil {
		log.Fatalf("Could not launch the browser. err: %v", err)
	}
	browser := rod.New().ControlURL(url)
	err = browser.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the browser. err: %v", err)
	}
	defer browser.Close()
	// // .SlowMotion(2 * time.Second)

	page := browser.MustPage("https://www.facebook.com/marketplace/108530542504412/search?minPrice=100&maxPrice=10000&daysSinceListed=1&maxMileage=100000&minMileage=50000&query=honda civic&exact=false")
	viewportOpts := proto.EmulationSetDeviceMetricsOverride{Width: 1680, Height: 1080, DeviceScaleFactor: 1, Mobile: false}
	page.SetViewport(&viewportOpts)

	closeButtonJoinedClasses := fmt.Sprintf("div.%s", joinClassNames(closeButtonClasses))
	selectLocationJoinedClasses := fmt.Sprintf("div.%s", joinClassNames(selectLocationClasses))
	// itemIMG := joinClassNames(itemsImgAndInformationClassName)
	// itemMileage := joinClassNames(itemMileageClassName)
	// joinedCloseButtonClasses := fmt.Sprintf("div.%s", closeButton)
	// joinedSelectLocationClasses := fmt.Sprintf("div.%s", selectLocation)
	signUpCloseButton, err := page.Element(closeButtonJoinedClasses)
	if err != nil {
		log.Fatalf("Could not locate sign up close button. err: %v", err)
	}
	signUpCloseButton.Click(proto.InputMouseButtonLeft, 1)
	// page.MustElement("div[aria-label='Close']").MustClick()

	locationButton, err := page.Element(selectLocationJoinedClasses)
	if err != nil {
		log.Fatalf("Could not locate the location button text. err: %v", err)
	}
	locationButton.Click(proto.InputMouseButtonLeft, 1)

	time.Sleep(3 * time.Second)

	// joinedRadiusClasses := fmt.Sprintf("div.%s", joinClassNames(selectRadiusClasses))
	locationInputBox, err := page.Element("label[aria-label='Location']")
	if err != nil {
		log.Fatalf("Could not locate the location input box. err: %v", err)

	}
	locationInputBox.Click(proto.InputMouseButtonLeft, 1)
	locationInputBox.Input("Newark, Delaware")

	page.Keyboard.Press(input.ArrowDown)
	time.Sleep(500 * time.Millisecond)
	page.Keyboard.Press(input.Enter)
	time.Sleep(500 * time.Millisecond)

	page.MustElement("label[aria-label='Radius']").MustClick()
	time.Sleep(1 * time.Second)
	// var arrowDown input.Key. = "ArrowDown"
	page.Keyboard.Press(input.ArrowDown)
	time.Sleep(500 * time.Millisecond)
	page.Keyboard.Press(input.Enter)
	time.Sleep(2 * time.Second)

	page.MustElement("div[aria-label='Apply']").MustClick()

	time.Sleep(2 * time.Second)

	// data := page.MustElements(fmt.Sprintf("div.%s", joinClassNames("x9f619 x78zum5 x1r8uery xdt5ytf x1iyjqo2 xs83m0k x1e558r4 x150jy0e x1iorvi4 xjkvuk6 xnpuxes x291uyu x1uepa24")))
	// te := fmt.Sprintf("a.%s", joinClassNames("x1i10hfl xjbqb8w x1ejq31n xd10rxx x1sy0etr x17r0tee x972fbf xcfux6l x1qhh985 xm0m39n x9f619 x1ypdohk xt0psk2 xe8uvvx xdj266r x11i5rnm xat24cr x1mh8g0r xexx8yu x4uap5 x18d9i69 xkhd6sd x16tdsg8 x1hl2dhg xggy1nq x1a2a7pz x1heor9g x1sur9pj xkrqix3 x1lku1pv"))
	// mileage := page.MustElements(fmt.Sprintf("div.%s", itemMileage))
	imgTagClassName := fmt.Sprintf("img.%s", joinClassNames("xt7dq6l xl1xv1r x6ikm8r x10wlt62 xh8yej3"))
	beforeMileageSpanClassName := fmt.Sprintf("span.%s", joinClassNames("x193iq5w xeuugli x13faqbe x1vvkbs x10flsy6 x1lliihq x1s928wv xhkezso x1gmr53x x1cpjm7i x1fgarty x1943h6x x1tu3fi x3x7a5m x1nxh6w3 x1sibtaa xo1l8bm xi81zsa"))
	spanMileage := fmt.Sprintf("span.%s", joinClassNames("x1lliihq x6ikm8r x10wlt62 x1n2onr6 xlyipyv xuxw1ft"))
	// data1 := page.MustElements(imgTagClassName)
	// data2 := page.MustElements(spanMileage)
	time.Sleep(3 * time.Second)

	scrollDuration := 10 * time.Second
	scrollInterval := 500 * time.Millisecond
	scrollEndTime := time.Now().Add(scrollDuration)
	for time.Now().Before(scrollEndTime) {
		page.Mouse.Scroll(0, 500, 10) // Scroll vertically by 500 pixels in 10 steps
		time.Sleep(scrollInterval)
	}

	parentDIV := fmt.Sprintf("div.%s", "xjp7ctv")
	data := page.MustElements(parentDIV)
	anchorClassName := fmt.Sprintf("div.x3ct3a4 a.%s", joinClassNames("x1i10hfl xjbqb8w x1ejq31n xd10rxx x1sy0etr x17r0tee x972fbf xcfux6l x1qhh985 xm0m39n x9f619 x1ypdohk xt0psk2 xe8uvvx xdj266r x11i5rnm xat24cr x1mh8g0r xexx8yu x4uap5 x18d9i69 xkhd6sd x16tdsg8 x1hl2dhg xggy1nq x1a2a7pz x1heor9g x1sur9pj xkrqix3 x1lku1pv"))
	fmt.Printf("%s %s", beforeMileageSpanClassName, spanMileage)
	var cars []carInformation
	for _, divs := range data {
		// fmt.Println(divs.Text())
		links, err := divs.Element(anchorClassName)
		if err != nil {
			fmt.Println("err while getting anchor element", err)
			continue
		}
		href, err := links.Attribute("href")
		if err != nil {
			fmt.Println("err while getting href element")
			continue
		}
		imgTagInfo, err := divs.Element(imgTagClassName)
		if err != nil {
			fmt.Println("err while getting alt element", err)
			continue
		}
		alt, err := imgTagInfo.Attribute("alt")
		if err != nil {
			fmt.Println("err while getting alt property")
			continue
		}

		mileageElement, err := divs.Elements(spanMileage)
		if err != nil {
			fmt.Println("err while getting mileage element")
			continue
		}
		mileage, err := mileageElement[1].Text()
		if err != nil {
			fmt.Println("err while getting mileage element text")
			continue
		}
		carNameAndLocation := strings.Split(*alt, "in")
		car := carInformation{CarName: carNameAndLocation[0], Location: carNameAndLocation[1], Mileage: mileage, CarDetailsURL: *href}
		cars = append(cars, car)
		// fmt.Printf("each item external link: %s, information from image alt prop: %s, mileage:%v\n", *href, *alt, mileage)

		// fmt.Printf("%s %s", beforeMileageSpanClassName, spanMileage)
		// fmt.Println(*href)
		// fmt.Println(*alt)
		// linksToItem := divs.MustElement(fmt.Sprintf("a.%s", joinClassNames("x1i10hfl xjbqb8w x1ejq31n xd10rxx x1sy0etr x17r0tee x972fbf xcfux6l x1qhh985 xm0m39n x9f619 x1ypdohk xt0psk2 xe8uvvx xdj266r x11i5rnm xat24cr x1mh8g0r xexx8yu x4uap5 x18d9i69 xkhd6sd x16tdsg8 x1hl2dhg xggy1nq x1a2a7pz x1heor9g x1sur9pj xkrqix3 x1lku1pv")))

		// href, err := linksToItem.Attribute("href")
		// if err != nil {
		// 	log.Println("Error finding <a>href element in this vehicle card:", err)
		// 	continue
		// }
		// fmt.Printf("img alt: %s\n", *href)

	}
	fmt.Println(cars)
	// fmt.Printf("and the mileage:%s", data)
	// page.HTML()
	// page.MustElement("#email").MustInput(email)
	// page.MustElement("#pass").MustInput(pass)
	// page.MustElement("button[name='login']").MustClick()
	// // fmt.Print(page.HTML())

	// // page.MustWaitLoad()
	// time.Sleep(3 * time.Second)

	// page.MustNavigate("https://www.facebook.com/marketplace/category/vehicles")

	// page.WaitStable(3 * time.Second)

	// // time.Sleep(5 * time.Second)

	// filtersDivID := page.MustElement("#seo_filters")
	// fmt.Println(filtersDivID)
	// firstChild := filtersDivID.MustElement(">:first-child")
	// // fmt.Println(firstChild)
	// firstChild.MustClick()

	// log.Println(page.MustInfo())
}

func joinClassNames(classes string) string {
	return strings.Join(strings.Split(classes, " "), ".")
}

// func addParamatersToURL(t string) string {
// 	te := t || ""
// }
