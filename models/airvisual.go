package airvisual

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"
)

var once sync.Once
var apiSingleInstance *Api

func New(APIKey string) *Api {
	if apiSingleInstance == nil {
		once.Do(
			func() {

				apiSingleInstance = &Api{
					client: http.DefaultClient,
					url:    "http://api.airvisual.com/v2/",
					APIKey: APIKey,
				}
			})
	}
	return apiSingleInstance
}

func (a *Api) GetCountries() ([]Country, error) {
	req, _ := http.NewRequest("GET", "http://api.airvisual.com/v2/countries", nil)
	q := req.URL.Query()
	q.Add("key", a.APIKey)
	req.URL.RawQuery = q.Encode()
	resp, _ := a.client.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		js := struct {
			baseResponse
			Data []Country `json:"data"`
		}{}
		json.NewDecoder(resp.Body).Decode(&js)
		return js.Data, nil
	}
	if resp.StatusCode == http.StatusBadRequest ||
		resp.StatusCode == http.StatusForbidden {

		js := struct {
			baseResponse
			Data errMsg `json:"data"`
		}{}
		json.NewDecoder(resp.Body).Decode(&js)
		return nil, errors.New(js.Data.Message)
	}
	return nil, errors.New("something went wrong")
}

func (a *Api) GetStatesByCountryName(countryName string) (states []State) {
	req, _ := http.NewRequest("GET", "http://api.airvisual.com/v2/states", nil)
	q := req.URL.Query()
	q.Add("key", a.APIKey)
	q.Add("country", countryName)
	req.URL.RawQuery = q.Encode()
	resp, _ := a.client.Do(req)
	defer resp.Body.Close()

	js := struct {
		Status string  `json:"status"`
		Data   []State `json:"data"`
	}{}

	json.NewDecoder(resp.Body).Decode(&js)

	states = js.Data
	return states
}

func (a *Api) GetCitiesByStateName(countryName, stateName string) (cities []City) {
	req, _ := http.NewRequest("GET", "http://api.airvisual.com/v2/states", nil)
	q := req.URL.Query()
	q.Add("key", a.APIKey)
	q.Add("country", countryName)
	q.Add("state", stateName)
	req.URL.RawQuery = q.Encode()
	resp, _ := a.client.Do(req)
	defer resp.Body.Close()

	js := struct {
		Status string `json:"status"`
		Data   []City `json:"data"`
	}{}

	json.NewDecoder(resp.Body).Decode(&js)

	cities = js.Data
	return cities
}

func (a *Api) GetCityByIP(ipAddress string) (city City) {
	req, _ := http.NewRequest("GET", "http://api.airvisual.com/v2/nearest_city", nil)
	q := req.URL.Query()
	q.Add("key", a.APIKey)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("x-forwarded-for", ipAddress)
	resp, _ := a.client.Do(req)
	defer resp.Body.Close()

	js := struct {
		Status string `json:"status"`
		Data   City   `json:"data"`
	}{}

	json.NewDecoder(resp.Body).Decode(&js)

	city = js.Data
	return city
}

func (a *Api) GetCityByCoordinates(lat, lon string) (city City) {
	req, _ := http.NewRequest("GET", "http://api.airvisual.com/v2/nearest_city", nil)
	q := req.URL.Query()
	q.Add("key", a.APIKey)
	q.Add("lat", lat)
	q.Add("lon", lon)
	req.URL.RawQuery = q.Encode()
	resp, _ := a.client.Do(req)
	defer resp.Body.Close()

	js := struct {
		Status string `json:"status"`
		Data   City   `json:"data"`
	}{}

	json.NewDecoder(resp.Body).Decode(&js)

	city = js.Data
	return city
}

func (a *Api) GetCity(countryName, stateName, cityName string) (city City) {
	req, _ := http.NewRequest("GET", "http://api.airvisual.com/v2/city", nil)
	q := req.URL.Query()
	q.Add("key", a.APIKey)
	q.Add("city", cityName)
	q.Add("state", stateName)
	q.Add("country", countryName)
	req.URL.RawQuery = q.Encode()
	resp, _ := a.client.Do(req)
	defer resp.Body.Close()

	js := struct {
		Status string `json:"status"`
		Data   City   `json:"data"`
	}{}

	json.NewDecoder(resp.Body).Decode(&js)

	city = js.Data
	return city
}

func (a *Api) getStationsByCityName(countryName, stateName, cityName string) (stations []Station) {
	req, _ := http.NewRequest("GET", "http://api.airvisual.com/v2/stations", nil)
	q := req.URL.Query()
	q.Add("key", a.APIKey)
	q.Add("city", cityName)
	q.Add("state", stateName)
	q.Add("country", countryName)
	req.URL.RawQuery = q.Encode()
	resp, _ := a.client.Do(req)
	defer resp.Body.Close()

	js := struct {
		Status string    `json:"status"`
		Data   []Station `json:"data"`
	}{}

	json.NewDecoder(resp.Body).Decode(&js)

	stations = js.Data
	return stations
}

func (a *Api) GetStationByIP(ipAddress string) (station Station) {
	req, _ := http.NewRequest("GET", "http://api.airvisual.com/v2/nearest_station", nil)
	q := req.URL.Query()
	q.Add("key", a.APIKey)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("x-forwarded-for", ipAddress)
	resp, _ := a.client.Do(req)
	defer resp.Body.Close()

	js := struct {
		Status string  `json:"status"`
		Data   Station `json:"data"`
	}{}

	json.NewDecoder(resp.Body).Decode(&js)

	station = js.Data
	return station
}

func (a *Api) GetStationByCoordinates(lat, lon string) (station Station) {
	req, _ := http.NewRequest("GET", "http://api.airvisual.com/v2/nearest_station", nil)
	q := req.URL.Query()
	q.Add("key", a.APIKey)
	q.Add("lat", lat)
	q.Add("lon", lon)
	req.URL.RawQuery = q.Encode()
	resp, _ := a.client.Do(req)
	defer resp.Body.Close()

	js := struct {
		Status string  `json:"status"`
		Data   Station `json:"data"`
	}{}

	json.NewDecoder(resp.Body).Decode(&js)

	station = js.Data
	return station
}

func (a *Api) getStation(countryName, stateName, cityName, stationName string) (station Station) {
	req, _ := http.NewRequest("GET", "http://api.airvisual.com/v2/station", nil)
	q := req.URL.Query()
	q.Add("key", a.APIKey)
	q.Add("station", stationName)
	q.Add("city", cityName)
	q.Add("state", stateName)
	q.Add("country", countryName)
	req.URL.RawQuery = q.Encode()
	resp, _ := a.client.Do(req)
	defer resp.Body.Close()

	js := struct {
		Status string  `json:"status"`
		Data   Station `json:"data"`
	}{}

	json.NewDecoder(resp.Body).Decode(&js)

	station = js.Data
	return station
}
