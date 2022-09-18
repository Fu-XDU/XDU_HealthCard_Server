package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xdu-health-card/model"
	"xdu-health-card/model/base"
	"xdu-health-card/service"
)

func GetSummary(c *gin.Context) {
	openid := c.GetString("openid")
	summary, err := service.GetSummary(openid)
	if err != nil {
		c.JSON(http.StatusOK, base.NewErrorResponse(err, base.BindDataFailed))
	} else {
		c.JSON(http.StatusOK, base.NewDataResponse(summary))
	}
}

func StorageThreeCheck(c *gin.Context) {
	var request *model.ThreeCheck
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, base.NewErrorResponse(err, base.BindDataFailed))
		return
	}

	err = service.StorageThreeCheck(request)
	if err != nil {
		c.JSON(http.StatusOK, base.NewErrorResponse(err, base.BindDataFailed))
	} else {
		c.JSON(http.StatusOK, base.NewResponse())
	}
}

func StorageHealthCard(c *gin.Context) {
	var request *model.HealthCard
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, base.NewErrorResponse(err, base.BindDataFailed))
		return
	}

	err = service.StorageHealthCard(request)
	if err != nil {
		c.JSON(http.StatusOK, base.NewErrorResponse(err, base.BindDataFailed))
	} else {
		c.JSON(http.StatusOK, base.NewResponse())
	}
}

func DeleteThreeCheck(c *gin.Context) {

}

func DeleteHealthCard(c *gin.Context) {

}
