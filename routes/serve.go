package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"xdu-health-card/auth"
	"xdu-health-card/controller"
	"xdu-health-card/model/base"
)

func addServeRoutes(rg *gin.RouterGroup) {
	serve := rg.Group("/serve")
	serve.Use(authMid)
	{
		serve.GET("/summary", controller.GetSummary)
		serve.POST("/three-check", controller.StorageThreeCheck)
		serve.POST("/health-card", controller.StorageHealthCard)
		serve.DELETE("/three-check", controller.DeleteThreeCheck)
		serve.DELETE("/health-card", controller.DeleteHealthCard)
	}
}

func authMid(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	authorizationArray := strings.Split(authorization, " ")
	if len(authorizationArray) != 2 || authorizationArray[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, base.NewErrorResponse(nil, base.Unauthorized))
		c.Abort()
		return
	}

	clames, err := auth.ParseJwt(authorizationArray[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, base.NewErrorResponse(err, base.Unauthorized))
		c.Abort()
		return
	}
	c.Set("openid", clames.OpenId)
	c.Next()
}
