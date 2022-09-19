package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"xdu-health-card/model"
	"xdu-health-card/model/base"
	"xdu-health-card/utils"
	"xdu-health-card/utils/flags"
)

type submitStatus struct {
	E int    `json:"e"`
	M string `json:"m"`
}

func LoginStuID(stuID, passwd string) (response *base.Response) {
	const loginURL = "https://xxcapp.xidian.edu.cn/uc/wap/login/check"
	var header = map[string]string{
		"Content-Type":     "application/x-www-form-urlencoded",
		"Accept":           "*/*",
		"Accept-Language":  "zh-cn",
		"Host":             "xxcapp.xidian.edu.cn",
		"Origin":           "https://xxcapp.xidian.edu.cn",
		"User-Agent":       "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1 Safari/605.1.15",
		"Referer":          "https://xxcapp.xidian.edu.cn/uc/wap/login?redirect=https%3A%2F%2Fxxcapp.xidian.edu.cn%2Fncov%2Fwap%2Fdefault%2Findex",
		"X-Requested-With": "XMLHttpRequest",
	}

	var data = url.Values{"username": {stuID}, "password": {passwd}}
	resp, body, err := utils.Post([]byte(data.Encode()), loginURL, header)
	if err != nil {
		return base.NewErrorResponse(err, base.LoginXidianFailed)
	}
	var submitRes submitStatus
	err = json.Unmarshal(body, &submitRes)
	if err != nil {
		return base.NewErrorResponse(err, base.LoginXidianFailed)
	}

	if submitRes.E != 0 {
		return base.NewErrorResponse(errors.New(submitRes.M), base.LoginXidianFailed)
	}

	cookies := resp.Cookies()
	cookiesStr := ""
	for i := range cookies {
		cookiesStr += cookies[i].Name + "=" + cookies[i].Value + ";"
	}
	return base.NewDataResponse(cookiesStr)
}

func SubmitThreeCheck(cookies, location string) (err error) {
	const submitURL = "https://xxcapp.xidian.edu.cn/xisuncov/wap/open-report/save"
	var header = map[string]string{
		"Accept":           "application/json, text/plain, */*",
		"Accept-Language":  "zh-cn",
		"Host":             "xxcapp.xidian.edu.cn",
		"Origin":           "https://xxcapp.xidian.edu.cn",
		"User-Agent":       "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1 Safari/605.1.15",
		"Connection":       "keep-alive",
		"Referer":          "https://xxcapp.xidian.edu.cn/site/ncov/xidiandailyup",
		"X-Requested-With": "XMLHttpRequest",
		"Cookie":           cookies,
	}

	data, _ := json.Marshal(model.LocationData[location])
	_, body, err := utils.Post(data, submitURL, header)
	if err != nil {
		return
	}
	var submitRes submitStatus
	err = json.Unmarshal(body, &submitRes)
	if err != nil {
		return
	}

	if submitRes.E != 0 {
		return errors.New(submitRes.M)
	}
	return
}

func SubmitHealthCard(cookies, location string) (err error) {
	const submitURL = "https://xxcapp.xidian.edu.cn/ncov/wap/default/save"
	var header = map[string]string{
		"Content-Type":              "application/x-www-form-urlencoded",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"Accept-Language":           "zh-CN,zh;q=0.9",
		"Host":                      "xxcapp.xidian.edu.cn",
		"Origin":                    "https://xxcapp.xidian.edu.cn",
		"User-Agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1 Safari/605.1.15",
		"Connection":                "keep-alive",
		"Referer":                   "https://xxcapp.xidian.edu.cn/uc/wap/login?redirect=https%3A%2F%2Fxxcapp.xidian.edu.cn%2Fncov%2Fwap%2Fdefault%2Findex",
		"X-Requested-With":          "XMLHttpRequest",
		"Sec-Fetch-Dest":            "document",
		"Sec-Fetch-Mode":            "navigate",
		"Sec-Fetch-Site":            "same-origin",
		"Sec-Fetch-User":            "?1",
		"Upgrade-Insecure-Requests": "1",
		"Cookie":                    cookies,
	}
	var locationStruct model.HealthCardLocation
	err = json.Unmarshal([]byte(location), &locationStruct)
	if err != nil {
		return
	}
	geoApiInfo, _ := json.Marshal(locationStruct.GeoApiInfo)
	data := url.Values{
		"MIME类型":       {"application/x-www-form-urlencoded; charset=UTF-8"},
		"szgjcs":       {""},  // 所在国家城市
		"szcs":         {""},  // 所在城市
		"szgj":         {""},  // 所在国家
		"zgfxdq":       {"0"}, // 今日是否在中高风险地区？（中高风险地区信息可通过国务院客户端小程序实时查询）
		"mjry":         {"0"}, // 是否接触密接人员
		"csmjry":       {"0"}, // 近14日内本人/共同居住者是否去过疫情发生场所（市场、单位、小区等）或与场所人员有过密切接触
		"tw":           {"2"}, // 体温：第三项，36-36.5
		"sfcxtz":       {"0"}, // 是否出现发热（37.3℃以上）、乏力、干咳、呼吸困难等任意症状之一
		"sfjcbh":       {"0"}, // 是否接触无症状感染/疑似/确诊人群
		"sfcxzysx":     {"0"}, // 是否有任何与疫情相关的， 值得注意的情况
		"qksm":         {"0"},
		"sfyyjc":       {"0"}, // 是否医院检查
		"jcjgqr":       {"0"}, // 检查结果确认
		"remark":       {""},
		"address":      {locationStruct.Address},  // 地址
		"geo_api_info": {string(geoApiInfo)},      // 定位系统详情
		"area":         {locationStruct.Area},     // 地区
		"province":     {locationStruct.Province}, // 省份
		"city":         {locationStruct.City},     // 城市
		"sfzx":         {"0"},                     // 是否在校 否
		"sfjcwhry":     {"0"},                     // 是否接触武汉人员 否
		"sfjchbry":     {"0"},                     // 是否接触湖北人员 否
		"sfcyglq":      {"0"},                     // 是否处于隔离期 否
		"gllx":         {""},
		"glksrq":       {""},
		"jcbhlx":       {""},
		"jcbhrq":       {""},
		"ismoved":      {"0"}, // 与上次地点是否有不同
		"bztcyy":       {""},
		"sftjhb":       {"0"}, // 是否途径湖北 否
		"sftjwh":       {"0"}, // 是否途径武汉 否
		"sfjcjwry":     {"0"}, // 是否接触境外人员 否
		"jcjg":         {""},
	}
	_, body, err := utils.Post([]byte(data.Encode()), submitURL, header)
	if err != nil {
		return
	}
	var submitRes submitStatus
	err = json.Unmarshal(body, &submitRes)
	if err != nil {
		return
	}

	if submitRes.E != 0 {
		return
	}
	return
}

func ConvertHealthCardLocation(latitude, longitude float32) (location string, err error) {
	var mapApiUrl = fmt.Sprintf("https://apis.map.qq.com/ws/geocoder/v1/?location=%v,%v&key=%s", latitude, longitude, flags.MapApiKey)
	_, body, err := utils.Get(mapApiUrl)
	if err != nil {
		return
	}
	var convertResult model.ConvertLocationResult
	err = json.Unmarshal(body, &convertResult)
	if err != nil {
		return
	}

	if convertResult.Status != 0 {
		err = errors.New(convertResult.Message)
		return
	}

	locationResult := convertResult.ToHealthCardLocation()
	locationBytes, err := json.Marshal(locationResult)
	location = string(locationBytes)
	return
}
