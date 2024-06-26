package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type UrlParameters struct {
	City            string
	State           string
	MinPrice        int
	MaxPrice        int
	MinMileage      int
	MaxMileage      int
	DaysSinceListed int
	QueryName       string
}

func Repl() UrlParameters {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter city: ")
	cityInput, _ := reader.ReadString('\n')
	cityInput = strings.TrimSpace(cityInput)

	fmt.Print("Enter State: ")
	stateInput, _ := reader.ReadString('\n')
	stateInput = strings.TrimSpace(stateInput)

	fmt.Print("Enter minPrice: ")
	minPriceInput, _ := reader.ReadString('\n')
	minPriceInput = strings.TrimSpace(minPriceInput)
	minPrice, err := strconv.Atoi(minPriceInput)
	if err != nil {
		fmt.Println("Invalid minPrice input")
		return UrlParameters{}
	}

	fmt.Print("Enter maxPrice: ")
	maxPriceInput, _ := reader.ReadString('\n')
	maxPriceInput = strings.TrimSpace(maxPriceInput)
	maxPrice, err := strconv.Atoi(maxPriceInput)
	if err != nil {
		fmt.Println("Invalid maxPrice input")
		return UrlParameters{}
	}

	fmt.Print("Enter minMileage: ")
	minMileageInput, _ := reader.ReadString('\n')
	minMileageInput = strings.TrimSpace(minMileageInput)
	minMileage, err := strconv.Atoi(minMileageInput)
	if err != nil {
		fmt.Println("Invalid minMileage input")
		return UrlParameters{}
	}

	fmt.Print("Enter maxMileage: ")
	maxMileageInput, _ := reader.ReadString('\n')
	maxMileageInput = strings.TrimSpace(maxMileageInput)
	maxMileage, err := strconv.Atoi(maxMileageInput)
	if err != nil {
		fmt.Println("Invalid maxMileage input")
	}

	fmt.Print("Enter days since last listed. You can choose either 1 or 7: ")
	daysSinceListedInput, _ := reader.ReadString('\n')
	daysSinceListedInput = strings.TrimSpace(daysSinceListedInput)
	daysSinceListed, err := strconv.Atoi(daysSinceListedInput)
	if err != nil {
		fmt.Println("Invalid daysSinceListed input")
		return UrlParameters{}
	}

	fmt.Print("Enter the car name you want to search: ")
	queryNameInput, err := reader.ReadString('\n')
	queryNameInput = strings.TrimSpace(queryNameInput)
	// queryName, err := strconv.Atoi(queryNameInput)
	if err != nil {
		fmt.Println("Invalid query name input")
		return UrlParameters{}
	}

	completedURl := UrlParameters{City: cityInput, State: stateInput, MinPrice: minPrice, MaxPrice: maxPrice, MinMileage: minMileage, MaxMileage: maxMileage, DaysSinceListed: daysSinceListed, QueryName: queryNameInput}

	return completedURl
}
