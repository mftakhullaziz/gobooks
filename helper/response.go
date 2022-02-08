package helper

type Response struct {
	Status		bool 	`json:status`
	message		string	`json:message`
	Error		interface{}	`json:errors`
	Data 		interface{}	`json:data`
}

type EmptyObj struct {

}

func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response {
		Status: status,
		Message: message,
		Errors:	nil,
		Data: data,
	}
	return res
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := string.Split(err, "\n")
	res := {
		Status: false,
		Message: message,
		Errors: splittedError,
		Data: data,
	}
	return res
}