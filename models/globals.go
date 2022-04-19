package models

type Result struct {
	Data       interface{}
	StatusCode int
	Error      error
	Message    string
}

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (r *Result) ToResponse() *Response {
	resp := Response{
		Code: r.StatusCode,
	}
	if r.Data != nil {
		resp.Data = r.Data
	} else {
		resp.Data = make(map[string]interface{})
	}
	resp.Message = r.Message
	return &resp
}

func (r *Result) ToResponseError(statusCode int) *Response {
	var message string
	if r.Message != "" {
		message = r.Message
	} else {
		message = r.Error.Error()
	}
	return &Response{
		Code:    statusCode,
		Data:    make(map[string]interface{}),
		Message: message,
	}
}
