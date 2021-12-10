package response

import (
	"blog/app/pager"
	"encoding/json"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Response) String() string {
	marshal, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func Success(data interface{}) *Response {
	return &Response{
		Code:    200,
		Message: "Success",
		Data:    data,
	}
}


func PagerResponse(pager *pager.Pager) *Response {
	return &Response{
		Code:    200,
		Message: "Success",
		Data:    pager,
	}
}
