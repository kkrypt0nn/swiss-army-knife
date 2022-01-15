package dispatcher

import (
	"context"
)

// Name is the unique name of the event.
type Name string

// Event is an interface that must contain the `Handle()` method.
// Each event can have its own structure, and therefore you can give separate attributes to each event.
type Event interface {
	Handle(ctx context.Context)
}
