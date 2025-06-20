package routes

import "tubes-arc-api/internals/utils"

type DisasterResponseData struct {
	Earthquakes    []utils.DisasterData `json:"earthquakes"`
	LocalDisasters []utils.DisasterData `json:"local_disasters"`
	Volcanoes      []utils.DisasterData `json:"volcanoes"`
}

type DisasterResponse struct {
	Data    DisasterResponseData `json:"data"`
	Sources map[string]string    `json:"sources"`
}
