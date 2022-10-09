package main

import (
	"fmt"

	"example.com/airvisual"
)

func main() {
	api := airvisual.New("Your key")
	countries, _ := api.GetCountries()
	fmt.Println(countries[0].Country)
	// states := api.GetStatesByCountryName(countries[0].Country)
	// fmt.Println(states)
}
