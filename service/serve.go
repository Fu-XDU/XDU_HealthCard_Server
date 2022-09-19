package service

import (
	"errors"
	"xdu-health-card/database/mysql"
	"xdu-health-card/model"
)

func GetSummary(openid string) (summary *model.Summary, err error) {
	healthCard, _ := mysql.GetNeededHealthCardByOpenid(openid)
	threeCheck, _ := mysql.GetNeededThreeCheckByOpenid(openid)

	summary = &model.Summary{
		HealthCard: healthCard,
		ThreeCheck: threeCheck,
	}
	return
}

func StorageThreeCheck(threeCheck *model.ThreeCheck) (err error) {
	registered, err := mysql.ThreeCheckRegisteredByOpenid(threeCheck.Openid)
	if registered {
		return errors.New("已经提交过了")
	}
	_, err = mysql.InsertThreeCheck(threeCheck)
	return
}

func StorageHealthCard(healthCard *model.HealthCard) (err error) {
	registered, err := mysql.HealthCardRegisteredByOpenid(healthCard.Openid)
	if registered {
		return errors.New("已经提交过了")
	}

	_, err = mysql.InsertHealthCard(healthCard)
	return
}

func DeleteThreeCheck(openid string) (err error) {
	err = mysql.DelThreeCheckByOpenid(openid)
	return
}

func DeleteHealthCard(openid string) (err error) {
	err = mysql.DelHealthCardByOpenid(openid)
	return
}
