package MoneyxErrors

type LockException struct {
	Message string
}

func (s *LockException) Error() string {
	return s.Message
}

func (s *LockException) LogMessage() string {
	return s.Message
}
