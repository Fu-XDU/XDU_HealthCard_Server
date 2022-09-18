package base

type PageResponse struct {
	Code        int         `json:"code"`
	Message     string      `json:"message"`
	Error       string      `json:"error"`
	Success     bool        `json:"success"`
	Data        interface{} `json:"data"`
	CurrentPage int         `json:"current"`
	PageSize    int         `json:"pageSize"`
	Total       int64       `json:"total"`
}

func NewPageResponse(currentPage, pageSize int, total int64) *PageResponse {
	return NewDataPageResponse(nil, currentPage, pageSize, total)
}

func NewDataPageResponse(data interface{}, currentPage, pageSize int, total int64) *PageResponse {
	return &PageResponse{
		Code:        SUCCESS.Code,
		Message:     SUCCESS.Message,
		Data:        data,
		CurrentPage: currentPage,
		PageSize:    pageSize,
		Total:       total,
		Success:     true,
	}
}

func NewErrorPageResponse(err error, retCode *RetCode, currentPage, pageSize int, total int64) *PageResponse {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}

	return &PageResponse{
		Code:        retCode.Code,
		Message:     retCode.Message,
		Error:       errMsg,
		CurrentPage: currentPage,
		PageSize:    pageSize,
		Total:       total,
		Success:     false,
	}
}
