package lib

// HTTPResponse type
type HTTPResponse struct {
	Code    int         `json:"code" xml:"code"`
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data,omitempty" xml:"data"`
}

// GetHTTPResponse : Retourne le type HTTPResponse
func GetHTTPResponse(code int, message string, data interface{}) HTTPResponse {
	return HTTPResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// GetHTTPInternalServerError return 500 error
func GetHTTPInternalServerError() HTTPResponse {
	return HTTPResponse{
		Code:    500,
		Message: "Internal Server Error",
		Data:    nil,
	}
}
