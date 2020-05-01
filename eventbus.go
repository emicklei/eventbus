package eventbus

import "sync"

// EventData represents the context of an event, a simple map.
type EventData map[string]interface{}

// Event is send on the bus to all subscribed listeners
type Event struct {
	Name string
	Data EventData
}

// EventListener is the signature of functions that can handle an Event.
type EventListener func(Event)

// The EventBus allows publish-subscribe-style communication between components
// without requiring the components to explicitly register with one another (and thus be aware of each other)
// Inspired by Guava EventBus ; this is a more lightweight implementation.
type EventBus struct {
	mutex     *sync.RWMutex
	listeners map[string][]EventListener
}

// NewEventBus return a new EventBus
func NewEventBus() *EventBus {
	return &EventBus{new(sync.RWMutex), map[string][]EventListener{}}
}

// Subscribe adds an EventListener to be called when an event is posted.
func (e *EventBus) Subscribe(name string, listener EventListener) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	list, ok := e.listeners[name]
	if !ok {
		list = []EventListener{}
	}
	list = append(list, listener)
	e.listeners[name] = list
}

// Post sends an event to all subscribed listeners.
// Parameter data is optional ; Post can only have one map parameter.
func (e *EventBus) Post(name string, data ...map[string]interface{}) {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	list, ok := e.listeners[name]
	if !ok {
		return
	}
	event := Event{Name: name}
	if len(data) == 1 {
		event.Data = EventData(data[0])
	}
	for _, each := range list[:] { // iterate over unmodifyable copy
		each(event)
	}
}
