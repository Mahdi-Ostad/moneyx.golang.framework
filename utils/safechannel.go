package utils

type SafeChannel[T any] struct {
	channel   chan T
	onFail    func(T)
	onSucceed func(T)
}

// bufferSize 0 means boundless
func NewSafeChannel[T any](bufferSize int, onSucceedToAdd func(T), onFailToAdd func(T)) SafeChannel[T] {
	return SafeChannel[T]{
		channel:   make(chan T, bufferSize),
		onSucceed: onSucceedToAdd,
		onFail:    onFailToAdd,
	}
}

func (s *SafeChannel[T]) Add(data T) bool {
	select {
	case s.channel <- data:
		if s.onSucceed != nil {
			s.onSucceed(data)
		}
		return true
	default:
		if s.onFail != nil {
			s.onFail(data)
		}
		return false
	}
}

func (s *SafeChannel[T]) AddForce(data T) {
	s.channel <- data
}

func (s *SafeChannel[T]) GetForce() T {
	return <-s.channel
}

func (s *SafeChannel[T]) GetOrDefault(defaultValue T) T {
	var value T
	select {
	case value = <-s.channel:
		return value
	default:
		return defaultValue
	}
}
