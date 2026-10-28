package response

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func ValidateErrors(errs validator.ValidationErrors) Response {
	var errMaps []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMaps = append(errMaps, fmt.Sprint("Required field ", err.Field()))
		case "url":
			errMaps = append(errMaps, fmt.Sprint("URL field ", err.Field()))
		default:
			errMaps = append(errMaps, fmt.Sprint(err))
		}
		return Response{
			Status: StatusError,
			Error:  strings.Join(errMaps, "; "),
		}
	}
	return Response{}
}
