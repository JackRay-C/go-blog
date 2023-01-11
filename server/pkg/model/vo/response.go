package vo

import "encoding/json"

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

func Message(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
	}
}

func Failed(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
	}
}
