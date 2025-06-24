package routes

import (
	"fmt"
	"tubes-arc-api/internals/utils"

	"github.com/gin-gonic/gin"
)

func DisastersHandler(c *gin.Context) {
	var response DisasterResponse

	bmkg, err := utils.FetchBMKG()
	if err != nil {
		fmt.Println("disasters: Failed to fetch BMKG result:", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	response.Data.Earthquakes = bmkg

	eonet, err := utils.FetchNASAEONET()
	if err != nil {
		fmt.Println("disasters: Failed to fetch NASA EONET result:", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	response.Data.Volcanoes = eonet

	adinet, err := utils.FetchAdinet()
	if err != nil {
		fmt.Println("adinet: Failed to fetch NASA EONET result:", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	response.Data.LocalDisasters = adinet

	response.Sources = map[string]string{
		"earthquakes": "https://data.bmkg.go.id/DataMKG/TEWS/gempaterkini.json",
		"tsunami":     "https://data.bmkg.go.id/DataMKG/TEWS/gempaterkini.json",
		"volcano":     "https://eonet.gsfc.nasa.gov/api/v3/events?bbox=94.0,6.1,141.0,-11.0",
		"flood":       "https://adinet.ahacentre.org/report/list?keywords=Indonesia&sort=new",
		"tornadoes":   "https://adinet.ahacentre.org/report/list?keywords=Indonesia&sort=new",
	}

	c.JSON(200, response)
}
