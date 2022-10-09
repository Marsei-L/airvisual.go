package main

import (
	"fmt"

	"github.com/Marsei-L/airvisual.go"
)

func main() {
	api := airvisual.New("Your key")
	countries, _ := api.GetCountries()
	fmt.Println(countries[0].Country)
	// states := api.GetStatesByCountryName(countries[0].Country)
	// fmt.Println(states)
}
