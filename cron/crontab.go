package cron

import (
	"github.com/robfig/cron"
)

func Start() {
	c := cron.New()
	_ = c.AddFunc("30 7 * * *", healthCardWorker)
	_ = c.AddFunc("30 8,12,18 * * *", threeCheckWorker)
	c.Start()
}

func healthCardWorker() {

}

func threeCheckWorker() {

}
