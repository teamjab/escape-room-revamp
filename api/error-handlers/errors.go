package errors

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	// 500 Error
	InternalErr = &Error{
		Code: http.StatusInternalServerError,
		Message: "Something went wrong on our side, but you probabaly broke it.",
	}

	// 422 Error
	EntityUnprocessabeErr = &Error{
		Code: http.StatusUnprocessableEntity,
		Message: "Unprocessable Entity",
	}

	// 404 Error
	NotFoundErr = &Error{
		Code: http.StatusNotFound,
		Message: "Not found!",
	}

	// 400 Error 
	ObjectNotFoundErr = &Error{
		Code: http.StatusBadRequest,
		Message: "Request should provide entity",
	}
)

type Error struct {
	Code int // this is status code,
	Message string // this is error message
}

func (err *Error) Error() string {
	return err.Error()
}

func (err *Error) CustomError() string {
	if err == nil {
		return ""
	}
	log.Fatalf("error: Code: %s Message: %s", http.StatusText(err.Code), err.Message)
	return fmt.Sprintf("error: Code: %s Message: %s", http.StatusText(err.Code), err.Message)
}

// Convert Error json object
func (err *Error) JSON() []byte {
	if err == nil {
		return []byte("{}")
	}

	res, _ := json.Marshal(err)
	return res
}

func (err *Error) StatusCode() int {
	if err == nil {
		return http.StatusOK
	}

	return err.Code
}