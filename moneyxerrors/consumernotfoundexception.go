package MoneyxErrors

import (
	"fmt"
	"reflect"
)

type ConsumerNotFoundException struct {
	MessageType reflect.Type
}

func (c *ConsumerNotFoundException) Error() string {
	return fmt.Sprintf("no consumers found for message messageType '%v'", c.MessageType)
}

func (c *ConsumerNotFoundException) LogMessage() string {
	return fmt.Sprintf("ConsumerNotFoundException: %v", c.MessageType)
}
