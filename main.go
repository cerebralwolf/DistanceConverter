package main

import (
	"fmt"
	"log"
	"strings"
)

func convertDistance(distance float64, unit string) (float64, string) {
	// constants for converting miles to km and km to miles
	const kmToMiles = 0.621371
	const milesToKm = 1.60934

	unit = strings.ToLower(unit)

	if unit == "km" {
		return distance * kmToMiles, "miles"
	} else if unit == "miles" {
		return distance * milesToKm, "km"
	}

	return 0, ""
}

func main() {
	var distance float64
	var unit string

	fmt.Print("Enter the distance: ")
	_, err := fmt.Scanf("%f", &distance)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter the unit in km or miles: ")
	_, err = fmt.Scanf("%s", &unit)
	if err != nil {
		log.Fatal(err)
	}

	convertedDistance, convertedUnit := convertDistance(distance, unit)
	if convertedUnit == "" {
		fmt.Println("Invalid unit. Please enter 'km' or 'miles'.")
		return
	}

	fmt.Printf("%.2f %s is equal to %.2f %s\n", distance, unit, convertedDistance, convertedUnit)

}
