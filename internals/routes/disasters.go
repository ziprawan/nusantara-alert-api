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

	c.JSON(200, response)
}
