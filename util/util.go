package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	// "time"

	// "github.com/go-rod/rod"
	// "github.com/go-rod/rod/lib/input"

	// "github.com/jeffrey1200/facebook-marketplace-scraper/internals/scraper"

	// "github.com/jeffrey1200/facebook-marketplace-scraper/internals/scraper"
	"github.com/jeffrey1200/facebook-marketplace-scraper/internals/models"
	"go.uber.org/zap"
)

func HandleError(err error, msg string, logger *zap.Logger) {
	if err != nil {
		logger.Error(msg, zap.Error(err))
	}
}

func FormatAndConvertPriceToInt(price string) (int, error) {

	actualCost := strings.Replace(price, "$", "", 1)
	actualCost = strings.Replace(actualCost, ",", "", 1)
	costInt, err := strconv.Atoi(actualCost)
	if err != nil {
		return 0, err
	}
	return costInt, nil

}

//	func JoinClassNames(classes string) string {
//		return strings.Join(strings.Split(classes, " "), ".")
//	}
//
// /home/jeffrey/Documents/facebook-marketplace-scraper
func SaveScrapedDataToJSON(cars models.CarData) {
	// cars = append(cars, time.Now().UTC())
	// dateObj := map[string]time.Time{"creation_date": time.Now().UTC()}

	fileName := fmt.Sprintf("/home/jeffrey/Documents/facebook-marketplace-scraper/output/%v_cars.json", cars.CreationDate)
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed to create JSON file: %v", err)
	}
	defer file.Close()
	// file.WriteString(fmt.Sprintf("creation_date:%v", time.Now().UTC()))
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(cars); err != nil {
		log.Fatalf("Failed to encode cars to JSON: %v", err)
	}

}

func JoinClassNames(classes string) string {
	return strings.Join(strings.Split(classes, " "), ".")
}

// var milesValues = map[int]int{
// 	0:  1,
// 	1:  2,
// 	2:  5,
// 	3:  10,
// 	4:  20,
// 	5:  40,
// 	6:  60,
// 	7:  80,
// 	8:  100,
// 	9:  250,
// 	10: 500,
// }

// // type milesValues map[int]string

// func SelectMiles(page *rod.Page, mile int, logger *zap.Logger) error {

// 	var selectedMileageIndex int
// 	for i, v := range milesValues {
// 		if mile == v {
// 			selectedMileageIndex = i
// 			break
// 		}
// 	}

// 	for i := 0; i < selectedMileageIndex+1; i++ {
// 		HandleError(page.Keyboard.Press(input.ArrowDown), "Error while trying to press keyboard down arrow", logger)
// 		// err := page.Keyboard.Press(input.ArrowDown)
// 		// if err != nil {
// 		// return err
// 		// }
// 		time.Sleep(300 * time.Millisecond)
// 	}
// 	HandleError(page.Keyboard.Press(input.Enter), "Error while pressing enter to select the desired miles", logger)
// 	return nil

// }
