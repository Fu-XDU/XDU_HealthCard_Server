package service

import (
	"xdu-health-card/database/mysql"
	"xdu-health-card/database/redis"
	"xdu-health-card/model"
	"xdu-health-card/model/base"
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

func StorageThreeCheck(threeCheck *model.ThreeCheck) (retCode *base.RetCode, err error) {
	var lockKey = "StorageThreeCheck_" + threeCheck.Openid

	retCode, err = redis.Lock(lockKey)
	if retCode.Code != base.SUCCESS.Code {
		return
	}

	defer redis.Unlock(lockKey)

	registered, err := mysql.ThreeCheckRegisteredByOpenid(threeCheck.Openid)
	if registered {
		return base.CannotSubmitAgain, err
	}

	_, err = mysql.InsertThreeCheck(threeCheck)
	if err != nil {
		retCode = base.RedisCurdFailed
	}
	return
}

func StorageHealthCard(healthCard *model.HealthCard) (retCode *base.RetCode, err error) {
	var lockKey = "StorageHealthCard_" + healthCard.Openid

	retCode, err = redis.Lock(lockKey)
	if retCode != base.SUCCESS {
		return
	}

	defer redis.Unlock(lockKey)

	registered, err := mysql.HealthCardRegisteredByOpenid(healthCard.Openid)
	if registered {
		return base.CannotSubmitAgain, err
	}

	_, err = mysql.InsertHealthCard(healthCard)
	if err != nil {
		retCode = base.RedisCurdFailed
	}
	return
}

func DeleteThreeCheck(openid string) (retCode *base.RetCode, err error) {
	var lockKey = "DeleteThreeCheck_" + openid

	retCode, err = redis.Lock(lockKey)
	if retCode != base.SUCCESS {
		return
	}

	defer redis.Unlock(lockKey)

	err = mysql.DelThreeCheckByOpenid(openid)
	if err != nil {
		retCode = base.RedisCurdFailed
	}

	return
}

func DeleteHealthCard(openid string) (retCode *base.RetCode, err error) {
	var lockKey = "DeleteHealthCard_" + openid

	retCode, err = redis.Lock(lockKey)
	if retCode != base.SUCCESS {
		return
	}

	defer redis.Unlock(lockKey)

	err = mysql.DelHealthCardByOpenid(openid)

	if err != nil {
		retCode = base.RedisCurdFailed
	}
	return
}
