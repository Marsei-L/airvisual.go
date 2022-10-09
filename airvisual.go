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

func (a *Api) GetCountries() ([]*Country, error) {
	req, _ := http.NewRequest("GET", "http://api.airvisual.com/v2/countries", nil)
	q := req.URL.Query()
	q.Add("key", a.APIKey)
	req.URL.RawQuery = q.Encode()
	resp, err := a.client.Do(req)
	if err != nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			js := struct {
				baseResponse
				Data []*Country `json:"data"`
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
	}
	return nil, errors.New("something went wrong")
}

func (a *Api) GetStatesByCountryName(countryName string) ([]*State, error) {
	req, _ := http.NewRequest("GET", "http://api.airvisual.com/v2/states", nil)
	q := req.URL.Query()
	q.Add("key", a.APIKey)
	q.Add("country", countryName)
	req.URL.RawQuery = q.Encode()
	resp, err := a.client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			js := struct {
				baseResponse
				Data []*State `json:"data"`
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
	}
	return nil, errors.New("something went wrong")
}

func (a *Api) GetCitiesByStateName(countryName, stateName string) ([]*City, error) {
	req, _ := http.NewRequest("GET", "http://api.airvisual.com/v2/states", nil)
	q := req.URL.Query()
	q.Add("key", a.APIKey)
	q.Add("country", countryName)
	q.Add("state", stateName)
	req.URL.RawQuery = q.Encode()
	resp, err := a.client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			js := struct {
				baseResponse
				Data []*City `json:"data"`
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
	}
	return nil, errors.New("something went wrong")
}

func (a *Api) GetCityByIP(ipAddress string) (*City, error) {
	req, _ := http.NewRequest("GET", "http://api.airvisual.com/v2/nearest_city", nil)
	q := req.URL.Query()
	q.Add("key", a.APIKey)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("x-forwarded-for", ipAddress)
	resp, err := a.client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			js := struct {
				baseResponse
				Data *City `json:"data"`
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
	}
	return nil, errors.New("something went wrong")
}

func (a *Api) GetCityByCoordinates(lat, lon string) (*City, error) {
	req, _ := http.NewRequest("GET", "http://api.airvisual.com/v2/nearest_city", nil)
	q := req.URL.Query()
	q.Add("key", a.APIKey)
	q.Add("lat", lat)
	q.Add("lon", lon)
	req.URL.RawQuery = q.Encode()
	resp, err := a.client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			js := struct {
				baseResponse
				Data *City `json:"data"`
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
	}
	return nil, errors.New("something went wrong")
}

func (a *Api) GetCity(countryName, stateName, cityName string) (*City, error) {
	req, _ := http.NewRequest("GET", "http://api.airvisual.com/v2/city", nil)
	q := req.URL.Query()
	q.Add("key", a.APIKey)
	q.Add("city", cityName)
	q.Add("state", stateName)
	q.Add("country", countryName)
	req.URL.RawQuery = q.Encode()
	resp, err := a.client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			js := struct {
				baseResponse
				Data *City `json:"data"`
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
	}
	return nil, errors.New("something went wrong")
}

func (a *Api) GetStationsByCityName(countryName, stateName, cityName string) ([]*Station, error) {
	req, _ := http.NewRequest("GET", "http://api.airvisual.com/v2/stations", nil)
	q := req.URL.Query()
	q.Add("key", a.APIKey)
	q.Add("city", cityName)
	q.Add("state", stateName)
	q.Add("country", countryName)
	req.URL.RawQuery = q.Encode()
	resp, err := a.client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			js := struct {
				baseResponse
				Data []*Station `json:"data"`
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
	}
	return nil, errors.New("something went wrong")
}

func (a *Api) GetStationByIP(ipAddress string) (*Station, error) {
	req, _ := http.NewRequest("GET", "http://api.airvisual.com/v2/nearest_station", nil)
	q := req.URL.Query()
	q.Add("key", a.APIKey)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("x-forwarded-for", ipAddress)
	resp, err := a.client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			js := struct {
				baseResponse
				Data *Station `json:"data"`
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
	}
	return nil, errors.New("something went wrong")
}

func (a *Api) GetStationByCoordinates(lat, lon string) (*Station, error) {
	req, _ := http.NewRequest("GET", "http://api.airvisual.com/v2/nearest_station", nil)
	q := req.URL.Query()
	q.Add("key", a.APIKey)
	q.Add("lat", lat)
	q.Add("lon", lon)
	req.URL.RawQuery = q.Encode()
	resp, err := a.client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			js := struct {
				baseResponse
				Data *Station `json:"data"`
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
	}

	return nil, errors.New("something went wrong")
}

func (a *Api) GetStation(countryName, stateName, cityName, stationName string) (*Station, error) {
	req, _ := http.NewRequest("GET", "http://api.airvisual.com/v2/station", nil)
	q := req.URL.Query()
	q.Add("key", a.APIKey)
	q.Add("station", stationName)
	q.Add("city", cityName)
	q.Add("state", stateName)
	q.Add("country", countryName)
	req.URL.RawQuery = q.Encode()
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, errors.New("something went wrong with while getting the data")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		js := struct {
			baseResponse
			Data *Station `json:"data"`
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
