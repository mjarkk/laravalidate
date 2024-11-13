# `Laravalidate`

A go package that has Laravel like validation.

It takes inspiration from Laravel's validation system but is not a 1 on 1 copy.

It's designed to make it easy to convert the validation rules from a Laravel project to a Go project.

```sh
go get -u github.com/mjarkk/laravalidate
```

```go
package main

import (
	"context"
	"fmt"

	"github.com/mjarkk/laravalidate"
)

func main() {
	input := struct {
		Name string `json:"name" validate:"required"`
	}{
		Name: "",
	}

	err := laravalidate.JsonValidate(context.Background(), nil, input)
	fmt.Println(err.Error()) // The name field is required.
}
```

## Rules

All the rules are defined in: [RULES.md](./RULES.md)

Most laraval rules are supported except for the database related rules.

Rules can be set using the `validate` tag on a struct field like:

```go
type UserRequest struct {
	Email string `json:"email" validate:"email"`
	Name  string `json:"name" validate:"required"`
}
```

Multiple rules can be set by separating them with a `|` like:

```go
type UserRequest struct {
	Email string `json:"email" validate:"required|email"`
}
```

When validating an array the array can be validated using the `validate` tag and it's elements can be validated using the `validateInner` tag like:

```go
type UserRequest struct {
	// The emails list is required and must at least contain one element.
	// The elements in the list must be valid email addresses.
	Emails []string `json:"emails" validate:"required" validateInner:"email"`
}
```

## More error info

```go
type UserRequest struct {
	Email string `json:"email" validate:"email"`
	Name  string `json:"name" validate:"required"`
}

func main() {
	err := laravalidate.JsonValidate(context.Background(), nil, UserRequest{})
	if err == nil {
		os.Exit(0)
	}

	errInfo := err.(*laravalidate.ValidationError)
	for _, fieldError := range errInfo.Errors {
		fmt.Println("field:", fieldError.JsonPath)
		for _, err := range fieldError.Errors {
			fmt.Printf("  %+v\n", err)
		}
	}
}
```

## Laravel style errors

For when you are converting a laravel application to a go application and want to keep the error messages the same.

```go
type UserRequest struct {
	Email string `json:"email" validate:"email"`
	Name  string `json:"name" validate:"required"`
}

func main() {
	err := laravalidate.JsonValidate(context.Background(), nil, UserRequest{})
	if err == nil {
		os.Exit(0)
	}

	laravelErr := err.(*laravalidate.ValidationError).ToLaravelError()
	laravelErrJson, err := json.MarshalIndent(laravelErr, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(laravelErrJson))
	/*{
	  "errors": {
	    "email": [
	      "The email field must be a valid email address."
	    ],
	    "name": [
	      "The name field is required."
	    ]
	  },
	  "message": "Form contains errors"
  }*/
}
```

## Go style errors

It might be that you are not validating json messages but rather go structs or something else.
In these cases you might want the go names rather than the json names in the error messages.

This can be done using the `laravalidate.GoValidate` method

```go
func main() {
	translations.RegisterNlTranslations()

	input := struct {
		Name string `json:"name" validate:"required"`
	}{
		Name: "",
	}

	err := laravalidate.GoValidate(context.Background(), nil, input)
	fmt.Println(err.Error()) // The Name field is required.
}
```

## Custom error messages

Sometimes you want to provide a custom error message for a specific field.

```go
type UserRequest struct {
	Email string `json:"email" validate:"email"`
	Name  string `json:"name" validate:"required"`
}

func (UserRequest) ValidationMessages() []laravalidate.CustomError {
	return []laravalidate.CustomError{
		{Key: "Email", Resolver: laravalidate.BasicMessageResolver("Bro email is required!")},
	}
}

func main() {
	err := laravalidate.JsonValidate(context.Background(), nil, UserRequest{})
	fmt.Println(err.Error()) // Bro email is required!
}
```

Nested fields can be accessed using the `.` separator.
For example when a field is nested under `Foo[3].Bar.Baz[2]` you can define the keys for it the following ways:

- `Foo.3.Bar.Baz.2` (very strict)
- `Foo.*.Bar.Baz.*` (lists have wild cards)
- `Foo.Bar.Baz` (same as the prevouse one)

specific validators can also be targeted like so:

```go
type UserRequest struct {
	Email string `json:"email" validate:"required|email"`
}

func (UserRequest) ValidationMessages() []laravalidate.CustomError {
	return []laravalidate.CustomError{
		{Key: "Email.required", Resolver: laravalidate.BasicMessageResolver("Bro email is required!")},
		{Key: "Email.email", Resolver: laravalidate.BasicMessageResolver("Bogus email!")},
	}
}
```

## Translations

You can provide a list of languages to the `JsonValidate` function to get translated errors.
This only works if you have registered custom error messages for a specific language.

```go
package main

import (
	"context"
	"fmt"

	"github.com/mjarkk/laravalidate"
	"github.com/mjarkk/laravalidate/translations"
	"golang.org/x/text/language"
)

func main() {
	// Note that we have to register the translations before we can use them.
	translations.RegisterNlTranslations()

	input := struct {
		Name string `json:"name" validate:"required"`
	}{
		Name: "",
	}

	err := laravalidate.JsonValidate(
		context.Background(),
		[]language.Tag{
			language.Dutch,
		},
		input,
	)
	fmt.Println(err.Error()) // Het name veld is verplicht.
}
```

Out of the box supported languages are:

- English (default)
- German `translations.RegisterDeTranslations()`
- Dutch `translations.RegisterNlTranslations()`
- French `translations.RegisterFrTranslations()`
- Spanish `translations.RegisterEsTranslations()`

## Custom translations

See how other translations are done inside of the [./translations](./translations) folder

## Writing a custom validator

```go
// Accepted is a custom validator that checks if the value is accepted.
//
// It returns 2 values:
// - string : Hint for the error message, this can be used to change error messages based on the hint.
// - bool : Passes, if false the validation failed and you'll get an error message
//
// Note that this validator is already part of the laravel validation rules and just here for example.
func Accepted(ctx *ValidatorCtx) (string, bool) {
	// Unwrap the pointer to the actual value if it is a pointer.
	// Note that this might result in a type but no value inside the ctx.
	ctx.UnwrapPointer()

	// Check if the value type is one of the following:
	if !ctx.IsKind(
		reflect.Bool,
		reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
	) {
		return "invalid_type", false
	}

	// Check if the ctx has a value
	if !ctx.HasValue() {
		return "unacceptable", false
	}

	switch ctx.Kind() {
	case reflect.Bool:
		if !ctx.Value.Bool() {
			return "unacceptable", false
		}
	case reflect.String:
		switch ctx.Value.String() {
		case "yes", "on", "1", "true":
			return "", true
		}
		return "unacceptable", false
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if ctx.Value.Int() == 1 {
			return "", true
		}
		return "unacceptable", false
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if ctx.Value.Uint() == 1 {
			return "", true
		}
		return "unacceptable", false
	default:
		return "invalid_type", false
	}

	return "", true
}

func main() {
	laravalidate.RegisterValidator("accepted", Accepted)
	laravalidate.BaseRegisterMessages(map[string]laravalidate.MessageResolver{
		"accepted": laravalidate.BasicMessageResolver("The :attribute field must be accepted."),
	})
	// Now you can use the accepted validator
}
```

There are a lot more methods on the `ValidatorCtx` that you can use to get the value of the field.

See the [rules.go](./rules.go) for examples.
