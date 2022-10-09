package airvisual

type City struct {
	City     string
	State    string
	Country  string
	Location *Location
	Forecast []*Forecast
	Current  *Current
	History  *History
}
