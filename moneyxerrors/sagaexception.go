package MoneyxErrors

import (
	"encoding/json"
	"fmt"

	ValueObjects "moneyx.golang.framework/valueobjects"
)

type SagaException struct {
	MessageTemplate     string                             `json:"messageTemplate"`
	DescriptionMetadata []ValueObjects.DescriptionMetadata `json:"descriptionMetadata"`
	InnerException      string                             `json:"innerException"`
}

func (d *SagaException) Error() string {
	return fmt.Sprintf("Error occured with tempalte: %s, and metedata which includes %v", d.MessageTemplate, d.DescriptionMetadata)
}

func (d *SagaException) LogMessage() string {
	b, _ := json.Marshal(d)
	return string(b)
}
