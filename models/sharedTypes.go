package airvisual

type Location struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Forecast struct {
	TS    string  `json:"ts"`
	AQIUS int     `json:"aqius"`
	AQICN int     `json:"aqicn"`
	TP    float64 `json:"tp,omitempty"`
	TPMin float64 `json:"tp_min,omitempty"`
	PR    float64 `json:"pr,omitempty"`
	HU    float64 `json:"hu,omitempty"`
	WS    float64 `json:"ws,omitempty"`
	WD    float64 `json:"wd,omitempty"`
	IC    string  `json:"ic,omitempty"`
}

type Weather struct {
	TS string  `json:"ts"`
	TP float64 `json:"tp"`
	PR float64 `json:"pr"`
	HU float64 `json:"hu"`
	WS float64 `json:"ws"`
	WD float64 `json:"wd"`
	IC string  `json:"ic"`
}

type Metric struct {
	CONC  float64 `json:"conc"`
	AQIUS int     `json:"aqius"`
	AQICN int     `json:"aqicn"`
}

type Pollution struct {
	TS     string  `json:"ts"`
	AQIUS  int     `json:"aqius"`
	MAINUS string  `json:"mainus"`
	AQICN  int     `json:"aqicn"`
	MAINCN string  `json:"maincn"`
	P2     *Metric `json:"p2,omitempty"`
	P1     *Metric `json:"p1,omitempty"`
	O3     *Metric `json:"o3,omitempty"`
	N2     *Metric `json:"n2,omitempty"`
	S2     *Metric `json:"s2,omitempty"`
	CO     *Metric `json:"co,omitempty"`
}

type Current struct {
	Weather   *Weather   `json:"weather"`
	Pollution *Pollution `json:"pollution"`
}

type History struct {
	Weather   []*Weather   `json:"weather"`
	Pollution []*Pollution `json:"pollution"`
}
