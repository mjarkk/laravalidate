package laravalidate

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/mjarkk/laravalidate/dates"
)

const (
	ParsedDateKey = "parsed-date" // Stores a custom parsed date, this is used by the date_format validator to store the results so they can be used by other validators
)

type ValidatorCtx struct {
	// Needle contains the field under validation
	Needle
	// Args are the arguments that were passed to the validator
	Args []string
	// If bail indicates that the validator should stop after the first error
	ctx context.Context
	// State is a object that lives trough the validation process of a single field
	// The ValidatorCtx is regenerated for each field and validator
	state *ValidatorCtxState
	// lastObtainedField should contain the last field requested using the (*ValidatorCtx).Field(..) method
	// Can be nil if no field was requested during the validation
	lastObtainedField *Needle
}

type ValidatorCtxState struct {
	bail      bool
	state     map[string]any
	stack     Stack
	validator *Validator
}

// NewValidatorCtx returns the underlying context of the validator
func (ctx *ValidatorCtx) Context() context.Context {
	return ctx.ctx
}

// Date tries to convert the value to a time.Time
func (ctx *ValidatorCtx) Date() (time.Time, ConvertStatus) {
	state, ok := ctx.GetState(ParsedDateKey)
	if ok {
		t, ok := state.(time.Time)
		if ok {
			return t, ConverstionOk
		}
	}

	return ctx.Needle.Date()
}

func (ctx *ValidatorCtx) DateFromArgs(argIndex int) (time.Time, bool) {
	if len(ctx.Args) <= argIndex {
		return time.Time{}, false
	}

	argWords := strings.Split(ctx.Args[argIndex], " ")
	for i := len(argWords) - 1; i >= 0; i-- {
		word := argWords[i]
		if word == "" {
			argWords = append(argWords[:i], argWords[i+1:]...)
		}
		argWords[i] = strings.ToLower(word)
	}

	now := time.Now()
	today := func(hour, minute, seconds int) time.Time {
		return time.Date(now.Year(), now.Month(), now.Day(), hour, minute, seconds, 0, nil)
	}
	nextWeekday := func(base time.Time, day time.Weekday) time.Time {
		if base.Weekday() == day {
			return base
		}

		diff := day - base.Weekday()
		if diff < 0 {
			diff += 7
		}
		return base.AddDate(0, 0, int(diff))
	}
	nextWeekdayFromToday := func(day time.Weekday) time.Time {
		return nextWeekday(today(0, 0, 0), day)
	}
	previousWeekday := func(base time.Time, day time.Weekday) time.Time {
		if base.Weekday() == day {
			return base
		}

		diff := day - base.Weekday()
		if diff > 0 {
			diff -= 7
		}
		return base.AddDate(0, 0, int(diff))
	}

	switch len(argWords) {
	case 1:
		switch argWords[0] {
		case "yesterday":
			return today(0, 0, 0).AddDate(0, 0, -1), true
		case "midnight":
			return today(0, 0, 0), true
		case "today":
			return today(0, 0, 0), true
		case "now":
			return now, true
		case "noon":
			return today(12, 0, 0), true
		case "tomorrow":
			return today(0, 0, 0).AddDate(0, 0, 1), true
		default:
			weekday, ok := dates.Weekday(argWords[0])
			if ok {
				return nextWeekdayFromToday(weekday), true
			}
		}
	case 2:
		offset, offsetErr := strconv.Atoi(argWords[0])
		if offsetErr == nil && offset != 0 {
			switch argWords[1] {
			case "seconds", "second":
				return now.Add(time.Duration(offset) * time.Second), true
			case "minutes", "minute":
				return now.Add(time.Duration(offset) * time.Minute), true
			case "hours", "hour":
				return now.Add(time.Duration(offset) * time.Hour), true
			case "days", "day":
				return now.AddDate(0, 0, offset), true
			case "weeks", "week":
				return now.AddDate(0, 0, offset*7), true
			case "weekdays", "weekday":
				resp := now
				// FIXME: This can be done way more efficient
				if offset > 0 {
					for i := 0; i < offset; i++ {
						if resp.Weekday() == time.Friday {
							resp = resp.AddDate(0, 0, 3)
						} else {
							resp = resp.AddDate(0, 0, 1)
						}
					}
				} else {
					for i := 0; i < offset; i++ {
						if resp.Weekday() == time.Monday {
							resp = resp.AddDate(0, 0, -3)
						} else {
							resp = resp.AddDate(0, 0, -1)
						}
					}
				}
				return now, true
			case "months", "month":
				return now.AddDate(0, offset, 0), true
			case "years", "year":
				return now.AddDate(offset, 0, 0), true
			}
		}
	case 3:
		hour, ok := dates.Hour(argWords[2])
		if ok {
			switch strings.Join(argWords[:2], " ") {
			case "back of":
				if hour == 24 {
					return today(0, 15, 0).AddDate(0, 0, 1), true
				}
				return today(hour, 15, 0), true
			case "front of":
				if hour == 0 {
					return today(23, 45, 0).AddDate(0, 0, -1), true
				}
				return today(hour-1, 45, 0), true
			}
		}
		weekDay, ok := dates.Weekday(argWords[0])
		if ok && argWords[2] == "week" {
			switch argWords[1] {
			case "last", "previous", "prev":
				return today(0, 0, 0).AddDate(0, 0, -int(now.Weekday())-7+int(weekDay)), true
			case "this":
				return today(0, 0, 0).AddDate(0, 0, -int(now.Weekday())+int(weekDay)), true
			case "next":
				return nextWeekday(now, weekDay).AddDate(0, 0, -int(now.Weekday())+7+int(weekDay)), true
			}
		}
	case 5:
		month, monthOk := dates.Month(argWords[3])
		year, yearOk := dates.Year(argWords[4])
		weekday, weekdayOk := dates.Weekday(argWords[0])
		if argWords[4] == "month" {
			if argWords[3] == "next" {
				offsetDate := now.AddDate(0, 1, 0)
				year, yearOk = offsetDate.Year(), true
				month, monthOk = offsetDate.Month(), true
			} else if argWords[3] == "previous" || argWords[3] == "prev" || argWords[3] == "last" {
				offsetDate := now.AddDate(0, -1, 0)
				year, yearOk = offsetDate.Year(), true
				month, monthOk = offsetDate.Month(), true
			}
		}
		if monthOk && yearOk && argWords[2] == "of" {
			firstDay := time.Date(year, month, 1, 0, 0, 0, 0, nil)
			if argWords[1] == "day" {
				switch argWords[0] {
				case "first": // "first day of [month] [year]"
					return firstDay, false
				case "last": // "last day of [month] [year]"
					return firstDay.AddDate(0, 1, 0).AddDate(0, 0, -1), false
				}
			} else if weekdayOk {
				switch argWords[0] {
				case "first": // "first [weekday] of [month] [year]"
					return nextWeekday(firstDay, weekday), false
				case "last": // "last [weekday] of [month] [year]"
					return previousWeekday(firstDay.AddDate(0, 1, 0).AddDate(0, 0, -1), weekday), false
				}
			}
		}
	}

	return dates.ParseStructuredDate(ctx.Value.String())
}

