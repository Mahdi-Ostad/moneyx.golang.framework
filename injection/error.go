package injection

import (
	"fmt"
	"reflect"
)

type InjectionError struct {
	Message      string
	InjectedType reflect.Type
}

func (e *InjectionError) Error() string {
	return fmt.Sprintf("Injection Error: %s could not be injected. Type: %v", e.Message, e.InjectedType)
}
