package routes

import (
	"github.com/gin-gonic/gin"
	"xdu-health-card/controller"
)

func addServeRoutes(rg *gin.RouterGroup) {
	rg.GET("/summary", controller.GetSummary)
	rg.POST("/three-check", controller.StorageThreeCheck)
	rg.POST("/health-card", controller.StorageHealthCard)
	rg.DELETE("/three-check", controller.DeleteThreeCheck)
	rg.DELETE("/health-card", controller.DeleteHealthCard)
}
