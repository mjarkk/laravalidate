package dates

import (
	"strconv"
	"strings"
	"time"
)

func Hour(input string) (int, bool) {
	isPm := false
	if len(input) > 2 {
		if strings.HasSuffix(input, "am") || strings.HasSuffix(input, "AM") {
			input = input[:len(input)-2]
		} else if strings.HasSuffix(input, "pm") || strings.HasSuffix(input, "PM") {
			isPm = true
			input = input[:len(input)-2]
		}
	}

	nr, err := strconv.Atoi(input)
	if err != nil {
		return -1, false
	}
	if nr < 0 {
		return -1, false
	}
	if isPm {
		nr += 12
	}

	return nr, nr <= 24
}

func Year(input string) (int, bool) {
	nr, err := strconv.Atoi(input)
	if err != nil {
		return -1, false
	}
	if nr < 1000 || nr > 3000 {
		return -1, false
	}

	return nr, true
}

func Month(input string) (time.Month, bool) {
	switch strings.ToLower(input) {
	case "january", "jan":
		return time.January, true
	case "february", "feb":
		return time.February, true
	case "march", "mar":
		return time.March, true
	case "april", "apr":
		return time.April, true
	case "may":
		return time.May, true
	case "june", "jun":
		return time.June, true
	case "july", "jul":
		return time.July, true
	case "august", "aug":
		return time.August, true
	case "september", "sep":
		return time.September, true
	case "october", "oct":
		return time.October, true
	case "november", "nov":
		return time.November, true
	case "december", "dec":
		return time.December, true
	default:
		return time.Month(0), false
	}
}

func Weekday(input string) (time.Weekday, bool) {
	switch strings.ToLower(input) {
	case "sunday", "sun":
		return time.Sunday, true
	case "monday", "mon":
		return time.Monday, true
	case "tuesday", "tue":
		return time.Tuesday, true
	case "wednesday", "wed":
		return time.Wednesday, true
	case "thursday", "thu", "thur":
		return time.Thursday, true
	case "friday", "fri":
		return time.Friday, true
	case "saturday", "sat":
		return time.Saturday, true
	default:
		return time.Weekday(-1), false
	}
}

func ParseUnix(t int64) (time.Time, bool) {
	if t < 100_000_000 {
		return time.Time{}, false
	}
	if t < 10_000_000_000 {
		return time.Unix(t, 0), true
	}
	if t < 10_000_000_000_000 {
		return time.UnixMilli(t), true
	}
	if t < 10_000_000_000_000_000 {
		return time.UnixMicro(t), true
	}
	return time.UnixMicro(t / 1000), true
}

func ParseStructuredDate(input string) (time.Time, bool) {
	layoutsToAttempt := []string{
		time.RFC3339Nano,
		time.RFC3339,
		time.DateTime,
		time.DateOnly,
	}
	for _, layout := range layoutsToAttempt {
		t, err := time.Parse(layout, input)
		if err == nil {
			return t, true
		}
	}

	t, err := strconv.ParseInt(input, 10, 64)
	if err == nil {
		date, ok := ParseUnix(t)
		if ok {
			return date, true
		}
	}

	return time.Time{}, false
}
