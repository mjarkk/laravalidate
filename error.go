package laravalidate

import (
	"encoding/json"
)

type ValidationError struct {
	Mode     Mode          `json:"mode"`
	Language []string      `json:"language"`
	Errors   []FieldErrors `json:"errors"`
}

const fallbackErrorMessage = "Validation Error"

func (v *ValidationError) Error() string {
	for _, err := range v.Errors {
		for _, field := range err.Errors {
			return field.Message
		}
	}

	return fallbackErrorMessage
}

type FieldErrors struct {
	Path   string                `json:"path"`
	Errors []FieldValidatorError `json:"errors"`
}

type FieldValidatorError struct {
	Rule    string `json:"rule"`
	Hint    string `json:"hint"`
	Message string `json:"message"`
}

type LaravelValidationError struct {
	Errors  LaravelErrorObj `json:"errors"`
	Message string          `json:"message"`
}

type LaravelErrorObjEntry struct {
	Key   string
	Value []string
}

type LaravelErrorObj []LaravelErrorObjEntry

func (obj LaravelErrorObj) MarshalJSON() ([]byte, error) {
	resp := []byte{'{'}
	for i, entry := range obj {
		if i != 0 {
			resp = append(resp, ',')
		}

		resp = append(resp, '"')
		resp = append(resp, []byte(entry.Key)...)
		resp = append(resp, '"', ':')

		b, err := json.Marshal(entry.Value)
		if err != nil {
			return nil, err
		}

		resp = append(resp, b...)
	}
	resp = append(resp, '}')
	return resp, nil
}

func (e *ValidationError) ToLaravelError() *LaravelValidationError {
	errors := LaravelErrorObj{}

	for _, entry := range e.Errors {
		path := entry.Path

		errorMessages := []string{}
		for _, validatorError := range entry.Errors {
			errorMessages = append(errorMessages, validatorError.Message)
		}
		errors = append(errors, LaravelErrorObjEntry{
			Key:   path,
			Value: errorMessages,
		})
	}

	return &LaravelValidationError{
		Errors:  errors,
		Message: "Form contains errors",
	}
}