// SetState sets a value in the state
func (ctx *ValidatorCtx) SetState(key string, value any) {
	ctx.state.state[key] = value
}

// GetState gets a value from the state
func (ctx *ValidatorCtx) GetState(key string) (any, bool) {
	value, oke := ctx.state.state[key]
	return value, oke
}

// Bail indicates that the validator should stop after the first error for this field
func (ctx *ValidatorCtx) Bail() {
	ctx.state.bail = true
}

// UnBail indicates that the validator should continue after the first error for this field
func (ctx *ValidatorCtx) UnBail() {
	ctx.state.bail = false
}

// BailStatus returns the current bail status, if true the validator will stop after the first error
func (ctx *ValidatorCtx) BailStatus() bool {
	return ctx.state.bail
}

// field tries to return a value from the input based on the requested path
// There are 2 main ways of using this function
//
// 1. Absolute path:
//   - "foo.1.bar" = Get from the input (struct) the field "foo", then when it's a list like get the element at index 1 from the list, then get the field "bar" from the struct
//   - "" = Get the source input
//
// 2. Relative path:
//   - ".foo" = Get relative to the currently processed struct the field "foo"
//   - ".1" = Get relative to the currently processed list the element at index 1
//   - "." = Get the currently processed struct
//   - "..foo" = Get the parent of the currentl`y processed struct and then get the field "foo" from it
//
// If nil is returned the field does not exist or path is invalid
// If a needle with only a reflect.Type is returned the path exists but the value is nil
func (ctx *ValidatorCtx) Field(key string) *Needle {
	needle := ctx.state.validator.field(ctx.state.stack, key)
	if needle == nil {
		return nil
	}

	ctx.lastObtainedField = needle
	return needle
}

type SimpleStackElement struct {
	GoName   string
	JsonName string
	Index    int // Only for kind == StackKindList
	Kind     StackKind
}

// Stack returns the path to the currently processed field
// !!DO NOT MODIFY THE STACK!!, it will break the validator and cause panics
func (ctx *ValidatorCtx) Stack() Stack {
	return ctx.state.stack
}

// ObjectFieldName returns the name of the currently processed field in the object
// If the currently processed field is not an object field it will return an empty string
func (ctx *ValidatorCtx) ObjectFieldName() string {
	stack := ctx.state.stack
	if len(stack) == 0 {
		return ""
	}

	element := stack[len(stack)-1]
	if element.Kind != StackKindObject {
		return ""
	}

	return element.GoName
}
