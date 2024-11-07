package laravalidate

import (
	"reflect"
	"strconv"
	"strings"
)

type StackKind uint8

const (
	StackKindObject StackKind = iota
	StackKindList
)

type StackElement struct {
	GoName     string
	JsonName   string
	Index      int // Only for kind == StackKindList
	Kind       StackKind
	Parent     *reflect.Value
	ParentType reflect.Type
}

type Stack []StackElement

func (s Stack) ToPaths() (golang, json string) {
	golang = ""
	json = ""

	for idx, e := range s {
		if idx == 0 {
			golang = e.GoName
			json = e.JsonName
			continue
		}
		golang += "." + e.GoName
		json += "." + e.JsonName
	}

	return
}

func (s Stack) AppendIndex(index int, parent *reflect.Value, parentType reflect.Type) Stack {
	indexStr := strconv.Itoa(index)
	return append(s, StackElement{
		GoName:     indexStr,
		JsonName:   indexStr,
		Index:      index,
		Kind:       StackKindList,
		Parent:     parent,
		ParentType: parentType,
	})
}

func (s Stack) AppendField(field reflect.StructField, parent *reflect.Value, parentType reflect.Type) Stack {
	jsonTag := field.Tag.Get("json")
	jsonTag = strings.Split(jsonTag, ",")[0]

	jsonName := field.Name
	if jsonTag != "" && jsonTag != "-" {
		jsonName = jsonTag
	}

	return append(s, StackElement{
		GoName:     field.Name,
		JsonName:   jsonName,
		Index:      -1,
		Kind:       StackKindObject,
		Parent:     parent,
		ParentType: parentType,
	})
}

// LooslyEquals checks if the stack is equal to the given key
// The key might ignore the index of the list elements and only check the object fields
func (s Stack) LooslyEqualsWithRule(key string, rule string) bool {
	keyParts := strings.Split(key, ".")
	if key == "" {
		keyParts = []string{}
	}

	stackCopy := make(Stack, len(s))
	copy(stackCopy, s)

	if len(keyParts) > 0 && keyParts[len(keyParts)-1] == rule {
		keyParts = keyParts[:len(keyParts)-1]
	}

	for _, part := range keyParts {
		if len(stackCopy) == 0 {
			return false
		}
		stackEl := stackCopy[0]

		parsedPart, err := strconv.Atoi(part)
		if err == nil && parsedPart >= 0 {
			// The part is an array index, the next stack element must be a list entry
			if stackEl.Kind != StackKindList {
				return false
			}
			if stackEl.Index != parsedPart {
				return false
			}

			stackCopy = stackCopy[1:]
			continue
		}

		if part == "*" {
			// The part is an array index wildcard, the next stack element must be a list entry
			if stackEl.Kind != StackKindList {
				return false
			}

			stackCopy = stackCopy[1:]
			continue
		}

		// The part is an object field, the next stack element must be an object
		for idx, elem := range stackCopy {
			if elem.Kind == StackKindList {
				continue
			}

			if elem.Kind != StackKindObject {
				return false
			}

			if elem.GoName != part {
				return false
			}

			stackCopy = stackCopy[1+idx:]
			break
		}
	}

	if len(stackCopy) == 0 {
		return true
	}

	skippedGoRuleName := false
	for _, elem := range stackCopy {
		if elem.Kind == StackKindList {
			continue
		}

		if elem.Kind != StackKindObject {
			return false
		}

		if elem.GoName == rule && !skippedGoRuleName {
			skippedGoRuleName = true
			continue
		}

		return false
	}

	return true
}
