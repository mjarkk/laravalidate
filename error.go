package laravalidate

import (
	"encoding/json"
)

type ValidationError struct {
	Mode     Mode          `json:"mode"`
	Language []string      `json:"language"`
	Errors   []FieldErrors `json:"errors"`
}

func (*ValidationError) Error() string {
	return "Validation Error"
}

type FieldErrors struct {
	Path     string                `json:"goPath"`
	JsonPath string                `json:"jsonPath"`
	Errors   []FieldValidatorError `json:"errors"`
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
		if e.Mode == JsonMode {
			path = entry.JsonPath
		}

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
