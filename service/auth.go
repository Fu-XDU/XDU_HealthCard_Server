package service

import (
	"encoding/json"
	"fmt"
	"xdu-health-card/model"
	"xdu-health-card/model/base"
	"xdu-health-card/utils"
	"xdu-health-card/utils/flags"
)

func Login(code string) (res interface{}, err error) {
	response := code2Session(code)
	if response.
	return
}

func code2Session(code string) (response *base.Response) {
	response = base.NewResponse()
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", flags.AppID, flags.Secret, auth.Code)

	get, err := utils.Get(url)
	if err != nil {
		return base.NewErrorResponse(err, base.ConnectWxFailed)
	}

	var authRes *model.AuthResult
	_ = json.Unmarshal(get, &authRes)

	switch authRes.Errcode {
	case -1:
		return base.NewErrorResponse(err, base.WxServerBusy)
	case 40029:
		return base.NewErrorResponse(err, base.CodeInvalid)
	case 40226:
		return base.NewErrorResponse(err, base.HighRiskUser)
	case 45011:
		return base.NewErrorResponse(err, base.FrequencyLimit)
	}

	//_, err = redis.SetSessionID(authRes.Openid, authRes.SessionKey)
	//if err != nil {
	//	return base.NewErrorResponse(err, base.SetSessionIDFailed)
	//}

	response = base.NewDataResponse(authRes)
	return
}
