// Package irdp is an implementation of distributed programming abstractions in
// go
package irdp

// Reactor is the event loop which will guarantee mutually exclusiive execution
// of the events.
type Reactor interface {
	React()
}
