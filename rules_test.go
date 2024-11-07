package laravalidate

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func validatorCtx(value *reflect.Value, valueType reflect.Type, args []string) *ValidatorCtx {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return &ValidatorCtx{
		ctx: ctx,
		Needle: Needle{
			Value: value,
			Type:  valueType,
		},
		Args: args,
		state: &ValidatorCtxState{
			bail:      false,
			state:     map[string]any{},
			stack:     Stack{},
			validator: newValidator(ctx, nil, reflect.ValueOf(struct{}{}), GoMode),
		},
	}
}

// TestRulesFuzz tests if no pannics occur when running all rules
func TestRulesFuzz(t *testing.T) {
	// var empty any

	args := [][]string{
		{},
		{"foo"},
		{"1"},
		{"2", "4"},
		{"2", "foo"},
	}

	testValue := "foo"
	testValuePtr := &testValue
	testValuePtrPtr := &testValuePtr
	testValuePtrPtrPtr := &testValuePtrPtr
	values := []any{
		testValue,
		testValuePtr,
		testValuePtrPtr,
		testValuePtrPtrPtr,
		3,
		uint8(3),
		1,
		uint8(1),
		int16(1),
		-1,
		0,
		3.5,
		true,
		false,
		"",
		"yes",
		"no",
		[]string{},
		[]string{"foo"},
		[]string{"foo", "bar"},
		`{"hello": "world"}`,
		[]byte("hello world"),
		[]rune("hello world"),
		'a',
		byte('a'),
		func() {},
		map[string]string{},
		map[string]string{
			"foo": "bar",
		},
	}

	for name, validator := range validators {
		t.Run(name, func(t *testing.T) {
			for _, arg := range args {
				for _, value := range values {
					valueReflection := reflect.ValueOf(value)
					valueType := valueReflection.Type()
					assert.NotPanicsf(t, func() {
						validator.Fn(validatorCtx(nil, valueType, arg))
					}, "value=%+v, args=%+v onlyType=true", value, arg)
					assert.NotPanicsf(t, func() {
						validator.Fn(validatorCtx(&valueReflection, valueType, arg))
					}, "value=%+v, args=%+v onlyType=false", value, arg)
				}
			}
		})
	}
}

func validationRulePasses(t *testing.T, validator ValidatorFn, value any, args []string, baseStruct ...any) {
	valueReflection := reflect.ValueOf(value)
	valueType := valueReflection.Type()

	failReason, passes := validator(validatorCtx(&valueReflection, valueType, args))
	assert.Truef(t, passes, "failHint=%s, value=%+v, args=%+v", failReason, value, args)
}

func validationRuleInvalid(t *testing.T, validator ValidatorFn, value any, args []string) {
	valueReflection := reflect.ValueOf(value)
	valueType := valueReflection.Type()

	_, passes := validator(validatorCtx(&valueReflection, valueType, args))
	assert.Falsef(t, passes, "value=%+v, args=%+v", value, args)
}

func TestURL(t *testing.T) {
	validationRulePasses(t, URL, "http://example.com", []string{"http", "https"})
	validationRulePasses(t, URL, "http://example.com", nil)
	validationRulePasses(t, URL, "steam://some_game", nil)
	validationRuleInvalid(t, URL, "this is not a url", nil)
	validationRuleInvalid(t, URL, "steam://some_game", []string{"http", "https"})
}

func TestUuid(t *testing.T) {
	validationRulePasses(t, Uuid, "550e8400-e29b-41d4-a716-446655440000", nil)
	validationRulePasses(t, Uuid, "550e8400-e29b-41d4-a716-446655440000", []string{"3", "4", "5"})
	validationRuleInvalid(t, Uuid, "this is not a uuid", nil)
	validationRuleInvalid(t, Uuid, "550e8400-e29b-41d4-a716-446655440000", []string{"1", "2", "3"})
}

func TestRegex(t *testing.T) {
	validationRulePasses(t, Regex, "foo", []string{"/^foo$/"})
	validationRuleInvalid(t, NotRegex, "foo", []string{"/^foo$/"})
	validationRuleInvalid(t, Regex, "foo", []string{"/^bar$/"})
	validationRulePasses(t, NotRegex, "foo", []string{"/^bar$/"})
}

func TestIp(t *testing.T) {
	validationRulePasses(t, IP, "1.1.1.1", nil)
	validationRulePasses(t, IPV4, "1.1.1.1", nil)
	validationRuleInvalid(t, IPV6, "1.1.1.1", nil)

	validationRulePasses(t, IP, "2606:4700:4700::1111", nil)
	validationRuleInvalid(t, IPV4, "2606:4700:4700::1111", nil)
	validationRulePasses(t, IPV6, "2606:4700:4700::1111", nil)

	validationRuleInvalid(t, IP, "This is not an ip", nil)
}

func TestMime(t *testing.T) {
	validationRulePasses(t, Mimetypes, "image/png", []string{"image/png"})
	validationRuleInvalid(t, Mimetypes, "image/jpg", []string{"image/png"})

	validationRulePasses(t, Mimetypes, "image/png", []string{"image/*"})
	validationRuleInvalid(t, Mimetypes, "application/html", []string{"image/*"})

	validationRulePasses(t, Mimes, "image/png", []string{"png"})
	validationRulePasses(t, Mimes, "image/jpeg", []string{"jpeg"})
}

func TestMac(t *testing.T) {
	validationRulePasses(t, MacAddress, "00:00:5e:00:53:01", nil)
	validationRuleInvalid(t, MacAddress, "00:00:5e", nil)
}

func TestConfirmed(t *testing.T) {
	v := &testValidator{t}

	type Test struct {
		Field             string `validate:"required|confirmed"`
		FieldConfirmation string
	}
	v.AssertInvalid(Test{
		Field:             "Foo",
		FieldConfirmation: "Bar",
	})

	v.AssertValid(Test{
		Field:             "Foo",
		FieldConfirmation: "Foo",
	})
}
