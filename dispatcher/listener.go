package dispatcher

import (
	"context"
)

// Listener is an interface that must contain the `Fire()` method.
type Listener interface {
	Fire(ctx context.Context, event interface{})
}
