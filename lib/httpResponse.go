package lib

// HTTPResponse type
type HTTPResponse struct {
	Code    int         `json:"code" xml:"code"`
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data,omitempty" xml:"data"`
}

// GetHTTPResponse returns HTTPResponse type
func GetHTTPResponse(code int, message string, data interface{}) HTTPResponse {
	return HTTPResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// GetHTTPInternalServerError returns 500 error
func GetHTTPInternalServerError(message string) HTTPResponse {
	if message == "" {
		message = "Internal Server Error"
	}

	return HTTPResponse{
		Code:    500,
		Message: message,
		Data:    nil,
	}
}
