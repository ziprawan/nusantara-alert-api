package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func FetchAdinet() ([]DisasterData, error) {
	var res []DisasterData

	resp, err := http.Get("https://adinet.ahacentre.org/report/list?keywords=Indonesia&sort=new")
	if err != nil {
		fmt.Println("bmkg: Failed to fetch URL:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var bmkg BMKGResponse
	err = json.NewDecoder(resp.Body).Decode(&bmkg)
	if err != nil {
		fmt.Println("bmkg: Failed to unmarshal response text:", err)
		return nil, err
	}

	for i, gempa := range bmkg.Infogempa.Gempa {
		var tmp DisasterData
		tmp.LocationName = gempa.Wilayah

		if strings.Contains((strings.ToLower(gempa.Potensi)), "tidak") {
			tmp.Type = "earthquake"
		} else {
			tmp.Type = "tsunami"
		}

		t, err := time.Parse(time.RFC3339, gempa.DateTime)
		if err != nil {
			fmt.Printf("bmkg: Failed to parse DateTime at index %d: %s\n", i, err.Error())
			return nil, fmt.Errorf("bmkg: Failed to parse DateTime at index %d: %s\n", i, err.Error())
		}
		tmp.IncidentTime = t.UnixMilli()

		splitted := strings.Split(gempa.Coordinates, ",")
		if len(splitted) != 2 {
			fmt.Printf("bmkg: Failed to parse Coordinates at index %d: Expected splitted length is 2, got %d\n", i, len(splitted))
			return nil, fmt.Errorf("bmkg: Failed to parse Coordinates at index %d: Expected splitted length is 2, got %d\n", i, len(splitted))
		}
		left, err := strconv.ParseFloat(splitted[0], 64)
		if err != nil {
			fmt.Printf("bmkg: Failed to parse Coordinates at index %d at position 0: %s\n", i, err.Error())
			return nil, fmt.Errorf("bmkg: Failed to parse Coordinates at index %d: Expected splitted length is 2, got %d\n", i, len(splitted))
		}
		right, err := strconv.ParseFloat(splitted[1], 64)
		if err != nil {
			fmt.Printf("bmkg: Failed to parse Coordinates at index %d at position 1: %s\n", i, err.Error())
			return nil, fmt.Errorf("bmkg: Failed to parse Coordinates at index %d: Expected splitted length is 2, got %d\n", i, len(splitted))
		}
		tmp.Coordinates = [2]float64{left, right}

		res = append(res, tmp)
	}

	return res, nil
}
