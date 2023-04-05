package helper

import (
	"strings"
)

type Response struct {
	StatusCode int         `json:statusCode`
	Success    bool        `json:success`
	Message    string      `json:message`
	Failed     bool        `json:failed`
	Data       interface{} `json:data`
}

type ErrorResponse struct {
	Success bool        `json:success`
	Message string      `json:message`
	Data    interface{} `json:data`
	Error   interface{} `json:error`
}

type EmptyObj struct{}

func BuildResponse(statusCode int, success bool, message string, failed bool, data interface{}) Response {
	res := Response{
		StatusCode: statusCode,
		Success:    success,
		Message:    message,
		Failed:     failed,
		Data:       data,
	}
	return res
}

func BuildErrorResponse(message string, err string, data interface{}) ErrorResponse {
	splitError := strings.Split(err, "\n")
	res := ErrorResponse{
		Success: false,
		Message: message,
		Error:   splitError,
		Data:    data,
	}
	return res
}
