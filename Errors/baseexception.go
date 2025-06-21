package Errors

type BaseException interface {
	Error() string
	LogMessage() string
}
