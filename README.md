# airvisual.go

```
go get github.com/Marsei-L/airvisual.go
```

# Code Example

```go
package main

import (
	"fmt"

	"github.com/Marsei-L/airvisual.go"
)

func main() {
	api := airvisual.New("Your key")
	countries, err := api.GetCountries()
	fmt.Println(countries[0].Country)
	states, err := api.GetStatesByCountryName(countries[0].Country)
	fmt.Println(states)
}
```
