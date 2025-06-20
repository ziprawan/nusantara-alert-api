package main

import (
	"tubes-arc-api/internals/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set this to gin.DebugMode if you want to debug this program
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	paths := r.Group("/")
	routes.APIRoutes(paths)

	r.Run(":8181")
}
