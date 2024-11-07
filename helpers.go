package laravalidate

import "reflect"

// equal is a helper function to compare two reflect.Value
// This method is similar to reflect.DeepEqual but it's less strict
// for example pointers don't have to point to the same memory address
func equal(a reflect.Value, b reflect.Value) bool {
	if a.Kind() != b.Kind() {
		return false
	}

	switch a.Kind() {
	case reflect.Invalid:
		return true
	case reflect.Bool:
		return a.Bool() == b.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return a.Int() == b.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return a.Uint() == b.Uint()
	case reflect.Float64, reflect.Float32:
		return a.Float() == b.Float()
	case reflect.Uintptr, reflect.Complex64, reflect.Complex128, reflect.UnsafePointer:
		// FIXME, not comparing these values
		return false
	case reflect.Slice:
		if a.IsNil() && b.IsNil() {
			return true
		}
		if a.IsNil() || b.IsNil() {
			return false
		}
		fallthrough
	case reflect.Array:
		if a.Len() != b.Len() {
			return false
		}

		for idx := 0; idx < a.Len(); idx++ {
			if !equal(a.Index(idx), b.Index(idx)) {
				return false
			}
		}

		return true
	case reflect.Chan, reflect.Func:
		// Uncomparable
		return false
	case reflect.Interface:
		// We cannot compare interfaces
		return false
	case reflect.Map:
		if a.IsNil() && b.IsNil() {
			return true
		}
		if a.IsNil() || b.IsNil() {
			return false
		}

		if a.Len() != b.Len() {
			return false
		}
		aType := a.Type()
		bType := b.Type()
		if aType.Key().Kind() != bType.Key().Kind() {
			return false
		}
		if aType.Elem().Kind() != bType.Elem().Kind() {
			return false
		}

		iter := a.MapRange()
		for iter.Next() {
			key := iter.Key()
			bValue := b.MapIndex(key)
			if bValue.Kind() == reflect.Invalid {
				return false
			}

			aValue := iter.Value()
			if !equal(aValue, bValue) {
				return false
			}
		}
		return true
	case reflect.Pointer:
		if a.IsNil() && b.IsNil() {
			return true
		}
		if a.IsNil() || b.IsNil() {
			return false
		}
		return equal(a.Elem(), b.Elem())
	case reflect.String:
		return a.String() == b.String()
	case reflect.Struct:
		aType := a.Type()
		bType := b.Type()

		if aType.Name() != bType.Name() || aType.PkgPath() != bType.PkgPath() {
			return false
		}

		if a.NumField() != b.NumField() {
			// This should not be needed but just to be sure
			return false
		}

		for idx := 0; idx < a.NumField(); idx++ {
			aValue := a.Field(idx)
			bValue := b.Field(idx)
			if !equal(aValue, bValue) {
				return false
			}
		}

		return true
	}

	return false
}
