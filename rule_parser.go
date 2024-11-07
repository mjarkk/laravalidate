package laravalidate

import (
	"fmt"
	"strings"
)

type validationRule struct {
	validator registeredValidatorT
	name      string
	args      []string
}

func validationRules(input string) []validationRule {
	rules := []validationRule{}
	if len(input) == 0 {
		return rules
	}

	sections := strings.Split(input, "|")
	for _, section := range sections {
		if len(section) == 0 {
			continue
		}

		nameAndArgs := strings.SplitN(section, ":", 2)
		name := nameAndArgs[0]
		if name == "" {
			continue
		}

		validator, ok := validators[name]
		if !ok {
			fmt.Printf(`Laravalidate: Unknown validation rule "%s"`+"\n", name)
			continue
		}

		args := []string{}
		if len(nameAndArgs) > 1 {
			args = strings.Split(nameAndArgs[1], ",")
		}

		rules = append(rules, validationRule{
			validator: validator,
			name:      name,
			args:      args,
		})
	}

	return rules
}
