package MoneyxErrors

import (
	"fmt"
)

type ValidatorException[T any] struct {
	input  T
	errors []error
}

func NewValidatorException[T any](input T, errors []error) *ValidatorException[T] {
	if len(errors) == 0 {
		return nil
	}
	return &ValidatorException[T]{
		input:  input,
		errors: errors,
	}
}

func (v *ValidatorException[T]) Error() string {
	return fmt.Sprintf("\"request\": %v ... \"errors\": %v", v.input, v.errors)
}

func (v *ValidatorException[T]) LogMessage() string {
	return fmt.Sprintf("Entry data validation failded, and command info is %v: %v", v.input, v.errors)
}
