package RoutineManagement

import (
	"reflect"
)

type TeeModel[T any] struct {
}

func (t *TeeModel[T]) Tee(done <-chan struct{}, in <-chan T, count int, size int) []<-chan T {
	channels := make([]chan T, count)
	outputs := make([]<-chan T, count)
	for i := 0; i < count; i++ {
		channels[i] = make(chan T, size)
		outputs[i] = channels[i]
	}
	go func() {
		for _, channel := range channels {
			defer close(channel)
		}
		for val := range t.orDone(done, in) {
			cases := make([]reflect.SelectCase, count)
			for i := 0; i < count; i++ {
				cases[i] = reflect.SelectCase{
					Dir:  reflect.SelectSend,
					Chan: reflect.ValueOf(channels[i]),
					Send: reflect.ValueOf(val),
				}
			}

			for i := 0; i < count; i++ {
				chosen, _, _ := reflect.Select(cases)
				// Disable the selected channel
				cases[chosen].Chan = reflect.ValueOf(nil)
			}
		}
	}()
	return outputs
}

func (t *TeeModel[T]) orDone(done <-chan struct{}, c <-chan T) <-chan T {
	valStream := make(chan T)
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

func (t *TeeModel[T]) SafeDo(done <-chan struct{}, ch <-chan T, fn func(T)) {
	for u := range t.orDone(done, ch) {
		fn(u)
	}
}
