package Errors

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
		Message: fmt.Sprintf("unable to create State instance with type '{stateType.FullName}' for Saga '{correlationId}'"),
	}
}
