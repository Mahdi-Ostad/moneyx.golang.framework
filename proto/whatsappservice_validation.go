// versions:
// 	gofr-cli v0.6.0
// 	gofr.dev v1.37.0
// 	source: wa.proto

package proto

import Validator "moneyx.golang.framework/validator"

// WhatsappServiceGoFrValidation defines the input validator implementation.
// Customize the struct with required dependencies and fields as needed.

type WhatsappServiceGoFrValidation struct {
}

func (s *WhatsappServiceGoFrValidation) CheckUserState(validator *Validator.EntryValidator[*EmptyWrapper]) {

}
func (s *WhatsappServiceGoFrValidation) Initialize(validator *Validator.EntryValidator[*EmptyWrapper]) {

}
func (s *WhatsappServiceGoFrValidation) SendMessage(validator *Validator.EntryValidator[*SendMessageCommandProtoWrapper]) {

}
func (s *WhatsappServiceGoFrValidation) SayHello(validator *Validator.EntryValidator[*EmptyWrapper]) {

}
func (s *WhatsappServiceGoFrValidation) SignOut(validator *Validator.EntryValidator[*StringIdArgWrapper]) {

}
