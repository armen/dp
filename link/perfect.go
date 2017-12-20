// Package link implements communication link abstraction.
package link

import (
	"github.com/armen/dp"
)

// Perfect defines the interface and properties of perfect point-to-point links.
//
// Properties:
// 	PL1: Reliable delivery:
// 		- If a correct process p sends a message m to a correct process q, then
// 		  q eventually delivers m.
// 	PL2: No duplication:
// 		- No message is delivered by a process more than once.
// 	PL3: No creation:
// 		- If some process q delivers a message m with sender p, then m was
// 		  previously sent to q by process p.
//
type Perfect interface {
	Send(q Peer, m Message) error    // Requests to send message m to process q
	Deliver(func(p Peer, m Message)) // Delivers message m sent by process p

	dp.Reactor
}
