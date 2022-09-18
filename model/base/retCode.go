package base

var (
	SUCCESS = NewRetCode(0, "Success")

	BindDataFailed = NewRetCode(10001, "Bind data failed")

	GetWeatherInfoFailed = NewRetCode(30001, "Get weather info from database failed")
	InsertDataFailed     = NewRetCode(30002, "Insert data to database failed")
	UploadDataFailed     = NewRetCode(30003, "Upload warehouse item failed")
	ConnectWxFailed      = NewRetCode(30004, "Connect wx server failed")
	SetSessionIDFailed   = NewRetCode(30005, "Set sessionID failed")
	WxServerBusy         = NewRetCode(30006, "System is busy, please try again later")
	CodeInvalid          = NewRetCode(30007, "Code invalid")
	FrequencyLimit       = NewRetCode(30008, "Frequency limit, 100 times per minute per user")
	HighRiskUser         = NewRetCode(30009, "High-risk user, login interception")
	WrongSessionKey      = NewRetCode(30010, "Wrong session key")
)

type RetCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewRetCode(code int, message string) *RetCode {
	return &RetCode{Code: code, Message: message}
}
