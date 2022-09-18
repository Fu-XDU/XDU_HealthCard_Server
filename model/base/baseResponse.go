package base

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

func NewResponse() *Response {
	return NewDataResponse(nil)
}

func NewDataResponse(data interface{}) *Response {
	return &Response{
		Code:    SUCCESS.Code,
		Message: SUCCESS.Message,
		Data:    data,
		Success: true,
	}
}

func NewErrorResponse(err error, retCode *RetCode) *Response {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}

	return &Response{
		Code:    retCode.Code,
		Message: retCode.Message,
		Error:   errMsg,
		Success: false,
	}
}
