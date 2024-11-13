package laravalidate

import (
	"encoding/json"
	"net"
	"net/mail"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

func init() {
	RegisterValidator("accepted", Accepted)

	// Accepted If

	RegisterValidator("active_url", ActiveUrl)
	RegisterValidator("after", AfterDate)
	RegisterValidator("after_or_equal", AfterOrEqualDate)
	RegisterValidator("alpha", Alpha)
	RegisterValidator("alpha_dash", AlphaDash)
	RegisterValidator("alpha_numeric", AlphaNumeric)
	// Unsupported: Array
	RegisterValidator("ascii", Ascii)
	RegisterValidator("bail", Bail)
	RegisterValidator("before", BeforeDate)
	RegisterValidator("before_or_equal", BeforeOrEqualDate)
	RegisterValidator("between", Between)
	RegisterValidator("boolean", Boolean)
	RegisterValidator("confirmed", Confirmed)

	// Contains
	// Current Password

	RegisterValidator("date", Date)

	// Date Equals

	RegisterValidator("date_format", DateFormat)
	// Unsupported: Decimal
	RegisterValidator("declined", Declined)

	// Declined If
	// Different

	RegisterValidator("digits", Digits)
	RegisterValidator("digits_between", DigitsBetween)

	// Dimensions (Image Files)
	// Distinct
	// Doesnt Start With
	// Doesnt End With

	RegisterValidator("email", Email)
	RegisterValidator("ends_with", EndsWith)

	// Enum
	// Exclude
	// Exclude If
	// Exclude Unless
	// Exclude With
	// Exclude Without

	// Provided by dbrules: Exists
	RegisterValidator("extensions", Extensions)

	// File

	RegisterValidator("filled", Filled)
	RegisterValidator("gt", Gt)
	RegisterValidator("gte", Gte)
	RegisterValidator("hex_color", HexColor)

	// Image (File)

	RegisterValidator("in", In)

	// In Array
	// Integer

	RegisterValidator("ip", IP)
	RegisterValidator("ipv4", IPV4)
	RegisterValidator("ipv6", IPV6)
	RegisterValidator("json", JSON)
	RegisterValidator("lt", Lt)
	RegisterValidator("lte", Lte)
	RegisterValidator("lowercase", Lowercase)
	// Unsupported: List
	RegisterValidator("mac_address", MacAddress)
	RegisterValidator("max", Max)
	RegisterValidator("max_digits", MaxDigits)
	RegisterValidator("mimetypes", Mimetypes)
	RegisterValidator("mimes", Mimes)
	RegisterValidator("min", Min)
	RegisterValidator("min_digits", MinDigits)

	// Missing
	// Missing If
	// Missing Unless
	// Missing With
	// Missing With All
	// Multiple Of

	RegisterValidator("not_nil", NotNil)
	RegisterValidator("not_in", NotIn)
	RegisterValidator("not_regex", NotRegex)
	// Unsupported: Nullable
	RegisterValidator("numeric", Numeric)

	// Present
	// Present If
	// Present Unless
	// Present With
	// Present With All
	// Prohibited
	// Prohibited If
	// Prohibited Unless
	// Prohibits

	RegisterValidator("regex", Regex)
	RegisterValidator("required", Required)

	// Required If
	// Required If Accepted
	// Required If Declined
	// Required Unless
	// Required With
	// Required With All
	// Required Without
	// Required Without All
	// Required Array Keys
	// Same

	RegisterValidator("size", Size)

	// Sometimes

	RegisterValidator("starts_with", StartsWith)
	// Unsupported: String

	// Timezone
	// Unique (Database)

	RegisterValidator("uppercase", Uppercase)
	RegisterValidator("url", URL)
	RegisterValidator("ulid", Ulid)
	RegisterValidator("uuid", Uuid)

	BaseRegisterMessages(map[string]MessageResolver{
		"accepted": BasicMessageResolver("The :attribute field must be accepted."),
		// "accepted_if": BasicMessageResolver("The :attribute field must be accepted when :other is :value."),
		"active_url":     BasicMessageResolver("The :attribute field must be a valid URL."),
		"after":          BasicMessageResolver("The :attribute field must be a date after :date."),
		"after_or_equal": BasicMessageResolver("The :attribute field must be a date after or equal to :date."),
		"alpha":          BasicMessageResolver("The :attribute field must only contain letters."),
		"alpha_dash":     BasicMessageResolver("The :attribute field must only contain letters, numbers, dashes, and underscores."),
		"alpha_numeric":  BasicMessageResolver("The :attribute field must only contain letters and numbers."),
		// "array":     BasicMessageResolver("The :attribute field must be an array."),
		"ascii":           BasicMessageResolver("The :attribute field must only contain single-byte alphanumeric characters and symbols."),
		"bail":            BasicMessageResolver("The :attribute field must pass."),
		"before":          BasicMessageResolver("The :attribute field must be a date before :date."),
		"before_or_equal": BasicMessageResolver("The :attribute field must be a date before or equal to :date."),
		"between": MessageHintResolver{
			Fallback: "The :attribute field must be between :arg0 and :arg1.",
			Hints: map[string]string{
				"array":   "The :attribute field must have between :arg0 and :arg1 items.",
				"file":    "The :attribute field must be between :arg0 and :arg1 kilobytes.",
				"numeric": "The :attribute field must be between :arg0 and :arg1.",
				"string":  "The :attribute field must be between :arg0 and :arg1 characters.",
			},
		},
		"boolean": BasicMessageResolver("The :attribute field must be true or false."),
		// "can": BasicMessageResolver("The :attribute field contains an unauthorized value."),
		"confirmed": BasicMessageResolver("The :attribute field confirmation does not match."),
		// "contains":  BasicMessageResolver("The :attribute field is missing a required value."),
		// "current_password": BasicMessageResolver("The password is incorrect."),
		"date": BasicMessageResolver("The :attribute field must be a valid date."),
		// "date_equals": BasicMessageResolver("The :attribute field must be a date equal to :date."),
		"date_format": BasicMessageResolver("The :attribute field must match the format :arg."),
		// "decimal": BasicMessageResolver("The :attribute field must have :arg decimal places."),
		"declined": BasicMessageResolver("The :attribute field must be declined."),
		// "declined_if": BasicMessageResolver("The :attribute field must be declined when :other is :value."),
		// "different":   BasicMessageResolver("The :attribute field and :other must be different."),
		"digits":         BasicMessageResolver("The :attribute field must be :digits digits."),
		"digits_between": BasicMessageResolver("The :attribute field must be between :arg0 and :arg1 digits."),
		// "dimensions":        BasicMessageResolver("The :attribute field has invalid image dimensions."),
		// "distinct":          BasicMessageResolver("The :attribute field has a duplicate value."),
		// "doesnt_end_with":   BasicMessageResolver("The :attribute field must not end with one of the following: :args."),
		// "doesnt_start_with": BasicMessageResolver("The :attribute field must not start with one of the following: :args."),
		"email":     BasicMessageResolver("The :attribute field must be a valid email address."),
		"ends_with": BasicMessageResolver("The :attribute field must end with one of the following: :args."),
		// "enum":    BasicMessageResolver("The selected :attribute is invalid."),
		"exists":     BasicMessageResolver("The selected :attribute is invalid."),
		"extensions": BasicMessageResolver("The :attribute field must have one of the following extensions: :args."),
		// "file":       BasicMessageResolver("The :attribute field must be a file."),
		"filled": BasicMessageResolver("The :attribute field must have a value."),
		"gt": MessageHintResolver{Hints: map[string]string{
			"array":   "The :attribute field must have more than :value items.",
			"file":    "The :attribute field must be greater than :value kilobytes.",
			"numeric": "The :attribute field must be greater than :value.",
			"string":  "The :attribute field must be greater than :value characters.",
		}},
		"gte": MessageHintResolver{Hints: map[string]string{
			"array":   "The :attribute field must have :value items or more.",
			"file":    "The :attribute field must be greater than or equal to :value kilobytes.",
			"numeric": "The :attribute field must be greater than or equal to :value.",
			"string":  "The :attribute field must be greater than or equal to :value characters.",
		}},
		"hex_color": BasicMessageResolver("The :attribute field must be a valid hexadecimal color."),
		// "image":     BasicMessageResolver("The :attribute field must be an image."),
		"in": BasicMessageResolver("The selected :attribute is invalid."),
		// "in_array":  BasicMessageResolver("The :attribute field must exist in :other."),
		// "integer":   BasicMessageResolver("The :attribute field must be an integer."),
		"ip":   BasicMessageResolver("The :attribute field must be a valid IP address."),
		"ipv4": BasicMessageResolver("The :attribute field must be a valid IPv4 address."),
		"ipv6": BasicMessageResolver("The :attribute field must be a valid IPv6 address."),
		"json": BasicMessageResolver("The :attribute field must be a valid JSON string."),
		// "list":      BasicMessageResolver("The :attribute field must be a list."),
		"lowercase": BasicMessageResolver("The :attribute field must be lowercase."),
		"lt": MessageHintResolver{Hints: map[string]string{
			"array":   "The :attribute field must have less than :value items.",
			"file":    "The :attribute field must be less than :value kilobytes.",
			"numeric": "The :attribute field must be less than :value.",
			"string":  "The :attribute field must be less than :value characters.",
		}},
		"lte": MessageHintResolver{Hints: map[string]string{
			"array":   "The :attribute field must not have more than :value items.",
			"file":    "The :attribute field must be less than or equal to :value kilobytes.",
			"numeric": "The :attribute field must be less than or equal to :value.",
			"string":  "The :attribute field must be less than or equal to :value characters.",
		}},
		"mac_address": BasicMessageResolver("The :attribute field must be a valid MAC address."),
		"max": MessageHintResolver{
			Fallback: "The :attribute field must not be greater than :arg.",
			Hints: map[string]string{
				"array":   "The :attribute field must not have more than :arg items.",
				"file":    "The :attribute field must not be greater than :arg kilobytes.",
				"numeric": "The :attribute field must not be greater than :arg.",
				"string":  "The :attribute field must not be greater than :arg characters.",
			},
		},
		"max_digits": BasicMessageResolver("The :attribute field must not have more than :max digits."),
		"mimes":      BasicMessageResolver("The :attribute field must be a file of type: :args."),
		"mimetypes":  BasicMessageResolver("The :attribute field must be a file of type: :args."),
		"min": MessageHintResolver{
			Fallback: "The :attribute field must be at least :arg.",
			Hints: map[string]string{
				"array":   "The :attribute field must have at least :arg items.",
				"file":    "The :attribute field must be at least :arg kilobytes.",
				"numeric": "The :attribute field must be at least :arg.",
				"string":  "The :attribute field must be at least :arg characters.",
			},
		},
		"min_digits": BasicMessageResolver("The :attribute field must have at least :arg digits."),
		// "missing":          BasicMessageResolver("The :attribute field must be missing."),
		// "missing_if":       BasicMessageResolver("The :attribute field must be missing when :other is :value."),
		// "missing_unless":   BasicMessageResolver("The :attribute field must be missing unless :other is :value."),
		// "missing_with":     BasicMessageResolver("The :attribute field must be missing when :args is present."),
		// "missing_with_all": BasicMessageResolver("The :attribute field must be missing when :args are present."),
		// "multiple_of":      BasicMessageResolver("The :attribute field must be a multiple of :value."),
		"not_nil":   BasicMessageResolver("The :attribute field must not be nil."),
		"not_in":    BasicMessageResolver("The selected :attribute is invalid."),
		"not_regex": BasicMessageResolver("The :attribute field format is invalid."),
		"numeric":   BasicMessageResolver("The :attribute field must be a number."),
		// "password": MessageHintResolver{Hints: map[string]string{
		// 	"letters":       "The :attribute field must contain at least one letter.",
		// 	"mixed":         "The :attribute field must contain at least one uppercase and one lowercase letter.",
		// 	"numbers":       "The :attribute field must contain at least one number.",
		// 	"symbols":       "The :attribute field must contain at least one symbol.",
		// 	"uncompromised": "The given :attribute has appeared in a data leak. Please choose a different :attribute.",
		// }},
		// "present":           BasicMessageResolver("The :attribute field must be present."),
		// "present_if":        BasicMessageResolver("The :attribute field must be present when :other is :value."),
		// "present_unless":    BasicMessageResolver("The :attribute field must be present unless :other is :value."),
		// "present_with":      BasicMessageResolver("The :attribute field must be present when :args is present."),
		// "present_with_all":  BasicMessageResolver("The :attribute field must be present when :args are present."),
		// "prohibited":        BasicMessageResolver("The :attribute field is prohibited."),
		// "prohibited_if":     BasicMessageResolver("The :attribute field is prohibited when :other is :value."),
		// "prohibited_unless": BasicMessageResolver("The :attribute field is prohibited unless :other is in :args."),
		// "prohibits":         BasicMessageResolver("The :attribute field prohibits :other from being present."),
		"regex":    BasicMessageResolver("The :attribute field format is invalid."),
		"required": BasicMessageResolver("The :attribute field is required."),
		// "required_array_keys":  BasicMessageResolver("The :attribute field must contain entries for: :args."),
		// "required_if":          BasicMessageResolver("The :attribute field is required when :other is :value."),
		// "required_if_accepted": BasicMessageResolver("The :attribute field is required when :other is accepted."),
		// "required_if_declined": BasicMessageResolver("The :attribute field is required when :other is declined."),
		// "required_unless":      BasicMessageResolver("The :attribute field is required unless :other is in :args."),
		// "required_with":        BasicMessageResolver("The :attribute field is required when :args is present."),
		// "required_with_all":    BasicMessageResolver("The :attribute field is required when :args are present."),
		// "required_without":     BasicMessageResolver("The :attribute field is required when :args is not present."),
		// "required_without_all": BasicMessageResolver("The :attribute field is required when none of :args are present."),
		// "same": BasicMessageResolver("The :attribute field must match :other."),
		"size": MessageHintResolver{
			Fallback: "The :attribute field must be of size :arg.",
			Hints: map[string]string{
				"array":   "The :attribute field must contain :arg items.",
				"file":    "The :attribute field must be :arg kilobytes.",
				"numeric": "The :attribute field must be :arg.",
				"string":  "The :attribute field must be :arg characters.",
			}},
		"starts_with": BasicMessageResolver("The :attribute field must start with one of the following: :args."),
		// "string":   BasicMessageResolver("The :attribute field must be a string."),
		// "timezone": BasicMessageResolver("The :attribute field must be a valid timezone."),
		// "unique":   BasicMessageResolver("The :attribute has already been taken."),
		// "uploaded": BasicMessageResolver("The :attribute failed to upload."),
		"uppercase": BasicMessageResolver("The :attribute field must be uppercase."),
		"url":       BasicMessageResolver("The :attribute field must be a valid URL."),
		"ulid":      BasicMessageResolver("The :attribute field must be a valid ULID."),
		"uuid":      BasicMessageResolver("The :attribute field must be a valid UUID."),
	})

	LogValidatorsWithoutMessages()
}

func Required(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	if !ctx.HasValue() {
		return "required", false
	}

	switch ctx.Kind() {
	case reflect.Chan, reflect.Func, reflect.UnsafePointer, reflect.Interface:
		if ctx.Value.IsNil() {
			return "required", false
		}
	case reflect.Map, reflect.Slice:
		if ctx.Value.IsNil() {
			return "required", false
		}
		if ctx.Value.Len() == 0 {
			return "required", false
		}
	case reflect.String:
		if ctx.Value.String() == "" {
			return "required", false
		}
	}

	return "", true
}

func NotNil(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	if !ctx.HasValue() {
		return "nil", false
	}

	switch ctx.Kind() {
	case reflect.Chan, reflect.Func, reflect.UnsafePointer, reflect.Interface, reflect.Map, reflect.Slice:
		if ctx.Value.IsNil() {
			return "nil", false
		}
	}

	return "", true
}

func Accepted(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	if !ctx.IsKind(
		reflect.Bool,
		reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
	) {
		return "invalid_type", false
	}

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

func Declined(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	if !ctx.IsKind(
		reflect.Bool,
		reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
	) {
		return "invalid_type", false
	}

	if !ctx.HasValue() {
		return "unacceptable", false
	}

	switch ctx.Kind() {
	case reflect.Bool:
		if ctx.Value.Bool() {
			return "unacceptable", false
		}
	case reflect.String:
		switch ctx.Value.String() {
		case "no", "off", "0", "false":
			return "", true
		}
		return "unacceptable", false
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if ctx.Value.Int() == 0 {
			return "", true
		}
		return "unacceptable", false
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if ctx.Value.Uint() == 0 {
			return "", true
		}
		return "unacceptable", false
	default:
		return "invalid_type", false
	}

	return "", true
}

func Boolean(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	if !ctx.IsKind(
		reflect.Bool,
		reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
	) {
		return "invalid_type", false
	}

	if !ctx.HasValue() {
		return "", true
	}

	switch ctx.Kind() {
	case reflect.Bool:
		return "", true
	case reflect.String:
		switch ctx.Value.String() {
		case "yes", "on", "1", "true":
			return "", true
		case "no", "off", "0", "false":
			return "", true
		}
		return "not_a_boolean", false
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch ctx.Value.Int() {
		case 0, 1:
			return "", true
		}
		return "not_a_boolean", false
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch ctx.Value.Uint() {
		case 0, 1:
			return "", true
		}
		return "not_a_boolean", false
	default:
		return "invalid_type", false
	}
}

func Ascii(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()
	str, status := ctx.StringLike()
	if !status.Oke() {
		return status.Response()
	}

	for _, c := range str {
		if c > 127 {
			return "not_ascii", false
		}
	}

	return "", true
}

func HexColor(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()
	strContent, status := ctx.String()
	if !status.Oke() {
		return status.Response()
	}

	switch len(strContent) {
	case 4, 5, 7, 9:
		// Continue
	default:
		return "invalid", false
	}

	if strContent[0] != '#' {
		return "invalid", false
	}
	for _, c := range strContent[1:] {
		if !(c >= '0' && c <= '9' || c >= 'a' && c <= 'f' || c >= 'A' && c <= 'F') {
			return "invalid", false
		}
	}

	return "", true
}

func URL(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()
	strContent, status := ctx.String()
	if !status.Oke() {
		return status.Response()
	}

	parsedUrl, err := url.ParseRequestURI(strContent)
	if err != nil {
		return "invalid", false
	}

	if len(ctx.Args) == 0 {
		return "", true
	}

	for _, acceptedSchema := range ctx.Args {
		if parsedUrl.Scheme == acceptedSchema {
			return "", true
		}
	}

	return "protocol", false
}

func Uppercase(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()
	strContent, status := ctx.StringLike()
	if !status.Oke() {
		return status.Response()
	}

	if strings.ToUpper(strContent) != strContent {
		return "not_uppercase", false
	}

	return "", true
}

func MacAddress(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	str, status := ctx.String()
	if !status.Oke() {
		return status.Response()
	}

	mac, err := net.ParseMAC(str)
	if err != nil || mac == nil {
		return "invalid", false
	}

	return "", true
}

func Lowercase(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()
	strContent, status := ctx.StringLike()
	if !status.Oke() {
		return status.Response()
	}

	if strings.ToLower(strContent) != strContent {
		return "not_lowercase", false
	}

	return "", true
}

func Ulid(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()
	strContent, status := ctx.String()
	if !status.Oke() {
		return status.Response()
	}

	_, err := ulid.Parse(strContent)
	if err != nil {
		return "invalid", false
	}

	return "", true
}

func Uuid(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	str, status := ctx.String()
	if !status.Oke() {
		return status.Response()
	}

	parsed, err := uuid.Parse(str)
	if err != nil {
		return "invalid", false
	}

	if len(ctx.Args) == 0 {
		return "", true
	}

	for _, allowedVersion := range ctx.Args {
		if len(allowedVersion) == 0 {
			continue
		}
		if allowedVersion[0] == 'v' {
			allowedVersion = allowedVersion[1:]
		}

		allowedVersionNumber, err := strconv.Atoi(allowedVersion)
		if err != nil {
			continue
		}

		if allowedVersionNumber <= 0 || allowedVersionNumber > 15 {
			continue
		}

		if uuid.Version(allowedVersionNumber) == parsed.Version() {
			return "", true
		}
	}

	return "version", false
}

func Numeric(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	switch ctx.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "", true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "", true
	case reflect.Float32, reflect.Float64:
		return "", true
	case reflect.String:
		str, status := ctx.String()
		if !status.Oke() {
			return status.Response()
		}

		_, err := strconv.Atoi(str)
		if err != nil {
			return "not_numeric", false
		}
		return "", true
	default:
		return "not_numeric", false
	}
}

func Max(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()
	if !ctx.IsNumeric() && !ctx.HasLen() {
		return "unsupported_type", false
	}

	if !ctx.HasValue() || len(ctx.Args) == 0 {
		return "", true
	}

	if ctx.Kind() == reflect.Float64 || ctx.Kind() == reflect.Float32 {
		maxLen, err := strconv.ParseFloat(ctx.Args[0], 64)
		if err != nil {
			return "", true
		}

		if ctx.Value.Float() <= maxLen {
			return "", true
		}

		return "numeric", false
	}

	maxLen, err := strconv.Atoi(ctx.Args[0])
	if err != nil {
		return "", true
	}

	if ctx.IsInt() {
		if ctx.Value.Int() > int64(maxLen) {
			return "numeric", false
		}
	} else if ctx.IsUint() {
		if ctx.Value.Uint() > uint64(maxLen) {
			return "numeric", false
		}
	} else if ctx.HasLen() {
		if ctx.Value.Len() > maxLen {
			if ctx.Kind() == reflect.String {
				return "string", false
			}
			return "array", false
		}
	} else {
		return "unsupported_type", false
	}

	return "", true
}

func Min(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()
	if !ctx.IsNumeric() && !ctx.HasLen() {
		return "unsupported_type", false
	}

	if !ctx.HasValue() || len(ctx.Args) == 0 {
		return "", true
	}

	if ctx.Kind() == reflect.Float64 || ctx.Kind() == reflect.Float32 {
		minLen, err := strconv.ParseFloat(ctx.Args[0], 64)
		if err != nil {
			return "", true
		}

		if ctx.Value.Float() >= minLen {
			return "", true
		}

		return "numeric", false
	}

	minLen, err := strconv.Atoi(ctx.Args[0])
	if err != nil {
		return "", true
	}

	if ctx.IsInt() {
		if ctx.Value.Int() < int64(minLen) {
			return "numeric", false
		}
	} else if ctx.IsUint() {
		if ctx.Value.Uint() < uint64(minLen) {
			return "numeric", false
		}
	} else if ctx.HasLen() {
		if ctx.Value.Len() < minLen {
			if ctx.Kind() == reflect.String {
				return "string", false
			}
			return "array", false
		}
	} else {
		return "unsupported_type", false
	}

	return "", true
}

func Between(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()
	if !ctx.IsNumeric() && !ctx.HasLen() {
		return "unsupported_type", false
	}

	if !ctx.HasValue() || len(ctx.Args) < 2 {
		return "", true
	}

	if ctx.Kind() == reflect.Float64 || ctx.Kind() == reflect.Float32 {
		min, err := strconv.ParseFloat(ctx.Args[0], 64)
		if err != nil {
			return "", true
		}
		max, err := strconv.ParseFloat(ctx.Args[1], 64)
		if err != nil {
			return "", true
		}

		val := ctx.Value.Float()
		if val >= min && val <= max {
			return "", true
		}

		return "numeric", false
	}

	min, err := strconv.Atoi(ctx.Args[0])
	if err != nil {
		return "", true
	}
	max, err := strconv.Atoi(ctx.Args[1])
	if err != nil {
		return "", true
	}

	if ctx.IsInt() {
		val := ctx.Value.Int()
		if val >= int64(min) && val <= int64(max) {
			return "", true
		}
		return "numeric", false
	} else if ctx.IsUint() {
		val := ctx.Value.Uint()
		if val >= uint64(min) && val <= uint64(max) {
			return "", true
		}
		return "numeric", false
	} else if ctx.HasLen() {
		val := ctx.Value.Len()
		if val >= min && val <= max {
			return "", true
		}
		if ctx.Kind() == reflect.String {
			return "string", false
		}
		return "array", false
	} else {
		return "unsupported_type", false
	}
}

func Size(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()
	if !ctx.IsInt() && !ctx.IsUint() && !ctx.HasLen() {
		return "unsupported_type", false
	}

	if len(ctx.Args) == 0 || !ctx.HasValue() {
		return "", true
	}

	size, err := strconv.Atoi(ctx.Args[0])
	if err != nil {
		return "", true
	}

	if ctx.IsInt() {
		if ctx.Value.Int() == int64(size) {
			return "", true
		}
		return "numeric", false
	} else if ctx.IsUint() {
		if ctx.Value.Uint() == uint64(size) {
			return "", true
		}
		return "numeric", false
	} else if ctx.HasLen() {
		if ctx.Value.Len() == size {
			return "", true
		}
		if ctx.Kind() == reflect.String {
			return "string", false
		}
		return "array", false
	} else {
		return "unsupported_type", false
	}
}

func StartsWith(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()
	str, status := ctx.StringLike()
	if !status.Oke() {
		return status.Response()
	}

	if len(ctx.Args) == 0 {
		return "", true
	}

	for _, allowedPrefix := range ctx.Args {
		if strings.HasPrefix(str, allowedPrefix) {
			return "", true
		}
	}

	return "starts_with", false
}

func EndsWith(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()
	str, status := ctx.StringLike()
	if !status.Oke() {
		return status.Response()
	}

	if len(ctx.Args) == 0 {
		return "", true
	}

	for _, allowedSuffix := range ctx.Args {
		if strings.HasSuffix(str, allowedSuffix) {
			return "", true
		}
	}

	return "ends_with", false
}

func JSON(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()
	str, status := ctx.StringLike()
	if !status.Oke() {
		return status.Response()
	}

	var jsonData any
	err := json.Unmarshal([]byte(str), &jsonData)
	if err != nil {
		return "json", false
	}

	return "", true
}

func IP(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()
	str, status := ctx.String()
	if !status.Oke() {
		return status.Response()
	}

	parsed := net.ParseIP(str)
	if parsed == nil {
		return "invalid", false
	}

	return "", true
}

func IPV4(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()
	str, status := ctx.String()
	if !status.Oke() {
		return "not_string", false
	}

	parsed := net.ParseIP(str)
	if parsed == nil {
		return "invalid", false
	}

	if len(parsed.To4()) != net.IPv4len {
		return "invalid", false
	}

	return "", true
}

func IPV6(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()
	str, status := ctx.String()
	if !status.Oke() {
		return status.Response()
	}

	parsed := net.ParseIP(str)
	if parsed == nil {
		return "invalid", false
	}

	if len(parsed.To4()) == net.IPv4len {
		return "invalid", false
	}

	return "", true
}

func Regex(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()
	str, status := ctx.String()
	if !status.Oke() {
		return "not_string", false
	}

	if len(ctx.Args) == 0 {
		return "", true
	}

	for _, pattern := range ctx.Args {
		if len(pattern) < 2 || !strings.HasPrefix(pattern, "/") || !strings.HasSuffix(pattern, "/") {
			continue
		}
		pattern = pattern[1 : len(pattern)-1]

		matched, err := regexp.MatchString(pattern, str)
		if err != nil {
			return "invalid", false
		}
		if matched {
			return "", true
		}
	}

	return "regex", false
}

func NotRegex(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()
	str, status := ctx.String()
	if !status.Oke() {
		return "not_string", false
	}

	if len(ctx.Args) == 0 {
		return "", true
	}

	for _, pattern := range ctx.Args {
		if len(pattern) < 2 || !strings.HasPrefix(pattern, "/") || !strings.HasSuffix(pattern, "/") {
			continue
		}
		pattern = pattern[1 : len(pattern)-1]

		matched, err := regexp.MatchString(pattern, str)
		if err != nil {
			return "invalid", false
		}
		if matched {
			return "matched", false
		}
	}

	return "", true
}

func In(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	if len(ctx.Args) == 0 || !ctx.HasValue() {
		return "", true
	}

	var val string
	var valSet bool

	switch ctx.Kind() {
	case reflect.String:
		val = ctx.Value.String()
		valSet = true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val = strconv.Itoa(int(ctx.Value.Int()))
		valSet = true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val = strconv.Itoa(int(ctx.Value.Uint()))
		valSet = true
	}

	if !valSet {
		return "not_in", false
	}

	for _, allowedValue := range ctx.Args {
		if val == allowedValue {
			return "", true
		}
	}

	return "not_in", false
}

func NotIn(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	if len(ctx.Args) == 0 || !ctx.HasValue() {
		return "", true
	}

	var val string
	var valSet bool

	switch ctx.Kind() {
	case reflect.String:
		val = ctx.Value.String()
		valSet = true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val = strconv.Itoa(int(ctx.Value.Int()))
		valSet = true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val = strconv.Itoa(int(ctx.Value.Uint()))
		valSet = true
	}

	if !valSet {
		return "", true
	}

	for _, allowedValue := range ctx.Args {
		if val == allowedValue {
			return "in", false
		}
	}

	return "", true
}

func Mimetypes(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()
	mimetype, status := ctx.String()
	if !status.Oke() {
		return status.Response()
	}

	if len(ctx.Args) == 0 {
		return "", true
	}

	for _, allowedMimetype := range ctx.Args {
		if strings.HasSuffix(allowedMimetype, "*") {
			// Do a prefix match
			if strings.HasPrefix(mimetype, allowedMimetype[:len(allowedMimetype)-1]) {
				return "", true
			}
			continue
		}

		if mimetype == allowedMimetype {
			return "", true
		}
	}

	return "mimetype", false
}

func Mimes(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()
	mimetype, status := ctx.String()
	if !status.Oke() {
		return status.Response()
	}

	if len(ctx.Args) == 0 {
		return "", true
	}

	for _, allowedExtension := range ctx.Args {
		allowedMimetype, ok := extensions[allowedExtension]
		if !ok {
			continue
		}

		if mimetype == allowedMimetype {
			return "", true
		}
	}

	return "mimetype", false
}

type DigitsStatus uint8

const (
	DigitStatusValid DigitsStatus = iota
	DigitStatusInvalidType
	DigitStatusNil
	DigitStatusStringWithoutDigits
)

func (s DigitsStatus) Oke() bool {
	return s == DigitStatusValid
}

func (s DigitsStatus) Response() (string, bool) {
	switch s {
	case DigitStatusInvalidType:
		return "invalid_type", false
	case DigitStatusStringWithoutDigits:
		return "string_without_digits", false
	default:
		return "", true
	}
}

func digits(ctx *ValidatorCtx) (int, DigitsStatus) {
	ctx.UnwrapPointer()

	if !ctx.IsKind(
		reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
	) {
		return 0, DigitStatusInvalidType
	}

	if !ctx.HasValue() {
		return 0, DigitStatusNil
	}

	var strValue string
	switch ctx.Kind() {
	case reflect.String:
		// Check if the string is a valid number
		strValue = ctx.Value.String()
		_, err := strconv.Atoi(strValue)
		if err != nil {
			return 0, DigitStatusStringWithoutDigits
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value := ctx.Value.Int()
		strValue = strconv.FormatInt(value, 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value := ctx.Value.Uint()
		strValue = strconv.FormatUint(value, 10)
	case reflect.Float32, reflect.Float64:
		value := ctx.Value.Float()
		strValue = strconv.FormatFloat(value, 'f', -1, 64)
	default:
		return 0, DigitStatusInvalidType
	}

	count := 0
	for _, char := range strValue {
		if char >= '0' && char <= '9' {
			count++
		} else if char == '.' || char == ',' || char == 'E' {
			break
		}
	}

	return 0, DigitStatusValid
}

func MinDigits(ctx *ValidatorCtx) (string, bool) {
	count, status := digits(ctx)
	if !status.Oke() {
		return status.Response()
	}

	if len(ctx.Args) == 0 {
		return "", true
	}

	min, err := strconv.Atoi(ctx.Args[0])
	if err != nil {
		return "", true
	}

	if count < min {
		return "min", false
	}

	return "", true
}

func MaxDigits(ctx *ValidatorCtx) (string, bool) {
	count, status := digits(ctx)
	if !status.Oke() {
		return status.Response()
	}

	if len(ctx.Args) == 0 {
		return "", true
	}

	max, err := strconv.Atoi(ctx.Args[0])
	if err != nil {
		return "", true
	}

	if count > max {
		return "max", false
	}

	return "", true
}

func Digits(ctx *ValidatorCtx) (string, bool) {
	count, status := digits(ctx)
	if !status.Oke() {
		return status.Response()
	}

	if len(ctx.Args) == 0 {
		return "", true
	}

	digits, err := strconv.Atoi(ctx.Args[0])
	if err != nil {
		return "", true
	}

	if count != digits {
		return "digits", false
	}

	return "", true
}

func DigitsBetween(ctx *ValidatorCtx) (string, bool) {
	count, status := digits(ctx)
	if !status.Oke() {
		return status.Response()
	}

	if len(ctx.Args) < 2 {
		return "", true
	}

	min, err := strconv.Atoi(ctx.Args[0])
	if err != nil {
		return "", true
	}

	max, err := strconv.Atoi(ctx.Args[0])
	if err != nil {
		return "", true
	}

	if count < min || count > max {
		return "between", false
	}

	return "", true
}

func AfterDate(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	fieldDate, status := ctx.Date()
	if !status.Oke() {
		return status.Response()
	}

	if len(ctx.Args) == 0 {
		return "", true
	}

	argDate, ok := ctx.DateFromArgs(0)
	if !ok {
		return "invalid_param", false
	}

	if fieldDate.After(argDate) {
		return "", true
	}

	return "after", false
}

func AfterOrEqualDate(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	fieldDate, status := ctx.Date()
	if !status.Oke() {
		return status.Response()
	}

	if len(ctx.Args) == 0 {
		return "", true
	}

	argDate, ok := ctx.DateFromArgs(0)
	if !ok {
		return "invalid_param", false
	}

	if fieldDate.Equal(argDate) || fieldDate.After(argDate) {
		return "", true
	}

	return "after", false
}

func BeforeDate(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	fieldDate, status := ctx.Date()
	if !status.Oke() {
		return status.Response()
	}

	if len(ctx.Args) == 0 {
		return "", true
	}

	argDate, ok := ctx.DateFromArgs(0)
	if !ok {
		return "invalid_param", false
	}

	if fieldDate.Before(argDate) {
		return "", true
	}

	return "before", false
}

func BeforeOrEqualDate(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	fieldDate, status := ctx.Date()
	if !status.Oke() {
		return status.Response()
	}

	if len(ctx.Args) == 0 {
		return "", true
	}

	argDate, ok := ctx.DateFromArgs(0)
	if !ok {
		return "invalid_param", false
	}

	if fieldDate.Equal(argDate) || fieldDate.Before(argDate) {
		return "", true
	}

	return "before", false
}

func Email(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	str, status := ctx.String()
	if !status.Oke() {
		return status.Response()
	}

	email, err := mail.ParseAddress(str)
	if err != nil {
		return "invalid", false
	}

	checkDns := false
	allowName := false
	requireName := false
	noLocalhost := false
	for _, arg := range ctx.Args {
		switch arg {
		case "dns":
			checkDns = true
		case "allow_name":
			allowName = true
		case "require_name":
			requireName = true
		case "no_localhost":
			noLocalhost = true
		}
	}

	host := strings.Split(strings.Split(email.Address, "@")[1], ":")[0]
	if checkDns {
		_, err := net.LookupIP(host)
		if err != nil {
			return "dns", false
		}
	}

	if noLocalhost {
		hostParts := strings.Split(host, ".")
		if strings.ToLower(hostParts[len(hostParts)-1]) == "localhost" {
			return "no_localhost", false
		}
		parsedIp := net.ParseIP(host)
		if parsedIp != nil {
			switch parsedIp.String() {
			case "127.0.0.1", "::1", "0.0.0.0", "::":
				return "no_localhost", false
			}
		}
	}

	if requireName {
		if email.Name == "" {
			return "invalid", false
		}
	} else if !allowName && email.Name != "" {
		return "invalid", false
	}

	return "", true
}

func Filled(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	if !ctx.HasValue() {
		return "", true
	}

	switch ctx.Kind() {
	case reflect.Map, reflect.Slice:
		if ctx.Value.IsNil() {
			return "", true
		}
		if ctx.Value.Len() == 0 {
			return "required", false
		}
	case reflect.String:
		if ctx.Value.String() == "" {
			return "required", false
		}
	}

	return "", true
}

func Date(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	_, status := ctx.Date()
	if !status.Oke() {
		return status.Response()
	}

	return "", true
}

func DateFormat(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	str, status := ctx.String()
	if !status.Oke() {
		return status.Response()
	}

	if len(ctx.Args) == 0 {
		return "", true
	}

	for _, format := range ctx.Args {
		t, err := time.Parse(format, str)
		if err != nil {
			ctx.SetState("parsed_date", t)
			return "", true
		}
	}

	return "format", false
}

func Extensions(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	str, status := ctx.String()
	if !status.Oke() {
		return status.Response()
	}

	if len(ctx.Args) == 0 {
		return "", true
	}

	if str == "" {
		return "extension", false
	}

	for _, allowedExtension := range ctx.Args {
		if allowedExtension == "" {
			continue
		}

		if allowedExtension[0] != '.' {
			allowedExtension = "." + allowedExtension
		}
		if allowedExtension == "" {
			continue
		}

		if strings.HasSuffix(str, allowedExtension) {
			return "", true
		}
	}

	return "extension", false
}

func validateAlpha(ctx *ValidatorCtx, numeric bool, dashes bool) (string, bool) {
	ctx.UnwrapPointer()
	str, status := ctx.StringLike()
	if !status.Oke() {
		return status.Response()
	}

	assci := false
	for _, arg := range ctx.Args {
		if arg == "ascii" {
			assci = true
			break
		}
	}

	if assci {
		for _, c := range str {
			if c >= 'A' && c <= 'Z' {
				continue
			}
			if c >= 'a' && c <= 'z' {
				continue
			}
			if numeric && c >= '0' && c <= '9' {
				continue
			}
			if dashes && (c == '-' || c == '_') {
				continue
			}
			return "invalid", false
		}
	}

	for _, c := range str {
		if unicode.IsLetter(c) {
			continue
		}
		if numeric && unicode.IsNumber(c) {
			continue
		}
		if dashes && (c == '-' || c == '_') {
			continue
		}
		return "invalid", false
	}

	return "", true
}

func Alpha(ctx *ValidatorCtx) (string, bool) {
	return validateAlpha(ctx, false, false)
}

func AlphaDash(ctx *ValidatorCtx) (string, bool) {
	return validateAlpha(ctx, true, true)
}

func AlphaNumeric(ctx *ValidatorCtx) (string, bool) {
	return validateAlpha(ctx, true, false)
}

func ActiveUrl(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	str, status := ctx.String()
	if !status.Oke() {
		return status.Response()
	}

	parsedUrl, err := url.ParseRequestURI(str)
	if err != nil {
		return "invalid", false
	}

	_, err = net.LookupIP(parsedUrl.Host)
	if err != nil {
		return "dns", false
	}

	return "", true
}

func Bail(ctx *ValidatorCtx) (string, bool) {
	ctx.Bail()
	return "", true
}

func Confirmed(ctx *ValidatorCtx) (string, bool) {
	ctx.UnwrapPointer()

	fieldName := ctx.ObjectFieldName()
	if fieldName == "" {
		return "field_not_in_struct", false
	}

	if !ctx.HasValue() {
		return "", true
	}

	toCompare := ctx.Field("." + fieldName + "Confirmation")
	if toCompare == nil || toCompare.Value == nil {
		return "field_missing", false
	}

	if !equal(*ctx.Value, *toCompare.Value) {
		return "not_equal", false
	}

	return "", true
}

type SizeCompareStatus uint8

const (
	SizeCompareStatusEq SizeCompareStatus = iota
	SizeCompareStatusLt
	SizeCompareStatusGt
)

func compareFieldsBase(ctx *ValidatorCtx) (SizeCompareStatus, ConvertStatus) {
	if len(ctx.Args) == 0 {
		return 0, Invalid
	}

	ctx.UnwrapPointer()

	if !ctx.HasValue() {
		return 0, ValueNil
	}

	other := ctx.Field(ctx.Args[0])
	if other == nil {
		return 0, Invalid
	}

	other.UnwrapPointer()

	if !other.HasValue() {
		return 0, Invalid
	}

	if ctx.IsNumeric() || other.IsNumeric() {
		if !ctx.IsNumeric() || !other.IsNumeric() {
			return 0, InvalidType
		}

		if ctx.IsFloat() || other.IsFloat() {
			aValue, aOk := ctx.Float64()
			bValue, bOk := other.Float64()
			if !aOk || !bOk {
				return 0, InvalidType
			}

			if aValue == bValue {
				return SizeCompareStatusEq, ConverstionOk
			}
			if aValue < bValue {
				return SizeCompareStatusLt, ConverstionOk
			}
			return SizeCompareStatusGt, ConverstionOk
		}

		if ctx.IsUint() && other.IsUint() {
			aValue := ctx.Value.Uint()
			bValue := other.Value.Uint()

			if aValue == bValue {
				return SizeCompareStatusEq, ConverstionOk
			}
			if aValue < bValue {
				return SizeCompareStatusLt, ConverstionOk
			}
			return SizeCompareStatusGt, ConverstionOk
		}

		aValue, aOk := ctx.Int64()
		bValue, bOk := other.Int64()
		if !aOk || !bOk {
			return 0, InvalidType
		}

		if aValue == bValue {
			return SizeCompareStatusEq, ConverstionOk
		}
		if aValue < bValue {
			return SizeCompareStatusLt, ConverstionOk
		}
		return SizeCompareStatusGt, ConverstionOk
	}

	compareLen := false
	if ctx.IsList() || other.IsList() {
		if !ctx.IsList() || !other.IsList() {
			return 0, InvalidType
		}
		compareLen = true
	}

	if ctx.Kind() == reflect.String && other.Kind() == reflect.String {
		compareLen = true
	}

	if compareLen {
		aLen := ctx.Value.Len()
		bLen := other.Value.Len()
		if aLen == bLen {
			return SizeCompareStatusEq, ConverstionOk
		}
		if aLen < bLen {
			return SizeCompareStatusLt, ConverstionOk
		}
		return SizeCompareStatusGt, ConverstionOk
	}

	return 0, InvalidType
}

func Gt(ctx *ValidatorCtx) (string, bool) {
	sizeStatus, status := compareFieldsBase(ctx)
	if !status.Oke() {
		return status.Response()
	}

	if sizeStatus != SizeCompareStatusGt {
		return "lt", false
	}

	return "", true
}

func Gte(ctx *ValidatorCtx) (string, bool) {
	sizeStatus, status := compareFieldsBase(ctx)
	if !status.Oke() {
		return status.Response()
	}

	if sizeStatus == SizeCompareStatusLt {
		return "lt", false
	}

	return "", true
}

func Lt(ctx *ValidatorCtx) (string, bool) {
	sizeStatus, status := compareFieldsBase(ctx)
	if !status.Oke() {
		return status.Response()
	}

	if sizeStatus != SizeCompareStatusLt {
		return "gt", false
	}

	return "", true
}

func Lte(ctx *ValidatorCtx) (string, bool) {
	sizeStatus, status := compareFieldsBase(ctx)
	if !status.Oke() {
		return status.Response()
	}

	if sizeStatus == SizeCompareStatusGt {
		return "gt", false
	}

	return "", true
}
