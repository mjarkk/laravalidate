# Rules

### `accepted`

The field under validation must be "yes", "on", 1, "1", true, or "true".
This is useful for validating "Terms of Service" acceptance or similar fields.

```go
type Body struct {
	// The terms must be accepted
	Terms1 bool `validate:"accepted"`
	Terms2 string `validate:"accepted"`
	Terms3 int `validate:"accepted"`
}
```

### `active_url`

The field under validation must be a valid URL according to [url.ParseRequestURI](https://pkg.go.dev/net/url#ParseRequestURI) and give a valid response for [net.LookupIP](https://pkg.go.dev/net#LookupIP).

### `after:date`

The field under validation must be a value after a given date.

Valid (string) field values an be found under: [Valid field datetime values](#valid-field-datetime-values)

Valid arguments for validator can be found under: [Valid argument datetime values](#valid-argument-datetime-values)

```go
type Body struct {
	StartDate string `validation:"date|after:tomorrow"`
}
```

**TODO**

Refer to another field

### `after_or_equal:date`

The field under validation must be a value after or equal to the given date.

This validator functions the same as the `after` validator.

### `alpha`

The field under validation must be entirely alphabetic characters.

Validated using Go's [unicode.IsLetter](https://pkg.go.dev/unicode#IsLetter)

To restrict the field to only ASCII characters, use the `ascii` argument:

```go
type Body struct {
	Field string `validate:"alpha:ascii"`
}
```

### `alpha_dash`

The field under validation may have alpha-numeric characters, as well as ASCII dashes (`-`) and ASCII underscores (`_`).

Validated using Go's [unicode.IsLetter](https://pkg.go.dev/unicode#IsLetter) and [unicode.IsNumber](https://pkg.go.dev/unicode#IsDigit)

To restrict the field to only ASCII characters, use the `ascii` argument:

```go
type Body struct {
	Field string `validate:"alpha_dash:ascii"`
}
```

### `alpha_numeric`

The field under validation may have alpha-numeric characters, as well as ASCII dashes and ASCII underscores.

Validated using Go's [unicode.IsLetter](https://pkg.go.dev/unicode#IsLetter) and [unicode.IsNumber](https://pkg.go.dev/unicode#IsDigit)

To restrict the field to only ASCII characters, use the `ascii` argument:

```go
type Body struct {
	Field string `validate:"alpha_numeric:ascii"`
}
```

### `ascii`

The field under validation must be entirely 7-bit ASCII characters.

Accepted types are: `string`, `[]byte`, `byte`, `[]rune` and `rune`.

### `bail`

Stop running validation rules after the first validation failure.

### `before:date`

The field under validation must be a value preceding the given date.

This validator functions the same as the `after` validator.

### `before_or_equal:date`

The field under validation must be a value preceding or equal to the given date.

This validator functions the same as the `after` validator.

### `between:min,max`

The field under validation must have a size between the given min and max (inclusive).
Strings, numerics, arrays, and files are evaluated in the same fashion as the size rule.

### `boolean`

The field under validation must be able to be cast as a boolean.

Accepted input are `true`, `false`, `1`, `0`, `"1"`, and `"0"`.

### `confirmed`

The field under validation must have a matching field with the `..Confirmation` suffix.
For example, if the field under validation is `Password`, a matching `PasswordConfirmation` field must be present in the input.

Note that the field name only applies to the Go struct field name, not the JSON tag so you can name the JSON tag differently.

### `date`

The field under validation must be a, non-relative date.

Valid (string) field values an be found under: [Valid field datetime values](#valid-field-datetime-values)

### `date_format:2006-01-02 15:04:05,..`

The field under validation must match the given date time format.

The format is parsed using Go's [time.Parse](https://pkg.go.dev/time#Parse) function.

If the date is parsed successfully, the date is cached and is reused by date validators after this validator.

### `declined`

The field under validation must be "no", "off", 0, "0", false, or "false".

### `digits:value`

The field under validation must be numeric and must have an exact length of value.

### `digits_between:min,max`

The field under validation must be numeric and must have a length between the given min and max.

### `email:flag,flag,..`

The field under validation must be formatted as an e-d address.

The field is validated using Go's [mail.ParseAddress](https://pkg.go.dev/net/mail#ParseAddress).

By default the name property of an email is disallowed like `Example <example@example.org>`.

Flags:

- `dns` - Lookup the domain name and check if there are any IP's attached to it.
- `allow_name` - Allow the name property of an email address like `Example <example@example.org>`.
- `require_name` - Require the name property of an email address
- `no_localhost` - Disallow localhost addresses `localhost`, `127.0.0.1`, `::1`, `0.0.0.0` and `::` _(ipv6 version of `0.0.0.0`)_

### `ends_with:foo,bar,...`

The field under validation must end with one of the given values.

### `exists`

[Requires dbrules to be setup!](./README.md#database-rules)



### `extensions:jpg,png,...`

The file under validation must have a matching extension.

### `filled`

The field under validation must not be empty when it is present.

This validator is almost equal to the `required` validator except that it allows nil values.

### `hex_color`

The field under validation must contain a valid color value in [hexadecimal](https://developer.mozilla.org/en-US/docs/Web/CSS/hex-color) format.

### `in:foo,bar,...`

The field under validation must be included in the given list of values.

### `ip`

The field under validation must be an IP address.

### `ipv4`

The field under validation must be an IPv4 address.

### `ipv6`

The field under validation must be an IPv6 address.

### `json`

The field under validation must be a valid JSON `string` or `[]byte`.

### `lowercase`

The field under validation must be lowercase.

### `mac_address`

The field under validation must be a MAC address.

The value is parsesed using Go's [net.ParseMAC](https://pkg.go.dev/net#ParseMAC)

### `max:value`

The field under validation must be less than or equal to a maximum value. Strings, numerics, arrays, and files are evaluated in the same fashion as the size rule.

### `max_digits:value`

The field under validation must have at most the given number of digits.

### `mimetypes:text/plain,...`

The field under validation must match one of the given MIME types.

### `mimes:jpg,png,...`

The file under validation must have a MIME type corresponding to one of the listed extensions.

A full listing of MIME types and their corresponding extensions may be found at the following location:

https://svn.apache.org/repos/asf/httpd/httpd/trunk/docs/conf/mime.types

### `min:value`

The field under validation must have a minimum value. Strings, numerics, arrays, and files are evaluated in the same fashion as the size rule.

### `min_digits:value`

The field under validation must have at least the given number of digits.

### `not_nil`

The field under validation must not be nil.

### `not_in:foo,bar,...`

The field under validation must not be included in the given list of values.

### `not_regex:pattern`

The field under validation must not match the given regular expression(s).

Examples:

```go
type Body struct {
	// Email can't end with @gmail.com or @outlook.com
	Email string `validate:"not_regex:/@gmail.com$/,/@outlook.com$/"`
}
```

### `numeric`

The field under validation must be numeric.

The value is validated using Go's [strconv.Atoi](https://pkg.go.dev/strconv#Atoi)

### `regex:pattern`

The field under validation must match the given regular expression(s).

Examples:

```go
type Body struct {
	// Email must end with @gmail.com or @outlook.com
	Email string `validate:"regex:/@gmail.com$/,/@outlook.com$/"`
}
```

### `required`

The field under validation must be present in the input data and not empty.
A field is "empty" if it meets one of the following criteria:

- The value is nil.
- The value is an empty string.
- The value is an empty slice or map.

**TODO**

- The value is an uploaded file with no path.

### `size:value`

The field under validation must have a size matching the given value. For string data, value corresponds to the number of characters. For numeric data, value corresponds to a given integer value (the attribute must also have the numeric or integer rule). For an array, size corresponds to the count of the array.
For files, size corresponds to the file size in kilobytes.

Let's look at some examples:

```go
type Body struct {
	// The title must be exactly 12 characters
	Title string `validate:"size:12"`
	// The seats must be exactly 4
	Seats int `validate:"size:4"`
	// The tags must have exactly 3 elements
	Tags []string `validate:"size:3"`
}
```

### `starts_with:foo,bar,...`

The field under validation must start with one of the given values.

### `uppercase`

The field under validation must be uppercase.

### `url`

The field under validation must be a valid URL.

The value is validated using Go's [url.Parse](https://pkg.go.dev/net/url#Parse) function.

If you would like to specify the URL protocols that should be considered valid,
you may pass the protocols as validation rule parameters:

```go
type Body struct {
	ImageUrl string `validate:"url:http,https"`
	gameUrl string `validate:"url:minecraft,steam"`
}
```

## `ulid`

The field under validation must be a valid Universally [Unique Lexicographically Sortable Identifier](https://github.com/ulid/spec) (ULID).

The values are validated using [github.com/oklog/ulid/v2](https://pkg.go.dev/github.com/oklog/ulid/v2)

## `uuid`

The field under validation must be a valid RFC 9562 and DCE 1.1 (version 1 to 7) universally unique identifier (UUID).

The values are validated using [github.com/google/uuid](https://pkg.go.dev/github.com/google/uuid)

If you would like to specify the UUID versions that should be considered valid,
you may pass the versions as validation rule parameters:

```go
type Body struct {
	Uuid string `validate:"uuid:v4"`
	OldUuid string `validate:"uuid:v1,v2,v4"`
}
```

## Valid field datetime values

Here are valid datetime field values

_Note that RFC3339(nano) is compatible with most ISO8601 strings_

| Format        | Example                            |
| ------------- | ---------------------------------- |
| `RFC3339`     | `2024-10-02T12:30:36+02:00`        |
| `RFC3339Nano` | `2024-10-02T12:30:36.808963+02:00` |
| `DateTime`    | `2024-10-02 12:30:36`              |
| `DateOnly`    | `2024-10-02`                       |
| `Unix`        | `1633157436`                       |
| `UnixMilli`   | `1633157436808`                    |
| `UnixMicro`   | `1633157436808963`                 |
| `UnixNano`    | `1633157436808963000`              |

## Valid argument datetime values

Here are valid datetime validation argument values

_Note that RFC3339(nano) is compatible with most ISO8601 strings_

| Format                                | Description                                                                                                | Example                                                 |
| ------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ------------------------------------------------------- |
| `2006-01-02T15:04:05.999999999Z07:00` | Parsed using Go's RFC3339Nano                                                                              | `2024-10-02T12:30:36.808963+02:00`                      |
| `2006-01-02T15:04:05Z07:00`           | Parsed using Go's RFC3339                                                                                  | `2024-10-02T12:30:36+02:00`                             |
| `2006-01-02 15:04:05`                 | Parsed using Go's DateTime                                                                                 | `2024-10-02 12:30:36`                                   |
| `2006-01-02`                          | Parsed using Go's DateOnly                                                                                 | `2024-10-02`                                            |
| `Unix`                                | Unix                                                                                                       | `1633157436`                                            |
| `UnixMilli`                           | UnixMilli                                                                                                  | `1633157436808`                                         |
| `UnixMicro`                           | UnixMicro                                                                                                  | `1633157436808963`                                      |
| `UnixNano`                            | UnixNano                                                                                                   | `1633157436808963000`                                   |
| `yesterday`                           | Midnight of yesterday                                                                                      | `yesterday`                                             |
| `midnight`                            | The time is set to 00:00:00                                                                                | `midnight`                                              |
| `today`                               | The time is set to 00:00:00                                                                                | `today`                                                 |
| `now`                                 | Nothing is changed                                                                                         | `now`                                                   |
| `noon`                                | The time is set to 12:00:00                                                                                | `noon`                                                  |
| `tomorrow`                            | Midnight of tomorrow                                                                                       | `tomorrow`                                              |
| `[weekday]`                           | Midnight the specified weekday                                                                             | `friday`                                                |
| `back of [hour]`                      | 15 minutes past the specified hour                                                                         | `back of 7pm` (7:15 today) `back of 16` (16:15 today)   |
| `front of [hour]`                     | 15 minutes before the specified hour                                                                       | `front of 7pm` (6:45 today) `front of 15` (14:45 today) |
| `first day of [month] [year]`         | The first day of a specified month and year                                                                | `first day of January 2024`                             |
| `last day of [month] [year]`          | The last day of a specified month and year                                                                 | `last day of january 2024`                              |
| `first [weekday] of [month] [year]`   | The first specified weekday of a specified month and year                                                  | `first Thursday of January 2024`                        |
| `last [weekday] of [month] [year]`    | The last specified weekday of a specified month and year                                                   | `last Fri of Jan 2024`                                  |
| `[number] [unit]`                     | Create a new date relative to now, supported units: seconds,minutes,hours,days,weeks,weekdays,months,years | `+1 days` `4 weeks` `-3 hours`                          |
| `[weekday] [textWeekOffset] week`     | A specified Weekday at the `last`/`this`/`next` week                                                       | `Wednesday last week`                                   |
