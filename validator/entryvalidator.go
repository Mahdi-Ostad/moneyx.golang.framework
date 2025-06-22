package Validator

import (
	"fmt"
	"reflect"

	"github.com/shivajichalise/validator"
	_ "github.com/shivajichalise/validator/rules"
	MoneyxErrors "moneyx.golang.framework/moneyxerrors"
)

type EntryValidator[T any] struct {
	data        T
	rules       map[string][]string
	customRules map[string]func(T) error
}

func NewEntryValidator[T any](input T) *EntryValidator[T] {
	return &EntryValidator[T]{
		data:        input,
		rules:       map[string][]string{},
		customRules: make(map[string]func(T) error),
	}
}

// Value must be a non-empty string
func (ev *EntryValidator[T]) NotEmpty(field string) *EntryValidator[T] {
	if _, ok := ev.rules[field]; !ok {
		ev.rules[field] = make([]string, 0)
	}
	ev.rules[field] = append(ev.rules[field], "string")
	return ev
}

// String length must be ≥ n
func (ev *EntryValidator[T]) MinLength(field string, n int) *EntryValidator[T] {
	if _, ok := ev.rules[field]; !ok {
		ev.rules[field] = make([]string, 0)
	}
	ev.rules[field] = append(ev.rules[field], fmt.Sprintf("min:%d", n))
	return ev
}

// String length must be ≤ n
func (ev *EntryValidator[T]) MaxLength(field string, n int) *EntryValidator[T] {
	if _, ok := ev.rules[field]; !ok {
		ev.rules[field] = make([]string, 0)
	}
	ev.rules[field] = append(ev.rules[field], fmt.Sprintf("max:%d", n))
	return ev
}

// Accepts int and float
func (ev *EntryValidator[T]) MustNumeric(field string) *EntryValidator[T] {
	if _, ok := ev.rules[field]; !ok {
		ev.rules[field] = make([]string, 0)
	}
	ev.rules[field] = append(ev.rules[field], "numeric")
	return ev
}

// Value must be an integer
func (ev *EntryValidator[T]) MustInt(field string) *EntryValidator[T] {
	if _, ok := ev.rules[field]; !ok {
		ev.rules[field] = make([]string, 0)
	}
	ev.rules[field] = append(ev.rules[field], "int")
	return ev
}

// Value must be a float64
func (ev *EntryValidator[T]) MustFloat(field string) *EntryValidator[T] {
	if _, ok := ev.rules[field]; !ok {
		ev.rules[field] = make([]string, 0)
	}
	ev.rules[field] = append(ev.rules[field], "float64")
	return ev
}

// Value must be greater than n
func (ev *EntryValidator[T]) GtFloat(field string, n float64) *EntryValidator[T] {
	if _, ok := ev.rules[field]; !ok {
		ev.rules[field] = make([]string, 0)
	}
	ev.rules[field] = append(ev.rules[field], fmt.Sprintf("gt:%.10f", n))
	return ev
}

// Value must be greater than n
func (ev *EntryValidator[T]) GtInt(field string, n int) *EntryValidator[T] {
	if _, ok := ev.rules[field]; !ok {
		ev.rules[field] = make([]string, 0)
	}
	ev.rules[field] = append(ev.rules[field], fmt.Sprintf("gt:%d", n))
	return ev
}

// Value must be less than n
func (ev *EntryValidator[T]) LtFloat(field string, n float64) *EntryValidator[T] {
	if _, ok := ev.rules[field]; !ok {
		ev.rules[field] = make([]string, 0)
	}
	ev.rules[field] = append(ev.rules[field], fmt.Sprintf("lt:%.10f", n))
	return ev
}

// Value must be less than n
func (ev *EntryValidator[T]) LtInt(field string, n int) *EntryValidator[T] {
	if _, ok := ev.rules[field]; !ok {
		ev.rules[field] = make([]string, 0)
	}
	ev.rules[field] = append(ev.rules[field], fmt.Sprintf("lt:%d", n))
	return ev
}

// Value must be a boolean
func (ev *EntryValidator[T]) MustBool(field string) *EntryValidator[T] {
	if _, ok := ev.rules[field]; !ok {
		ev.rules[field] = make([]string, 0)
	}
	ev.rules[field] = append(ev.rules[field], "boolean")
	return ev
}

// Value must be between min and max
func (ev *EntryValidator[T]) BetweenInt(field string, min, max int) *EntryValidator[T] {
	if _, ok := ev.rules[field]; !ok {
		ev.rules[field] = make([]string, 0)
	}
	ev.rules[field] = append(ev.rules[field], fmt.Sprintf("between:%d,%d", min, max))
	return ev
}

// Value must be between min and max
func (ev *EntryValidator[T]) BetweenFloat(field string, min, max float64) *EntryValidator[T] {
	if _, ok := ev.rules[field]; !ok {
		ev.rules[field] = make([]string, 0)
	}
	ev.rules[field] = append(ev.rules[field], fmt.Sprintf("between:%.10f,%.10f", min, max))
	return ev
}

func (ev *EntryValidator[T]) AddCustomRule(name string, fn func(T) error) error {
	if _, ok := ev.customRules[name]; ok {
		return fmt.Errorf("custom rule name %s duplicate", name)
	}
	ev.customRules[name] = fn
	return nil
}

func (ev *EntryValidator[T]) Validate() *MoneyxErrors.ValidatorException[T] {
	errors := make([]error, 0)
	for name, fn := range ev.customRules {
		err := fn(ev.data)
		if err != nil {
			errors = append(errors, fmt.Errorf("custom rule %s validation: %w", name, err))
		}
	}
	// data := make(map[string]any)
	// val := reflect.ValueOf(ev.data)
	// typ := reflect.TypeOf(ev.data)

	// if val.Kind() == reflect.Ptr {
	// 	val = val.Elem()
	// 	typ = typ.Elem()
	// }

	// if val.Kind() != reflect.Struct {
	// 	return errors
	// }

	// for i := 0; i < val.NumField(); i++ {
	// 	field := typ.Field(i)
	// 	if field.PkgPath != "" {
	// 		// skip unexported fields
	// 		continue
	// 	}
	// 	data[field.Name] = val.Field(i).Interface()
	// }

	data := structToMap(ev.data, "")

	v := validator.Make(data, ev.rules)

	if !v.Validate() {
		for field, errs := range v.Errors() {
			for _, msg := range errs {
				errors = append(errors, fmt.Errorf("%s: %s", field, msg))
			}
		}
	}

	return MoneyxErrors.NewValidatorException(ev.data, errors)
}

func structToMap(v interface{}, parentPrefix string) map[string]any {
	result := make(map[string]any)

	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() != reflect.Struct {
		return result
	}

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		if field.PkgPath != "" {
			continue // unexported field
		}

		fieldVal := val.Field(i)

		switch fieldVal.Kind() {
		case reflect.Struct:
			result[field.Name] = structToMap(fieldVal.Interface(), fmt.Sprintf("%s%s.", parentPrefix, field.Name))
		case reflect.Ptr:
			if fieldVal.IsNil() {
				result[field.Name] = nil
			} else if fieldVal.Elem().Kind() == reflect.Struct {
				result[field.Name] = structToMap(fieldVal.Interface(), fmt.Sprintf("%s%s.", parentPrefix, field.Name))
			} else {
				result[field.Name] = fieldVal.Interface()
			}
		case reflect.Slice:
			// Skip slice fields entirely
			continue
		default:
			result[field.Name] = fieldVal.Interface()
		}
	}

	return result
}
