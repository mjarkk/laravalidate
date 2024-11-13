package laravalidate

import (
	"reflect"
	"time"

	"github.com/mjarkk/laravalidate/dates"
)

type Needle struct {
	Type reflect.Type
	// Value might be nil if we are validating a pointer to a nil struct.
	// Also running ctx.UnwrapPointer() might cause the value to be unset
	Value *reflect.Value
}

// HasValue returns if the needle has a value
func (n *Needle) HasValue() bool {
	return n.Value != nil
}

// Kind returns the kind of the value if available an falls back to the kind of the type
func (n *Needle) Kind() reflect.Kind {
	if n.Value != nil {
		return n.Value.Kind()
	}

	return n.Type.Kind()
}

// UnwrapPointer unwraps the pointer
func (n *Needle) UnwrapPointer() bool {
	if n.Value != nil {
		for n.Value.Kind() == reflect.Ptr {
			if n.Value.IsNil() {
				n.Type = n.Type.Elem()
				n.Value = nil
				break
			}

			*n.Value = n.Value.Elem()
			n.Type = n.Value.Type()
		}
	}

	// Do not change this to an else block as in the code above we might set this value to nil
	if n.Value == nil {
		for n.Type.Kind() == reflect.Ptr {
			n.Type = n.Type.Elem()
		}

		return false
	}

	return true
}

type ConvertStatus uint8

const (
	ConverstionOk ConvertStatus = iota
	InvalidType                 // Usually means you should return ("invalid_type", false)
	Invalid                     // Usually means you should return ("invalid", false)
	ValueNil                    // Usually means you should return ("", true)
)

func (s ConvertStatus) Response() (string, bool) {
	if s == InvalidType {
		return "invalid_type", false
	} else if s == Invalid {
		return "invalid", false
	}
	return "", true
}

func (s ConvertStatus) Oke() bool {
	return s == ConverstionOk
}

// String is a wrapper around (*ValidatorCtx).value.String()
// This one checks the following extra things and if one of them does not match returns ("", false)
// 1. The value is set
// 2. The value is a string
//
// Example:
// ```
//
//	str, status := ctx.String()
//
//	if !status.Oke() {
//	  return status.Response()
//	}
//
//	fmt.Println(str)
//
// ```
func (n *Needle) String() (string, ConvertStatus) {
	if n.Kind() != reflect.String {
		return "", InvalidType
	}

	if !n.HasValue() {
		return "", ValueNil
	}

	return n.Value.String(), ConverstionOk
}

// StringLike returns all values that can be interpreted as a string
//
// Supported:
// - string
// - []byte
// - []rune
// - rune
// - byte
//
// Example:
// ```
//
//	str, status := ctx.StringLike()
//
//	if !status.Oke() {
//	  return status.Response()
//	}
//
//	fmt.Println(str)
//
// ```
func (n *Needle) StringLike() (string, ConvertStatus) {
	if !n.IsKind(reflect.String, reflect.Int32, reflect.Uint8) {
		if n.Kind() != reflect.Slice {
			return "", InvalidType
		}

		sliceElementKind := n.Type.Elem().Kind()
		if sliceElementKind != reflect.Uint8 && sliceElementKind != reflect.Int32 {
			return "", InvalidType
		}
	}

	if !n.HasValue() {
		return "", ValueNil
	}

	kind := n.Kind()
	if kind == reflect.String {
		return n.Value.String(), ConverstionOk
	} else if kind == reflect.Int32 {
		return string(rune(n.Value.Int())), ConverstionOk
	} else if kind == reflect.Uint8 {
		return string(uint8(n.Value.Uint())), ConverstionOk
	} else if kind != reflect.Slice {
		return "", InvalidType
	}

	// Value is a slice
	contentKind := n.Type.Elem().Kind()
	if contentKind == reflect.Uint8 {
		return string(n.Value.Bytes()), ConverstionOk
	} else if contentKind == reflect.Int32 {
		return string(n.Value.Interface().([]rune)), ConverstionOk
	}

	return "", InvalidType
}

// IsKinds asserts the .Kind method for one of the input values
// If it matches one it returns true
func (n *Needle) IsKind(kinds ...reflect.Kind) bool {
	currentKind := n.Kind()
	for _, kind := range kinds {
		if kind == currentKind {
			return true
		}
	}

	return false
}

// IsInt returns true if the type kind is a int[..]
func (n *Needle) IsInt() bool {
	switch n.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}
}

// IsUint returns true if the type kind is a uint[..]
func (n *Needle) IsUint() bool {
	switch n.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	default:
		return false
	}
}

// IsFloat returns true if the type kind is a float
func (n *Needle) IsFloat() bool {
	switch n.Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

// IsNumeric returns true if the type kind is a int[..] , uint[..] or float[..]
func (n *Needle) IsNumeric() bool {
	switch n.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

// IsBool returns true if the type kind is a list
func (n *Needle) IsList() bool {
	switch n.Kind() {
	case reflect.Slice, reflect.Array:
		return true
	default:
		return false
	}
}

// HasLen returns true if the type kind supports the reflect.Len method
func (n *Needle) HasLen() bool {
	switch n.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return true
	default:
		return false
	}
}

// Date tries to convert the value to a time.Time
func (n *Needle) Date() (time.Time, ConvertStatus) {
	if n.IsKind(reflect.String, reflect.Int64) {
		// Continue
	} else if n.Kind() == reflect.Struct && n.Type.ConvertibleTo(reflect.TypeOf(time.Time{})) {
		// Continue
	} else {
		return time.Time{}, InvalidType
	}

	if !n.HasValue() {
		return time.Time{}, ValueNil
	}

	if n.Kind() == reflect.Struct {
		return n.Value.Interface().(time.Time), ConverstionOk
	}

	switch n.Kind() {
	case reflect.String:
		t, ok := dates.ParseStructuredDate(n.Value.String())
		if ok {
			return t, ConverstionOk
		}
		return time.Time{}, Invalid
	case reflect.Int64:
		t, ok := dates.ParseUnix(n.Value.Int())
		if !ok {
			return time.Time{}, Invalid
		}
		return t, ConverstionOk
	default:
		return time.Time{}, InvalidType
	}
}

// Float64 tries to convert the number to a float64
func (n *Needle) Float64() (float64, bool) {
	if !n.HasValue() {
		return 0, false
	}

	switch n.Kind() {
	case reflect.Float32, reflect.Float64:
		return n.Value.Float(), true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(n.Value.Int()), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(n.Value.Uint()), true
	}

	return 0, false
}

// Int64 tries to convert the number to a int64
func (n *Needle) Int64() (int64, bool) {
	if !n.HasValue() {
		return 0, false
	}

	switch n.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return n.Value.Int(), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(n.Value.Uint()), true
	}

	return 0, false
}
