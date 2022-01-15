package dispatcher

import (
	"context"
	"fmt"
)

// The Dispatcher structure which contains a list of registered event names with their listeners.
type Dispatcher struct {
	eventsList map[Name]Listener
}

// NewDispatcher creates a new event Dispatcher.
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		eventsList: make(map[Name]Listener),
	}
}

// Register registers one or multiple new events and the listener that belongs to it or them.
// If the event is already registered, it returns an error.
func (d *Dispatcher) Register(listener Listener, names ...Name) error {
	for _, name := range names {
		_, exists := d.eventsList[name]
		if exists {
			return fmt.Errorf("the event '%s' is already registered", name)
		}
		d.eventsList[name] = listener
	}
	return nil
}

// Dispatch is used to dispatch an event. It fires the listener that contains the event.
// If the event being dispatched is not registered, it returns an error.
func (d *Dispatcher) Dispatch(ctx context.Context, name Name, event interface{}) error {
	_, exists := d.eventsList[name]
	if !exists {
		return fmt.Errorf("the event '%s' is not registered", name)
	}
	d.eventsList[name].Fire(ctx, event)
	return nil
}
