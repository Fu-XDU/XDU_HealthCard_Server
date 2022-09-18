package base

import "fmt"

type PageRequest struct {
	PageSize    int `form:"pageSize"`
	CurrentPage int `form:"current"`
}

func (w *PageRequest) Validate() (err error) {
	if w.CurrentPage <= 0 {
		err = fmt.Errorf("invalid current: %v", w.CurrentPage)
	}

	if w.PageSize <= 0 {
		err = fmt.Errorf("invalid pageSize: %v", w.PageSize)
	}
	return
}
