package MoneyxErrors

type BaseException interface {
	Error() string
	LogMessage() string
}
