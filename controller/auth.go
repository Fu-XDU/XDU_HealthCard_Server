package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xdu-health-card/model/base"
	"xdu-health-card/service"
)

func Login(c *gin.Context) {
	var request *struct {
		Code string `json:"code"`
	}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, base.NewErrorResponse(err, base.BindDataFailed))
		return
	}
	response := service.Login(request.Code)
	c.JSON(http.StatusOK, response)
}
