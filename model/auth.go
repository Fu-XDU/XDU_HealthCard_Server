package model

type AuthRequest struct {
	Code string `json:"code"`
}

type AuthResult struct {
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}
