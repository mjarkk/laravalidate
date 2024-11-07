package laravalidate

import (
	"reflect"
)

type CustomError struct {
	Key      string
	Resolver MessageResolver
}

var customErrorType = reflect.TypeOf(CustomError{})

func (v *Validator) CustomValidationRules() []CustomError {
	if v.customValidationMessagesCache != nil {
		return v.customValidationMessagesCache
	}

	validationMessagesMethod := v.inputValue.MethodByName("ValidationMessages")
	if !validationMessagesMethod.IsValid() {
		return nil
	}

	if !validatorMethodValid(validationMessagesMethod.Type()) {
		return nil
	}

	customValidationMessages := validationMessagesMethod.Call([]reflect.Value{})
	customMsg := customValidationMessages[0].Interface().([]CustomError)
	if len(customMsg) == 0 {
		return nil
	}

	v.customValidationMessagesCache = customMsg
	return v.customValidationMessagesCache
}

func validatorMethodValid(fnType reflect.Type) bool {
	if fnType.NumIn() != 0 {
		return false
	}
	if fnType.NumOut() != 1 {
		return false
	}

	outType := fnType.Out(0)
	if outType.Kind() != reflect.Slice {
		return false
	}

	return outType.Elem().String() == customErrorType.String()
}

func (v *Validator) CustomValidationRule(ruleName string, stack Stack) MessageResolver {
	customErrors := v.CustomValidationRules()
	if customErrors == nil {
		return nil
	}

	for _, err := range customErrors {
		if stack.LooslyEqualsWithRule(err.Key, ruleName) {
			return err.Resolver
		}
	}

	return nil
}
