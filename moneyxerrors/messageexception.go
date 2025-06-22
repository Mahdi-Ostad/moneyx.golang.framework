package MoneyxErrors

type MessageException struct {
	Message string
}

func (s *MessageException) Error() string {
	return s.Message
}

func (s *MessageException) LogMessage() string {
	return s.Message
}
