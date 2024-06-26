package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jeffrey1200/facebook-marketplace-scraper/internals/cli"
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

func SaveScrapedDataToJSON(cars models.CarData, searchParameters cli.UrlParameters) error {
	// cars = append(cars, time.Now().UTC())
	// dateObj := map[string]time.Time{"creation_date": time.Now().UTC()}

	fileName := fmt.Sprintf("/home/jeffrey/Documents/facebook-marketplace-scraper/output/%v_%s.json", cars.CreationDate, searchParameters.QueryName)
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed to create JSON file: %v", err)
		return err
	}
	defer file.Close()
	// file.WriteString(fmt.Sprintf("creation_date:%v", time.Now().UTC()))
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(cars); err != nil {
		log.Fatalf("Failed to encode cars to JSON: %v", err)
		return err
	}

	return nil

}

func JoinClassNames(classes string) string {
	return strings.Join(strings.Split(classes, " "), ".")
}
