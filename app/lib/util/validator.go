package util

import (
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

// ErrMsg maps error names to error values
type ErrMsg map[string]string

// Error implements the error interface
// and sends a concatenation of all ErrMsg values seperated
// by a newline

func (e ErrMsg) Error() string {

	msg := ""
	for _, v := range e {
		msg += v
		msg += "\n"
	}

	return msg
}

// CustomValidator for request body
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate is called by the echo instance after binding
// the response body to a the validation object i
// it is expected that the interface i be a struct

// Validate ...
func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.Validator.Struct(i)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		errmsg := ErrMsg{}
		for _, err := range err.(validator.ValidationErrors) {
			key := strings.ToLower(err.Field())
			errmsg[key] = fmt.Sprintf("There is an error with '%s' field", err.Field())

		}
		return errmsg
	}
	return nil
}

// Name returns the name of the validator
func (cv *CustomValidator) Name() string {
	return "A Custom Validator"
}
