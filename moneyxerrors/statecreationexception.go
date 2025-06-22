package MoneyxErrors

import (
	"fmt"
	"reflect"

	"github.com/google/uuid"
)

type StateCreationException struct {
	Message string
}

func NewStateCreationException(tpe reflect.Type, correlationId uuid.UUID) *StateCreationException {
	return &StateCreationException{
		Message: fmt.Sprintf("unable to create State instance with type '%v' for Saga '%v'", tpe, correlationId),
	}
}

func (s *StateCreationException) Error() string {
	return s.Message;
}

func (s *StateCreationException) LogMessage() string {
	return s.Message;
}