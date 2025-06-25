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

// adds to channel. If it's full, onFail will run, otherwise onSuccess will run
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

// blocks until data is added to the channel
func (s *SafeChannel[T]) AddForce(data T) {
	s.channel <- data
}

// blocks until it can read from channel
func (s *SafeChannel[T]) GetForce() T {
	return <-s.channel
}

// reads from channel, returns default value if channel is empty
func (s *SafeChannel[T]) GetOrDefault(defaultValue T) T {
	var value T
	select {
	case value = <-s.channel:
		return value
	default:
		return defaultValue
	}
}
