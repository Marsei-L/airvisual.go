package airvisual

type Station struct {
	Name     string
	State    string
	Country  string
	Location *Location
	Forecast []*Forecast
	Current  *Current
	History  *History
}
