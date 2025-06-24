package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func FetchNASAEONET() ([]DisasterData, error) {
	var res []DisasterData

	resp, err := http.Get("https://eonet.gsfc.nasa.gov/api/v3/events?bbox=94.0,6.1,141.0,-11.0")
	if err != nil {
		fmt.Println("nasa: Failed to fetch URL:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var nasa NASAEONETResult
	err = json.NewDecoder(resp.Body).Decode(&nasa)
	if err != nil {
		fmt.Println("nasa: Failed to unmarshal response text:", err)
		return nil, err
	}

	for i, volc := range nasa.Events {
		var tmp DisasterData
		tmp.Type = "volcano"

		splitted := strings.Split(volc.Title, ",")
		tmp.LocationName = splitted[0]

		if len(volc.Geometry) == 0 {
			fmt.Printf("eonet: Failed to parse Geometry at index %d: Expected geometry length is 1 or more, got %d", i, len(volc.Geometry))
			return nil, fmt.Errorf("eonet: Failed to parse Geometry at index %d: Expected geometry length is 1 or more, got %d", i, len(volc.Geometry))
		}
		coord := volc.Geometry[0].Coordinates
		if len(coord) != 2 {
			fmt.Printf("eonet: Failed to parse Geometry at index %d: Expected coordinates length is 2, got %d", i, len(coord))
			return nil, fmt.Errorf("eonet: Failed to parse Geometry at index %d: Expected coordinates length is 2, got %d", i, len(coord))
		}
		tmp.Coordinates = [2]float64{coord[1], coord[0]}

		t, err := time.Parse(time.RFC3339, volc.Geometry[0].Date)
		if err != nil {
			fmt.Printf("eonet: Failed to parse Geomatry at index %d: Failed to parse Date: %s", i, err.Error())
			return nil, err
		}
		tmp.IncidentTime = t.UnixMilli()

		res = append(res, tmp)
	}

	return res, nil
}
