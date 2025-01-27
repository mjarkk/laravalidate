package laravalidate

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/language"
)

type Validator struct {
	inputValue reflect.Value
	ctx        context.Context
	errors     []FieldErrors
	languages  []string
	mode       Mode
	// Cache
	customValidationMessagesCache []CustomError
}

func newValidator(ctx context.Context, languages []language.Tag, value reflect.Value, mode Mode) *Validator {
	return &Validator{
		inputValue: value,
		ctx:        ctx,
		errors:     []FieldErrors{},
		languages:  lookupLanguages(languages),
		mode:       mode,
	}
}

type Mode uint8

const (
	GoMode   Mode = iota
	JsonMode      // Tries to use json struct tags
	FormMode      // Tries to use form struct tags
)

// JsonValidate should be used to validate a json parsed message, errors returned will have a json paths.
// These look at the `json="xx"` struct tag for hints how to name the error keys.
//
// If an error is returned the type should be of *ValidationError.
//
// Ctx can be set to nil, default value will be context.Background().
// Languages can be set to nil, default value will be []language.Tag{language.English}.
func JsonValidate(ctx context.Context, languages []language.Tag, input any) error {
	return validate(ctx, languages, input, JsonMode)
}

// FormValidate should be used to validate a form parsed message, errors returned will have a form paths.
// These look at the `form="xx"` struct tag for hints how to name the error keys.
//
// If an error is returned the type should be of *ValidationError.
//
// Ctx can be set to nil, default value will be context.Background().
// Languages can be set to nil, default value will be []language.Tag{language.English}.
func FormValidate(ctx context.Context, languages []language.Tag, input any) error {
	return validate(ctx, languages, input, FormMode)
}

// GoValidate should be used to validate something within a go codebase with validation errors that apply to the go codebase.
//
// If an error is returned the type should be of *ValidationError.
//
// Ctx can be set to nil, default value will be context.Background().
// Languages can be set to nil, default value will be []language.Tag{language.English}.
func GoValidate(ctx context.Context, languages []language.Tag, input any) error {
	return validate(ctx, languages, input, GoMode)
}

func validate(ctx context.Context, languages []language.Tag, input any, mode Mode) error {
	value := reflect.ValueOf(input)

	if ctx == nil {
		ctx = context.Background()
	}
	v := newValidator(ctx, languages, value, mode)

	for value.Kind() == reflect.Ptr {
		if !value.IsNil() {
			value = value.Elem()
			continue
		}

		v.Nil(Stack{}, value.Type().Elem())
		return v.Error()
	}

	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		v.List(Stack{}, value, nil)
	case reflect.Struct:
		v.Struct(Stack{}, value)
	default:
		return nil
	}

	return v.Error()
}

func (v *Validator) Error() error {
	if len(v.errors) == 0 {
		return nil
	}

	return &ValidationError{
		Mode:     v.mode,
		Language: v.languages,
		Errors:   v.errors,
	}
}

func (v *Validator) Nil(stack Stack, valueType reflect.Type) {
	if len(stack) > 100 {
		return
	}

	for valueType.Kind() == reflect.Ptr {
		valueType = valueType.Elem()
	}

	if valueType.Kind() != reflect.Struct {
		return
	}

	v.NilStruct(stack, valueType)
}

func (v *Validator) List(stack Stack, value reflect.Value, validateInner []validationRule) {
	if value.IsNil() {
		return
	}
	if len(stack) > 100 {
		return
	}

	var innerStack Stack
	var element reflect.Value
outer:
	for idx := 0; idx < value.Len(); idx++ {
		element = value.Index(idx)
		innerStack = stack.AppendIndex(idx, &value, value.Type())

		v.Validate(innerStack, &element, element.Type(), validateInner)

		for element.Kind() == reflect.Ptr {
			if element.IsNil() {
				v.Nil(innerStack, element.Type().Elem())
				continue outer
			}

			element = element.Elem()
		}

		switch element.Kind() {
		case reflect.Struct:
			v.Struct(innerStack, element)
		case reflect.Slice, reflect.Array:
			v.List(innerStack, element, nil)
		}
	}
}

