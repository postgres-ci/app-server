package errors

import (
	"github.com/lib/pq"

	"database/sql"
	"net/http"
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) IsFatal() bool {

	return e.Code == http.StatusInternalServerError
}

func (e *Error) Error() string {

	return e.Message
}

func Wrap(err error) error {

	if err, ok := err.(*pq.Error); ok {

		switch err.Code.Name() {
		case "no_data_found":
			return &Error{
				Code:    http.StatusNotFound,
				Message: "Not found",
			}
		case "invalid_password":
			return &Error{
				Code:    http.StatusForbidden,
				Message: "Invalid password",
			}
		}

		switch err.Message {
		case "LOGIN_ALREADY_EXISTS":
			return &Error{
				Code:    http.StatusBadRequest,
				Message: "Login already exists",
			}
		case "EMAIL_ALREADY_EXISTS":
			return &Error{
				Code:    http.StatusBadRequest,
				Message: "Email already exists",
			}
		case "INVALID_EMAIL":
			return &Error{
				Code:    http.StatusBadRequest,
				Message: "Invalid email",
			}
		}
	}

	return &Error{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	}
}

func IsNotFound(err error) bool {

	if err == sql.ErrNoRows {

		return true
	}

	if err, ok := err.(*pq.Error); ok && err.Code.Name() == "no_data_found" {

		return true
	}

	return false
}

func IsInvalidPassword(err error) bool {

	if err, ok := err.(*pq.Error); ok && err.Code.Name() == "invalid_password" {

		return true
	}

	return false
}
