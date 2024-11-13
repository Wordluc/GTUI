package EventManager

import (
	"errors"

	"github.com/Wordluc/GTUI/Core"
)

type EventType int8

const (
	//Event to call when a component need an instant screee refresh
	Refresh EventType = iota
)

type EventManager struct {
	subscribers map[EventType][]func(Core.IComponent)
}

var eventManager *EventManager=nil

func Setup()error {
	if eventManager!=nil {
		return errors.New("EventManager already setup")
	}
	eventManager = &EventManager{subscribers: make(map[EventType][]func(Core.IComponent))}
	return nil
}

func Call(typeEvent EventType,caller Core.IComponent)error {
	if eventManager==nil {
		return errors.New("EventManager not setup")
	}
	for _, f := range eventManager.subscribers[typeEvent] {
		f(caller)
	}
	return nil
}

func Subscribe(typeEvent EventType, f func(Core.IComponent)) {
	if eventManager==nil {
		return
	}
	eventManager.subscribers[typeEvent] = append(eventManager.subscribers[typeEvent], f)
}
