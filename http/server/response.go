package server

// Response defines the server response
type Response struct {
	Data         interface{} `json:"data,omitempty"`
	ErrorCode    string      `json:"errorCode,omitempty"`
	ErrorMessage string      `json:"errorMessage,omitempty"`
}

// ResponseFromData returns a successful response without error
func ResponseFromData(data interface{}) *Response {
	return &Response{
		Data: data,
	}
}

// ErrorResponse return an response when server errored
func ErrorResponse(err error, code string) *Response {
	return &Response{
		ErrorMessage: err.Error(),
		ErrorCode:    code,
	}
}
