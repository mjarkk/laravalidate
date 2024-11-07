package laravalidate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMsgTemplate(t *testing.T) {
	testCases := []struct {
		msg               string
		expectedVariables []string
	}{
		{
			"",
			nil,
		},
		{
			"foo bar baz bar",
			[]string{},
		},
		{
			":foo",
			[]string{":foo"},
		},
		{
			"text :variable text",
			[]string{":variable"},
		},
		{
			"foo :variable bar :var2 another one :var3",
			[]string{":variable", ":var2", ":var3"},
		},
		{
			"::::::::",
			nil,
		},
		{
			"::foo a:bar -:baz :variable",
			[]string{":variable"},
		},
	}

	for _, tc := range testCases {
		variables := parseMsgTemplate([]byte(tc.msg))
		if len(tc.expectedVariables) == 0 {
			assert.Empty(t, variables, fmt.Sprintf("expected \"%s\" to have no variables", tc.msg))
			continue
		}

		variablesStr := []string{}
		for _, variable := range variables {
			variablesStr = append(variablesStr, string(tc.msg[variable.from:variable.to]))
		}

		assert.Equal(t, tc.expectedVariables, variablesStr, tc.msg)
	}
}
