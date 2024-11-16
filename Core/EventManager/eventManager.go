package EventManager

import (
	"errors"
	"sync"
	"time"
)

type EventType int8

const (
	//Event to call when a component need an instant screee refresh
	Refresh EventType = iota
	//Event to call when a element change its position
	ReorganizeElements
)

type eventPartioning struct {
	caller     []any
	subscriber *subscriber
	offsetTime int
}
type subscriber struct {
	event func([]any)
	offsetTime int
}
type EventManager struct {
	subscribers     map[EventType]*subscriber
	eventpartitions map[EventType]*eventPartioning
	mu              sync.Mutex
}

var eventManager *EventManager = nil

func Setup() error {
	if eventManager != nil {
		return errors.New("EventManager already setup")
	}
	eventManager = &EventManager{subscribers: make(map[EventType]*subscriber), eventpartitions: make(map[EventType]*eventPartioning), mu: sync.Mutex{}}
	return nil
}

func Call(typeEvent EventType, caller []any) error {
	if eventManager == nil {
		return errors.New("EventManager not setup")
	}
	eventManager.mu.Lock()
	defer eventManager.mu.Unlock()
	partition := eventManager.eventpartitions[typeEvent]
	if eventManager.subscribers[typeEvent] == nil {
		return errors.New("Event not subscribed")
	}
	if partition == nil {
		eventManager.eventpartitions[typeEvent] = &eventPartioning{caller: caller, subscriber: eventManager.subscribers[typeEvent]}
		time.AfterFunc(time.Millisecond*100, func() {
			eventManager.mu.Lock()
			defer eventManager.mu.Unlock()
			partition := eventManager.eventpartitions[typeEvent]
			if partition == nil {
				return
			}
			partition.subscriber.event(partition.caller)
			eventManager.eventpartitions[typeEvent] = nil
		})
		return nil
	}
	partition.caller = append(partition.caller, caller...)
	return nil
}

func Subscribe(typeEvent EventType,offsetTime int, f func([]any)) error {
	if eventManager == nil {
		return errors.New("EventManager not setup")
	}
	if _, ok := eventManager.subscribers[typeEvent]; ok {
		return errors.New("Event already subscribed")
	}
	eventManager.subscribers[typeEvent] = &subscriber{event: f, offsetTime: offsetTime}
	return nil
}
