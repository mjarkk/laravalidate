package laravalidate

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testValidator struct {
	*testing.T
}

func (v *testValidator) AssertValid(input any) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	assert.Nilf(v.T, GoValidate(ctx, nil, input), "expected no validation error %T%+v", input, input)
}

func (v *testValidator) AssertInvalid(input any) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	assert.NotNilf(v.T, GoValidate(ctx, nil, input), "expected a validation error %T%+v", input, input)
}

func TestValidate(t *testing.T) {
	v := testValidator{t}

	v.AssertValid(struct{}{})

	type SimpleNameStruct struct {
		Name *string `validate:"required"`
	}

	v.AssertInvalid(SimpleNameStruct{})
	name := "John Doe"
	v.AssertValid(SimpleNameStruct{
		Name: &name,
	})

	// Test nested structs
	type A struct {
		A *string `validate:"required"`
	}
	a := A{&name}
	type B struct {
		B  *A
		BB A
	}
	b := B{
		B:  &a,
		BB: a,
	}
	type C struct {
		C  *B
		CC B
	}
	c := C{
		C:  &b,
		CC: b,
	}
	type D struct {
		D  *C `validate:"required"`
		DD C
	}
	d := D{
		D:  &c,
		DD: c,
	}

	v.AssertValid(d)
	d.D = nil
	v.AssertInvalid(d)

	type Empty struct{}

	// Test validateInner
	type ValidateInner struct {
		List []*string `validate:"required" validateInner:"required"`
	}

	v.AssertInvalid(ValidateInner{})
	v.AssertInvalid(ValidateInner{List: []*string{nil}})
	v.AssertValid(ValidateInner{List: []*string{&name}})

	type ValidateOnlyInner struct {
		List []*string `validateInner:"required"`
	}
	v.AssertValid(ValidateOnlyInner{})
	v.AssertInvalid(ValidateOnlyInner{List: []*string{nil}})
	v.AssertValid(ValidateOnlyInner{List: []*string{&name}})

	type ValidateOnlyOuter struct {
		List []*string `validate:"required"`
	}
	v.AssertInvalid(ValidateOnlyOuter{})
	v.AssertValid(ValidateOnlyOuter{List: []*string{nil}})
	v.AssertValid(ValidateOnlyOuter{List: []*string{&name}})

	// Test nested slices
	type ValidateNestedSlice struct {
		list [][][]struct{} `validate:"not_nil"`
	}

	v.AssertInvalid(ValidateNestedSlice{})
	v.AssertValid(ValidateNestedSlice{list: [][][]struct{}{}})
	v.AssertValid(ValidateNestedSlice{list: [][][]struct{}{{nil}}})

	// Nil validation
	type NilValidation struct {
		Field string `validate:"required"`
	}
	var nilValidation *NilValidation
	v.AssertInvalid(nilValidation)

	type NilValidation2 struct {
		Foo *any `validate:"required"`
		Bar any  `validate:"required"`
	}
	v.AssertInvalid(NilValidation2{})
}

func TestErrorMessages(t *testing.T) {
	err := JsonValidate(nil, nil, struct {
		Message string `validate:"required" json:"message"`
	}{})
	assert.NotNil(t, err)
	typedErr, ok := err.(*ValidationError)
	assert.True(t, ok)
	assert.Len(t, typedErr.Errors, 1)

	firstErr := typedErr.Errors[0]
	assert.Equal(t, "message", firstErr.Path)
	assert.Len(t, firstErr.Errors, 1)

	firstValidatorErr := firstErr.Errors[0]
	assert.Equal(t, FieldValidatorError{
		Rule:    "required",
		Hint:    "required",
		Message: "The message field is required.",
	}, firstValidatorErr)

	laravelError := typedErr.ToLaravelError()
	laravelErrorBytes, err := json.Marshal(laravelError)
	assert.NoError(t, err)

	assert.Equal(t, `{"errors":{"message":["The message field is required."]},"message":"Form contains errors"}`, string(laravelErrorBytes))
}

type TestCustomErrorMessageT struct {
	Inner []TestCustomErrorMessageInnerT
}

type TestCustomErrorMessageInnerT struct {
	Inner2 struct {
		Inner3 string `validate:"required"`
	}
}

func (TestCustomErrorMessageT) ValidationMessages() []CustomError {
	return []CustomError{
		{"Inner.Inner2.Inner3.required", BasicMessageResolver("Yay custom error message!")},
	}
}

func TestCustomErrorMessages(t *testing.T) {
	err := JsonValidate(nil, nil, &TestCustomErrorMessageT{Inner: []TestCustomErrorMessageInnerT{{}}})

	assert.NotNil(t, err)
	typedErr, ok := err.(*ValidationError)
	assert.True(t, ok)
	assert.Len(t, typedErr.Errors, 1)

	firstErr := typedErr.Errors[0]
	assert.Equal(t, "Inner.0.Inner2.Inner3", firstErr.Path)
	assert.Len(t, firstErr.Errors, 1)

	firstValidatorErr := firstErr.Errors[0]
	assert.Equal(t, FieldValidatorError{
		Rule:    "required",
		Hint:    "required",
		Message: "Yay custom error message!",
	}, firstValidatorErr)
}
