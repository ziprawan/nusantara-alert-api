package utils

// Parsed result
type DisasterData struct {
	Type         string     `json:"type"`
	Coordinates  [2]float64 `json:"coordinates"`
	LocationName string     `json:"location_name"`
	IncidentTime int64      `json:"incident_time"`
}

// BMKG
type BMKGgempa struct {
	Tanggal     string
	Jam         string
	DateTime    string
	Coordinates string
	Lintang     string
	Bujur       string
	Magnitude   string
	Kedalaman   string
	Wilayah     string
	Potensi     string
}
type BMKGInfogempa struct {
	Gempa []BMKGgempa `json:"gempa"`
}
type BMKGResponse struct {
	Infogempa BMKGInfogempa
}

// NASA
type EONETEventCategory struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
type EONETEventSource struct {
	ID  string `json:"id"`
	Url string `json:"url"`
}
type EONETEventGeometry struct {
	Date        string    `json:"date"`
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
type EONETEvent struct {
	ID         string               `json:"id"`
	Title      string               `json:"title"`
	Link       string               `json:"link"`
	Categories []EONETEventCategory `json:"categories"`
	Sources    []EONETEventSource   `json:"sources"`
	Geometry   []EONETEventGeometry `json:"geometry"`
}
type NASAEONETResult struct {
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Link        string       `json:"link"`
	Events      []EONETEvent `json:"events"`
}
