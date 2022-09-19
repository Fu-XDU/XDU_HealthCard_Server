package mysql

import "xdu-health-card/model"

func InsertThreeCheck(threeCheck *model.ThreeCheck) (ID int, err error) {
	err = db.Create(threeCheck).Error
	ID = threeCheck.ID
	return
}

func GetThreeCheckByOpenid(openid string) (threeCheck *model.ThreeCheck, err error) {
	registered, err := ThreeCheckRegisteredByOpenid(openid)
	if !registered {
		return &model.ThreeCheck{}, err
	}

	err = db.Where("openid = ?", openid).First(&threeCheck).Error
	return
}

func DelThreeCheckByOpenid(openid string) (err error) {
	err = db.Where("openid = ?", openid).Delete(model.ThreeCheck{}).Error
	return
}

func GetThreeCheckByID(ID int) (threeCheck *model.ThreeCheck, err error) {
	registered, err := ThreeCheckRegisteredByID(ID)
	if !registered {
		return &model.ThreeCheck{}, err
	}

	err = db.Where("id = ?", ID).First(&threeCheck).Error
	return
}

func DelThreeCheckByID(ID int) (err error) {
	err = db.Delete(model.ThreeCheck{}, ID).Error
	return
}

func GetNeededThreeCheckByOpenid(openid string) (threeCheck *model.ThreeCheck, err error) {
	registered, err := ThreeCheckRegisteredByOpenid(openid)
	if !registered {
		return &model.ThreeCheck{}, err
	}

	selectCol := []string{"id", "stu_id", "location"}
	err = db.Where("openid = ?", openid).Select(selectCol).First(&threeCheck).Error
	return
}

func GetAllThreeCheck() (threeCheck []model.ThreeCheck, err error) {
	threeCheck = []model.ThreeCheck{}
	err = db.Find(&threeCheck).Error
	return
}

func ThreeCheckRegisteredByOpenid(openid string) (registered bool, err error) {
	count := int64(0)
	err = db.Model(&model.ThreeCheck{}).Where("openid = ?", openid).Count(&count).Error
	registered = count > 0
	return
}

func ThreeCheckRegisteredByID(ID int) (registered bool, err error) {
	count := int64(0)
	err = db.Model(&model.ThreeCheck{}).Where("id = ?", ID).Count(&count).Error
	registered = count > 0
	return
}
