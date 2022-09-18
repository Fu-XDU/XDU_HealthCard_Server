package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xdu-health-card/model/base"
	"xdu-health-card/service"
)

func Login(c *gin.Context) {
	code := c.GetString("code")
	data, err := service.Login(code)
	if err != nil {
		c.JSON(http.StatusOK, base.NewErrorResponse(err, base.BindDataFailed))
	} else {
		c.JSON(http.StatusOK, base.NewDataResponse(data))
	}
}
