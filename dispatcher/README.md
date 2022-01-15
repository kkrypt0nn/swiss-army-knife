# Dispatcher

This package provides an easy to use event dispatcher. It contains an interface for events.

## Example

At first we need to create a new event (`event_sleep.go`):

```go
// Sleep is the name of the event, it is used to recognize the event and should be unique.
const Sleep event.Name = "event.sleep"

// SleepEvent is being created.
type SleepEvent struct {
	Duration time.Duration
}

// Handle is the function that is called when the event is being fired from the listener.
func (e SleepEvent) Handle(ctx context.Context) {
	time.Sleep(e.Duration)
	log.Printf("[%s] Slept %f seconds since dispatch time.", Sleep, e.Duration.Seconds())
}
```

Then we need to setup a listener (`listener.go`) so that we only fire the events that we want:

```go
type MyListener struct{}

// Fire is executed when an event is getting fired by the dispatcher.
func (u MyListener) Fire(ctx context.Context, eventToRegister interface{}) {
	switch eventToRegister := eventToRegister.(type) {
	case SleepEvent:
		eventToRegister.Handle(ctx)
	default:
		log.Println("Tried to fire an invalid event:", eventToRegister)
	}
}
```

Then we can start creating the main function with the dispatcher being created, an event being registered and later dispatched (`main.go`):

```go
func main() {
	// Create a new event dispatcher.
	dispatcher := NewDispatcher()

	// Register new event with its listener to the dispatcher.
	if err := dispatcher.Register(MyListener{}, Sleep); err != nil {
		log.Fatalln(err)
	}

	// Dispatch the event 'event.sleep'.
	log.Println("Dispatching event...")
	err := dispatcher.Dispatch(context.Background(), Sleep, SleepEvent{
		Duration: 3 * time.Second,
	})
	if err != nil {
		log.Println(err)
	}
}

```
