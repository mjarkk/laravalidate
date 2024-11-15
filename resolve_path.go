package laravalidate

import (
	"reflect"
	"strconv"
)

func resolveWithValue(value reflect.Value, path []string) *Needle {
	if len(path) == 0 {
		valueType := value.Type()

		return &Needle{
			Value: &value,
			Type:  valueType,
		}
	}

	for value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return resolveWithType(value.Type().Elem(), path)
		}
		value = value.Elem()
	}

	needle := path[0]
	path = path[1:]

	if needle == "" {
		return nil
	}

	kind := value.Kind()
	switch kind {
	case reflect.Struct:
		field, ok := value.Type().FieldByName(needle)
		if !ok {
			return nil
		}

		return resolveWithValue(value.FieldByIndex(field.Index), path)
	case reflect.Slice, reflect.Array:
		needleNumber, err := strconv.Atoi(needle)
		if err != nil {
			return nil
		}

		if kind == reflect.Slice && value.IsNil() {
			return resolveWithType(value.Type().Elem(), path)
		}

		if needleNumber < 0 || needleNumber >= value.Len() {
			return nil
		}

		return resolveWithValue(value.Index(needleNumber), path)
	case reflect.Map:
		var key reflect.Value
		keySet := false
		valueType := value.Type()
		keyKind := valueType.Key().Kind()

		switch keyKind {
		case reflect.Bool:
			if needle == "true" {
				key = reflect.ValueOf(true)
				keySet = true
			} else if needle == "false" {
				key = reflect.ValueOf(false)
				keySet = true
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			index, err := strconv.ParseInt(needle, 10, 64)
			if err != nil {
				return nil
			}

			switch keyKind {
			case reflect.Int:
				key = reflect.ValueOf(int(index))
			case reflect.Int8:
				if int64(int8(index)) != index {
					return nil
				}
				key = reflect.ValueOf(int8(index))
			case reflect.Int16:
				if int64(int16(index)) != index {
					return nil
				}
				key = reflect.ValueOf(int16(index))
			case reflect.Int32:
				if int64(int32(index)) != index {
					return nil
				}
				key = reflect.ValueOf(int32(index))
			case reflect.Int64:
				key = reflect.ValueOf(int64(index))
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			index, err := strconv.ParseUint(needle, 10, 64)
			if err != nil {
				return nil
			}

			switch keyKind {
			case reflect.Uint:
				key = reflect.ValueOf(int(index))
			case reflect.Int8:
				if uint64(uint8(index)) != index {
					return nil
				}
				key = reflect.ValueOf(int8(index))
			case reflect.Uint16:
				if uint64(uint16(index)) != index {
					return nil
				}
				key = reflect.ValueOf(int16(index))
			case reflect.Uint32:
				if uint64(uint32(index)) != index {
					return nil
				}
				key = reflect.ValueOf(int32(index))
			case reflect.Uint64:
				key = reflect.ValueOf(int64(index))
			}
		case reflect.String:
			key = reflect.ValueOf(needle)
			keySet = true
		}

		if !keySet {
			return nil
		}
		field := value.MapIndex(key)
		if field.Kind() == reflect.Invalid {
			return resolveWithType(valueType.Elem(), path)
		}

		return resolveWithValue(field, path)
	default:
		return nil
	}
}

func resolveWithType(valueType reflect.Type, path []string) *Needle {
	if len(path) == 0 {
		return &Needle{
			Type: valueType,
		}
	}

	for valueType.Kind() == reflect.Ptr {
		valueType = valueType.Elem()
	}

	needle := path[0]
	path = path[1:]

	if needle == "" {
		return nil
	}

	kind := valueType.Kind()
	switch kind {
	case reflect.Struct:
		field, ok := valueType.FieldByName(needle)
		if !ok {
			return nil
		}
		return resolveWithType(field.Type, path)
	}

	return nil
}
