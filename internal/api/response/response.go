package response

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOk      = "OK"
	StatusError   = "ERROR"
	ValidateError = "validation error:"
)

func OK() Response {
	return Response{
		Status: StatusOk,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("%s failed %s is a required field", ValidateError, err.Field()))
		case "url":
			errMsgs = append(errMsgs, fmt.Sprintf("%s failed %s is a value URL", ValidateError, err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("%s failed %s is not valid", ValidateError, err.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}
}
