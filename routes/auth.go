package routes

import (
	"github.com/gin-gonic/gin"
	"xdu-health-card/controller"
)

func addAuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		auth.POST("/login", controller.Login)
	}
}
