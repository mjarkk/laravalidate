package laravalidate

import "sort"

func CreateGoError(errors map[string]string) error {
	return createCustomError(GoMode, errors)
}

func CreateJsonError(errors map[string]string) error {
	return createCustomError(JsonMode, errors)
}

func CreateFormError(errors map[string]string) error {
	return createCustomError(FormMode, errors)
}

func createCustomError(mode Mode, errors map[string]string) error {
	fieldErrors := []FieldErrors{}
	for key, message := range errors {
		fieldErrors = append(fieldErrors, FieldErrors{
			Path: key,
			Errors: []FieldValidatorError{{
				Rule:    "custom",
				Message: message,
			}},
		})
	}

	sort.Slice(fieldErrors, func(i, j int) bool {
		return fieldErrors[i].Path < fieldErrors[j].Path
	})

	return &ValidationError{
		Mode:     mode,
		Language: nil, // Maybe we should do something with this?
		Errors:   fieldErrors,
	}
}