func (v *Validator) Struct(stack Stack, value reflect.Value) {
	if len(stack) > 100 {
		return
	}

	var fieldType reflect.StructField
	var field reflect.Value
	var innerStack Stack
outer:
	for idx := 0; idx < value.NumField(); idx++ {
		fieldType = value.Type().Field(idx)
		field = value.Field(idx)

		innerStack = stack.AppendField(fieldType, &value, value.Type())

		validate := validationRules(fieldType.Tag.Get("validate"))
		validateInner := validationRules(fieldType.Tag.Get("validateInner"))
		if len(validate) > 0 || len(validateInner) > 0 {
			v.Validate(innerStack, &field, fieldType.Type, validate)
		}

		for field.Kind() == reflect.Ptr {
			if field.IsNil() {
				v.Nil(innerStack, field.Type().Elem())
				continue outer
			}
			field = field.Elem()
		}

		switch field.Kind() {
		case reflect.Struct:
			v.Struct(innerStack, field)
		case reflect.Slice, reflect.Array:
			v.List(innerStack, field, validateInner)
		}
	}
}

func (v *Validator) NilStruct(stack Stack, valueType reflect.Type) {
	if len(stack) > 100 {
		return
	}

	for idx := 0; idx < valueType.NumField(); idx++ {
		fieldType := valueType.Field(idx)
		innerStack := stack.AppendField(fieldType, nil, valueType)

		validate := validationRules(fieldType.Tag.Get("validate"))
		validateInner := validationRules(fieldType.Tag.Get("validateInner"))
		if len(validate) == 0 && len(validateInner) == 0 {
			continue
		}

		v.Validate(innerStack, nil, fieldType.Type, validate)

		for fieldType.Type.Kind() == reflect.Ptr {
			fieldType.Type = fieldType.Type.Elem()
		}

		switch fieldType.Type.Kind() {
		case reflect.Struct:
			v.NilStruct(innerStack, fieldType.Type)
		}
	}
}

func (v *Validator) Validate(stack Stack, value *reflect.Value, valueType reflect.Type, rules []validationRule) {
	if len(rules) == 0 {
		return
	}

	errors := []FieldValidatorError{}
	state := &ValidatorCtxState{
		bail:      false,
		state:     map[string]any{},
		stack:     stack,
		validator: v,
	}
	for _, rule := range rules {
		ctx := &ValidatorCtx{
			ctx:   v.ctx,
			Args:  rule.args,
			state: state,
			Needle: Needle{
				Value: value,
				Type:  valueType,
			},
		}
		hint, ok := rule.validator.Fn(ctx)
		if ok {
			continue
		}

		errors = append(errors, FieldValidatorError{
			Rule:    rule.name,
			Hint:    hint,
			Message: v.ErrorMessage(rule.name, rule.validator.Messages, hint, ctx),
		})
		if state.bail {
			break
		}
	}

	if len(errors) == 0 {
		return
	}

	goPath, jsonPath, formPath := stack.ToPaths()
	path := goPath
	if v.mode == JsonMode {
		path = jsonPath
	} else if v.mode == FormMode {
		path = formPath
	}

	v.errors = append(v.errors, FieldErrors{
		Path:   path,
		Errors: errors,
	})
}

