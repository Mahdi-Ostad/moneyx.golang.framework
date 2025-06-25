package RoutineManagement

import "sync"

type fanElementModel[T any] struct {
	value      T
	shouldSkip bool
}

type FanModel[T any] struct {
	channel chan fanElementModel[T]
	action  func(T) error
	onFail  func(error)
}

func NewFanModel[T any](action func(T) error, onFail func(error)) *FanModel[T] {
	return &FanModel[T]{
		action: action,
		onFail: onFail,
	}
}

func FanInFanOut[T any](input []T, models []FanModel[T]) {
	count := len(input)
	if count == 0 {
		return
	}
	for i := 0; i < count; i++ {
		models[i].channel = make(chan fanElementModel[T], count)
	}
	var wg sync.WaitGroup
	for index, model := range models {
		wg.Add(1)
		var nextChan chan<- fanElementModel[T]
		if index+1 < count {
			nextChan = models[index+1].channel
		} else {
			nextChan = nil
		}
		go model.singleFanRun(&wg, count, nextChan)
	}
	for _, in := range input {
		models[0].channel <- fanElementModel[T]{
			value:      in,
			shouldSkip: false,
		}
	}
	wg.Wait()
}

func (f *FanModel[T]) singleFanRun(wg *sync.WaitGroup, count int, nextChan chan<- fanElementModel[T]) {
	defer wg.Done()
	defer close(f.channel)
	cnt := 0
	for data := range f.channel {
		cnt++
		shouldSkip := false
		if !data.shouldSkip {
			err := f.action(data.value)
			if err != nil {
				f.onFail(err)
				shouldSkip = true
			}
		} else {
			shouldSkip = true
		}
		if nextChan != nil {
			nextChan <- fanElementModel[T]{
				value:      data.value,
				shouldSkip: shouldSkip,
			}
		}
		if cnt == count {
			break
		}
	}
}
