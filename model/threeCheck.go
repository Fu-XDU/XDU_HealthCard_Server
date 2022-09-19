package model

type ThreeCheck struct {
	ID       int    `json:"id"`
	Openid   string `gorm:"type:varchar(40);" json:"openid,omitempty"`
	StuID    string `gorm:"type:varchar(20);" json:"stu_id"`
	Passwd   string `json:"passwd,omitempty"`
	Location string `json:"location"`
}

var LocationData = map[string]map[string]string{
	"西电南校区": {
		"MIME类型":       "application/json; charset=UTF-8",
		"sfzx":         "1", // 是否在校 是
		"tw":           "1", // 体温 36-36.5
		"area":         "陕西省 西安市 长安区",
		"city":         "西安市",
		"province":     "陕西省",
		"address":      "陕西省西安市长安区郭杜街道雷甘路西安电子科技大学长安校区",
		"geo_api_info": `{"type": "complete","position": {"Q": 34.130520019532, "R": 108.83373643663197, "lng": 108.833736, "lat": 34.13052},"location_type": "html5","message": "Get ipLocation failed.Get geolocation success.Convert Success.Get address success.","accuracy": 165, "isConverted": true, "status": 1,"addressComponent": {"citycode": "029", "adcode": "610116", "businessAreas": [],"neighborhoodType": "", "neighborhood": "", "building": "","buildingType": "", "street": "海棠二路", "streetNumber": "252号","country": "中国", "province": "陕西省", "city": "西安市", "district": "长安区","township": "郭杜街道"}, "formattedAddress": "陕西省西安市长安区郭杜街道雷甘路西安电子科技大学长安校区","roads": [], "crosses": [], "pois": [], "info": "SUCCESS"}`,
		"sfcyglq":      "0", // 是否处于隔离期 否
		"sfyzz":        "0", // 是否出现发力干咳呼吸困难等症状 否
		"qtqk":         "",  // 其他情况 不填
		"ymtys":        "0", //一码通颜色 绿
	}, "西电北校区": {
		"MIME类型":       "application/json; charset=UTF-8",
		"sfzx":         "1", // 是否在校 是
		"tw":           "1", // 体温 36-36.5
		"area":         "陕西省 西安市 雁塔区",
		"city":         "西安市",
		"province":     "陕西省",
		"address":      "陕西省西安市雁塔区白沙路9号西安电子科技大学北校区",
		"geo_api_info": `{"type": "complete","position": {"Q": 34.232548, "R": 108.914364, "lng": 108.914364, "lat": 34.232548},"location_type": "html5","message": "Get ipLocation failed.Get geolocation success.Convert Success.Get address success.","accuracy": 165, "isConverted": true, "status": 1,"addressComponent": {"citycode": "029", "adcode": "610116", "businessAreas": [],"neighborhoodType": "", "neighborhood": "", "building": "","buildingType": "", "street": "白沙路", "streetNumber": "9号","country": "中国", "province": "陕西省", "city": "西安市", "district": "雁塔区","township": "白沙路9号"}, "formattedAddress": "陕西省西安市雁塔区白沙路9号西安电子科技大学北校区","roads": [], "crosses": [], "pois": [], "info": "SUCCESS"}`,
		"sfcyglq":      "0", // 是否处于隔离期 否
		"sfyzz":        "0", // 是否出现发力干咳呼吸困难等症状 否
		"qtqk":         "",  // 其他情况 不填
		"ymtys":        "0", //一码通颜色 绿
	},
}
