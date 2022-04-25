package helper

import (
	"strings"
)

type Response struct {
	Responses string      `json:responses`
	Status    bool        `json:status`
	Message   string      `json:message`
	Errors    interface{} `json:errors`
	Data      interface{} `json:data`
}

type EmptyObj struct{}

func BuildResponse(responses string, status bool, message string, data interface{}) Response {
	res := Response{
		Responses: responses,
		Status:    status,
		Message:   message,
		Errors:    nil,
		Data:      data,
	}
	return res
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}

// func NewUserArrayResponse(users []entity.User) []Response {
// 	usersRes := []Response{}
// 	return usersRes
// }
