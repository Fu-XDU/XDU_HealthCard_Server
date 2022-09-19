package mysql

import "xdu-health-card/model"

func InsertHealthCard(healthCard *model.HealthCard) (ID int, err error) {
	err = db.Create(healthCard).Error
	ID = healthCard.ID
	return
}

func GetHealthCardByOpenid(openid string) (healthCard *model.HealthCard, err error) {
	registered, err := HealthCardRegisteredByOpenid(openid)
	if !registered {
		return &model.HealthCard{}, err
	}

	err = db.Where("openid = ?", openid).First(&healthCard).Error
	return
}

func DelHealthCardByOpenid(openid string) (err error) {
	err = db.Where("openid = ?", openid).Delete(model.HealthCard{}).Error
	return
}

func GetHealthCardByID(ID int) (healthCard *model.HealthCard, err error) {
	registered, err := HealthCardRegisteredByID(ID)
	if !registered {
		return &model.HealthCard{}, err
	}

	err = db.Where("id = ?", ID).First(&healthCard).Error
	return
}

func DelHealthCardByID(ID int) (err error) {
	err = db.Delete(model.HealthCard{}, ID).Error
	return
}

func GetNeededHealthCardByOpenid(openid string) (healthCard *model.HealthCard, err error) {
	registered, err := HealthCardRegisteredByOpenid(openid)
	if !registered {
		return &model.HealthCard{}, err
	}

	selectCol := []string{"id", "stu_id", "location"}
	err = db.Where("openid = ?", openid).Select(selectCol).First(&healthCard).Error
	return
}

func GetAllHealthCard() (healthCard []model.HealthCard, err error) {
	healthCard = []model.HealthCard{}
	err = db.Find(&healthCard).Error
	return
}

func HealthCardRegisteredByOpenid(openid string) (registered bool, err error) {
	count := int64(0)
	err = db.Model(&model.HealthCard{}).Where("openid = ?", openid).Count(&count).Error
	registered = count > 0
	return
}

func HealthCardRegisteredByID(ID int) (registered bool, err error) {
	count := int64(0)
	err = db.Model(&model.HealthCard{}).Where("id = ?", ID).Count(&count).Error
	registered = count > 0
	return
}
