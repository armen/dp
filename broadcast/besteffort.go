// Package broadcast implements broadcast abstraction.
package broadcast

import (
	"github.com/armen/dp"
	"github.com/armen/dp/link"
)

// BestEffort defines the interface and properties of best-effort broadcast.
//
// Properties:
// 	BEB1: Validity:
// 		- If a correct process broadcasts a message m, then every
// 		  correct process eventually delivers m.
// 	BEB2: No duplication:
// 		- No message is delivered more than once.
// 	BEB3: No creation:
// 		- If a process delivers a message m with sender s, then m was
// 		  previously broadcast by process s.
//
type BestEffort interface {
	Broadcast(m link.Message)                  // Broadcasts a message m to all processes
	Deliver(func(p link.Peer, m link.Message)) // Delivers a message m broadcast by process p

	dp.Reactor
}
