package helpers

import "github.com/stianeikeland/go-rpio"

// ContextWithValue store Name and Pin of each device within a context.
type ContextWithValue struct {
	Name string
	Pin  rpio.Pin
}
