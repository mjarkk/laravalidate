package laravalidate

import (
	"fmt"
	"strings"

	"golang.org/x/text/language"
)

type ValidatorFn func(ctx *ValidatorCtx) (string, bool)

type registeredValidatorT struct {
	Fn ValidatorFn
	// The map index is the language
	Messages map[string]MessageResolver
}

var validators = map[string]registeredValidatorT{}

type MessageResolver interface {
	Resolve(hint string) string
}

var FallbackMessageResolver = BasicMessageResolver("The :attribute field is invalid")

// BasicMessageResolver is a simple message resolver that always returns the same message, no matter the language or hint
type BasicMessageResolver string

func (d BasicMessageResolver) Resolve(hint string) string {
	return string(d)
}

type MessageHintResolver struct {
	Fallback string
	Hints    map[string]string
}

func (d MessageHintResolver) Resolve(hint string) string {
	msg, ok := d.Hints[hint]
	if ok {
		return msg
	}

	return d.Fallback
}

// TODO: LanguageMessageResolver

// RegisterValidator registers a new validator function
func RegisterValidator(name string, validator ValidatorFn) {
	if validator != nil {
		validators[name] = registeredValidatorT{Fn: validator, Messages: map[string]MessageResolver{}}
	}
}

func BaseRegisterMessages(resolvers map[string]MessageResolver) {
	registerMessagesForLangs([]string{"en", "en-us", "en-gb"}, resolvers)
}

// RegisterMessages registers messages for a validator
func RegisterMessages(lang language.Tag, resolvers map[string]MessageResolver) {
	langStr := strings.ToLower(lang.String())
	langParts := strings.SplitN(langStr, "-", 2)
	langs := []string{langStr}
	if len(langParts) == 2 {
		langs = append(langs, langParts[0])
	}

	registerMessagesForLangs(langs, resolvers)
}

// RegisterMessagesStrict registers messages for a validator
// Compared to RegisterMessages this function does not try to match the language with the base language
// For example if you register a message for "en-GB" it will only be used for "en-GB" and not for "en"
func RegisterMessagesStrict(lang language.Tag, resolvers map[string]MessageResolver) {
	langStr := strings.ToLower(lang.String())
	registerMessagesForLangs([]string{langStr}, resolvers)
}

func registerMessagesForLangs(langs []string, resolvers map[string]MessageResolver) {
	if len(resolvers) == 0 {
		return
	}

	for name, resolver := range resolvers {
		validator, ok := validators[name]
		if !ok {
			fmt.Printf(`Laravalidate: Trying to register error message for validation rule that does not exists "%s"`+"\n", name)
			continue
		}
		for _, lang := range langs {
			validator.Messages[lang] = resolver
		}
	}
}

func LogValidatorsWithoutMessages() {
	for name, validator := range validators {
		if len(validator.Messages) == 0 {
			fmt.Printf(`Laravalidate: No error messages registered for validation rule "%s"`+"\n", name)
		}
	}
}