func (v *Validator) ErrorMessage(ruleName string, resolvers map[string]MessageResolver, hint string, ctx *ValidatorCtx) string {
	template := v.ErrorMessageTemplate(ruleName, resolvers, hint, ctx.state.stack)

	replaceVariable := func(location templateVariableT, a string) {
		template = template[:location.from] + a + template[location.to:]
	}

	variables := parseMsgTemplate([]byte(template))

outer:
	for idx := len(variables) - 1; idx >= 0; idx-- {
		variable := variables[idx]
		variableName := template[variable.from:variable.to]
		switch variableName[1:] {
		case "attribute":
			stack := ctx.state.stack
			if len(stack) == 0 {
				replaceVariable(variable, "")
				continue outer
			}

			stackElement := stack[len(stack)-1]
			if v.mode == JsonMode {
				replaceVariable(variable, stackElement.JsonName)
			} else if v.mode == FormMode {
				replaceVariable(variable, stackElement.FormName)
			} else {
				replaceVariable(variable, stackElement.GoName)
			}
			continue outer
		case "other":
			if ctx.lastObtainedField == nil || !ctx.lastObtainedField.HasValue() {
				replaceVariable(variable, "")
				continue outer
			}

			if v.mode == JsonMode {
				jsonValue, err := json.Marshal(ctx.lastObtainedField.Value.Interface())
				if err == nil {
					replaceVariable(variable, string(jsonValue))
					continue outer
				}
			}

			replaceVariable(variable, fmt.Sprintf("%+v", ctx.lastObtainedField.Value.Interface()))
			continue outer
		case "value":
			if !ctx.HasValue() {
				replaceVariable(variable, "")
				continue outer
			}

			if v.mode == JsonMode {
				jsonValue, err := json.Marshal(ctx.Value.Interface())
				if err == nil {
					replaceVariable(variable, string(jsonValue))
					continue outer
				}
			}

			replaceVariable(variable, fmt.Sprintf("%+v", ctx.Value.Interface()))
			continue outer
		case "date":
			t, ok := ctx.DateFromArgs(0)
			if !ok {
				replaceVariable(variable, "")
				continue outer
			}

			replaceVariable(variable, t.Format(time.DateTime))
			continue outer
		}

		if strings.HasPrefix(variableName[1:], "arg") {
			if variableName == ":args" {
				replaceVariable(variable, strings.Join(ctx.Args, ", "))
			} else if variableName == ":arg" {
				if len(ctx.Args) == 0 {
					replaceVariable(variable, "")
					continue
				}

				replaceVariable(variable, ctx.Args[0])
			} else {
				suffix := variableName[4:]
				idx, err := strconv.Atoi(suffix)
				if err != nil && idx < 0 {
					continue
				}

				if idx >= len(ctx.Args) {
					replaceVariable(variable, "")
					continue
				}

				replaceVariable(variable, ctx.Args[idx])
			}
		}
	}

	return template
}

func (v *Validator) ErrorMessageTemplate(ruleName string, resolvers map[string]MessageResolver, hint string, stack Stack) string {
	customResolver := v.CustomValidationRule(ruleName, stack)
	if customResolver != nil {
		return customResolver.Resolve(hint)
	}

	for _, lang := range v.languages {
		langResolver, ok := resolvers[lang]
		if !ok {
			continue
		}

		msg := langResolver.Resolve(hint)
		if msg == "" {
			break
		}

		return msg
	}

	return FallbackMessageResolver.Resolve(hint)
}

// field tries to return a value from the input based on the requested path
// There are 2 main ways of using this function
//
// 1. Absolute path:
//   - "foo.1.bar" = Get from the input (struct) the field "foo", then when it's a list like get the element at index 1 from the list, then get the field "bar" from the struct
//   - "" = Get the source input
//
// 2. Relative path:
//   - ".foo" = Get relative to the currently processed struct the field "foo"
//   - ".1" = Get relative to the currently processed list the element at index 1
//   - "." = Get the currently processed struct
//   - "..foo" = Get the parent of the currentl`y processed struct and then get the field "foo" from it
//
// If nil is returned the field does not exist or path is invalid
// If a needle with only a reflect.Type is returned the path exists but the value is nil
func (v *Validator) field(stack Stack, path string) *Needle {
	relativity := 0
	endRelative := false
	pathParts := []string{}

	for _, part := range strings.Split(path, ".") {
		part = strings.TrimSpace(part)
		if part == "" {
			if endRelative {
				return nil
			} else {
				relativity++
			}
			continue
		} else {
			endRelative = true
		}
		pathParts = append(pathParts, part)
	}

	if relativity == 0 || len(stack) == 0 || relativity > len(stack) {
		// Absolute path
		return resolveWithValue(v.inputValue, pathParts)
	}

	// Relative to the currently processed struct
	stackElement := stack[len(stack)-relativity]
	if stackElement.Parent == nil {
		return resolveWithType(stackElement.ParentType, pathParts)
	}

	return resolveWithValue(*stackElement.Parent, pathParts)
}
