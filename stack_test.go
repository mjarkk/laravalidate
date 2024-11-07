package laravalidate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackLooslyEquals(t *testing.T) {
	nestedStack := Stack{
		{Kind: StackKindObject, GoName: "Foo"},
		{Kind: StackKindList, Index: 1},
		{Kind: StackKindObject, GoName: "Bar"},
		{Kind: StackKindList, Index: 2},
	}

	testCases := []struct {
		Name     string
		Stack    Stack
		Key      string
		Expected bool
	}{
		{
			"empty",
			Stack{},
			"",
			true,
		},
		{
			"simple object",
			Stack{{Kind: StackKindObject, GoName: "Name"}},
			"Name",
			true,
		},
		{
			"simple object (no match)",
			Stack{{Kind: StackKindObject, GoName: "SussyBakka"}},
			"Name",
			false,
		},
		{
			"nested (not matches key length)",
			nestedStack,
			"Foo",
			false,
		},
		{
			"nested",
			nestedStack,
			"Foo.Bar",
			true,
		},
		{
			"nested with index",
			nestedStack,
			"Foo.1.Bar.2",
			true,
		},
		{
			"nested with wildcard index",
			nestedStack,
			"Foo.*.Bar.*",
			true,
		},
		{
			"nested with rule name",
			nestedStack,
			"Foo.Bar.required",
			true,
		},
	}

	for _, testCase := range testCases {
		equals := testCase.Stack.LooslyEqualsWithRule(testCase.Key, "required")
		if testCase.Expected {
			assert.True(t, equals, testCase.Name)
		} else {
			assert.False(t, equals, testCase.Name)
		}
	}
}
