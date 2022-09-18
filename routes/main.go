package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

var (
	router = gin.Default()
)

// Run will start the server
func Run(port string) {
	getRoutes()
	log.Fatal(router.Run(":" + port))
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {
	v1 := router.Group("/v1")
	addServeRoutes(v1)
	addAuthRoutes(v1)
}
