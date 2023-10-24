package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Router
	router := gin.Default()

	// Initilize routes
	initializeRoutes(router)

	// Ru nthe server
	router.Run(":8080")
}
