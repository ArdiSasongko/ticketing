package helper

// ResponseClient is a standard response structure for API responses
type ResponseClientModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewResponseClient creates a new ResponseClient instance
func ResponseClient(code int, message string, data interface{}) ResponseClientModel {
	return ResponseClientModel{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
