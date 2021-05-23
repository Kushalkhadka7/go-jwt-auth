package util

import (
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Response struct {
	code    codes.Code
	message string
	err     error
}

type Responser interface {
	Error() error
}

func NewResponse(code codes.Code, message string, err error) Responser {
	return &Response{
		code:    code,
		message: message,
		err:     err,
	}
}

// Error log the error to console and return it with status code.
func (res *Response) Error() error {
	log.Printf("%s: [%v] : %v", res.message, res.code, res.err)

	formattedMessage := res.message + ":%v"

	return status.Errorf(res.code, formattedMessage, res.err)
}
