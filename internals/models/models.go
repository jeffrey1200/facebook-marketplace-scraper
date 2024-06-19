package models

type CarInformation struct {
	Price         int    `json:"price"`
	PriceWasAt    int    `json:"price_was_at"`
	CarName       string `json:"car_name"`
	Location      string `json:"location"`
	Mileage       string `json:"mileage"`
	CarDetailsURL string `json:"car_details_url"`
	ImageURL      string `json:"image_url"`
}

type CarData struct {
	CreationDate string           `json:"creation_date"`
	AmountOfCars int              `json:"amount_of_cars"`
	Cars         []CarInformation `json:"cars"`
}
