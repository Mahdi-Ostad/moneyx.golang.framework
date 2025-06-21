package RoutineManagement

import (
	"log"
	"reflect"
	"time"
)

type HeartBeatManager struct {
	heartbeats []*heartBeat
}

type heartBeat struct {
	funcName         string
	heartbeatChannel chan struct{}
	timer            *time.Ticker
	interval         time.Duration
}

func (h *HeartBeatManager) AddHeartbeat(funcName string, initialDelay time.Duration, interval time.Duration) chan<- struct{} {
	newHeartBeat := &heartBeat{
		funcName:         funcName,
		interval:         interval,
		heartbeatChannel: make(chan struct{}),
		timer:            time.NewTicker(initialDelay),
	}
	h.heartbeats = append(h.heartbeats, newHeartBeat)
	return newHeartBeat.heartbeatChannel
}

func (h *HeartBeatManager) StartHealthCheck() {
	for {
		count := len(h.heartbeats)
		if count == 0 {
			time.Sleep(time.Second)
			continue
		}
		cases := make([]reflect.SelectCase, count*2)
		for i := 0; i < count; i++ {
			cases[2*i] = reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(h.heartbeats[i].heartbeatChannel),
			}
			cases[2*i+1] = reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(h.heartbeats[i].timer.C),
			}
		}
		chosen, _, ok := reflect.Select(cases)
		if !ok {
			continue
		}
		heartbeat := h.heartbeats[chosen/2]
		if chosen%2 == 1 {
			log.Fatalf("Function %s heartbeat not received", heartbeat.funcName)
		}
		log.Printf("Function %s heartbeat received", heartbeat.funcName)
		heartbeat.timer = time.NewTicker(heartbeat.interval)
	}
}
