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

	request.Openid = c.GetString("openid")

	loginRes := service.LoginStuID(request.StuID, request.Passwd)
	if !loginRes.Success {
		c.JSON(http.StatusOK, loginRes)
		return
	}

	retCode, err := service.StorageThreeCheck(request)
	c.JSON(http.StatusOK, base.NewErrorResponse(err, retCode))
}

func StorageHealthCard(c *gin.Context) {
	var request *model.HealthCardRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, base.NewErrorResponse(err, base.BindDataFailed))
		return
	}

	request.Openid = c.GetString("openid")

	loginRes := service.LoginStuID(request.StuID, request.Passwd)
	if !loginRes.Success {
		c.JSON(http.StatusOK, loginRes)
		return
	}

	location, err := service.ConvertHealthCardLocation(request.Latitude, request.Longitude)
	if err != nil {
		c.JSON(http.StatusOK, base.NewErrorResponse(err, base.BindDataFailed))
		return
	}
	retCode, err := service.StorageHealthCard(request.ToHealthCard(location))
	c.JSON(http.StatusOK, base.NewErrorResponse(err, retCode))
}

func DeleteThreeCheck(c *gin.Context) {
	openid := c.GetString("openid")
	retCode, err := service.DeleteThreeCheck(openid)
	c.JSON(http.StatusOK, base.NewErrorResponse(err, retCode))
}

func DeleteHealthCard(c *gin.Context) {
	openid := c.GetString("openid")
	retCode, err := service.DeleteHealthCard(openid)
	c.JSON(http.StatusOK, base.NewErrorResponse(err, retCode))
}
