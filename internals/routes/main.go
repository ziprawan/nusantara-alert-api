package routes

import "github.com/gin-gonic/gin"

func APIRoutes(router *gin.RouterGroup) {
	// /api/*
	apiRouter := router.Group("/api")
	apiRouter.GET("/disasters", DisastersHandler)
}
