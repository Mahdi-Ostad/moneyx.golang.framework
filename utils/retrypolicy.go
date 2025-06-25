package utils

import "time"

type RetryPolicy struct {
	onRetry      func(int, error)
	waitDuration time.Duration
}

func NewRetryPolicy(waitDuration time.Duration, onRetry func(int, error)) *RetryPolicy {
	return &RetryPolicy{
		onRetry:      onRetry,
		waitDuration: waitDuration,
	}
}

func (r *RetryPolicy) RetryUntil(maxCount int, fn func() error) (err error) {
	for i := 1; i <= maxCount; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		r.onRetry(i, err)
		time.Sleep(r.waitDuration)
	}
	return err
}

func (r *RetryPolicy) RetryForever(fn func() error) {
	for i := 1; ; i++ {
		err := fn()
		if err == nil {
			return
		}
		r.onRetry(i, err)
		time.Sleep(r.waitDuration)
	}
}
