package base

var (
	SUCCESS = NewRetCode(0, "Success")

	BindDataFailed = NewRetCode(10001, "Bind data failed")

	NewJwtFailed       = NewRetCode(30001, "New jwt failed")
	Unauthorized       = NewRetCode(30002, "Unauthorized")
	LoginXidianFailed  = NewRetCode(30003, "Login xidian failed")
	ConnectWxFailed    = NewRetCode(30004, "Connect wx server failed")
	SetRedisLockFailed = NewRetCode(30005, "Set redis lock failed")
	WxServerBusy       = NewRetCode(30006, "System is busy, please try again later")
	CodeInvalid        = NewRetCode(30007, "Code invalid")
	FrequencyLimit     = NewRetCode(30008, "Frequency limit, 100 times per minute per user")
	HighRiskUser       = NewRetCode(30009, "High-risk user, login interception")
	LockExist          = NewRetCode(30010, "Redis lock has existed")
	CannotSubmitAgain  = NewRetCode(30011, "Cannot submit again")
	RedisCurdFailed    = NewRetCode(30012, "Redis curd failed")
)

type RetCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewRetCode(code int, message string) *RetCode {
	return &RetCode{Code: code, Message: message}
}
