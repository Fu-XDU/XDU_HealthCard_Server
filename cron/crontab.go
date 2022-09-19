package cron

import (
	"github.com/labstack/gommon/log"
	"github.com/robfig/cron"
	"xdu-health-card/database/mysql"
	"xdu-health-card/service"
)

func Start() {
	c := cron.New()
	_ = c.AddFunc("30 7 * * *", healthCardWorker)
	_ = c.AddFunc("30 8,12,18 * * *", threeCheckWorker)
	c.Start()
}

func healthCardWorker() {
	allHealthCard, err := mysql.GetAllHealthCard()
	if err != nil {
		log.Error(err)
		return
	}

	for i := range allHealthCard {
		item := allHealthCard[i]
		response := service.LoginStuID(item.StuID, item.Passwd)
		if !response.Success {
			log.Error(response)
			continue
		}

		cookies := response.Data.(string)
		err = service.SubmitHealthCard(cookies, item.Location)
		if err != nil {
			log.Error(response)
		}
	}
}

func threeCheckWorker() {
	allThreeCheck, err := mysql.GetAllThreeCheck()
	if err != nil {
		log.Error(err)
		return
	}

	for i := range allThreeCheck {
		item := allThreeCheck[i]
		response := service.LoginStuID(item.StuID, item.Passwd)
		if !response.Success {
			log.Error(response)
			continue
		}

		cookies := response.Data.(string)
		err = service.SubmitThreeCheck(cookies, item.Location)
		if err != nil {
			log.Error(response)
		}
	}
}
