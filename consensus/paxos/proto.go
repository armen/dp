// Package paxos implements different variants of paxos algorithm.
package paxos

import (
	"encoding/gob"
)

func init() {
	gob.Register(Prepare{})
	gob.Register(Promise{})
	gob.Register(Accept{})
	gob.Register(Accepted{})
	gob.Register(Decided{})
	gob.Register(Nack{})
	gob.Register(Ballot{})
}

// Prepare request
type Prepare struct {
	Ballot *Ballot
}

// Promise response to a Prepare request
type Promise struct {
	Ballot   *Ballot
	Accepted *Ballot
	Val      interface{} // The associated value
}

// Accept request
type Accept struct {
	Ballot *Ballot
	Val    interface{} // The associated value
}

// Accepted response to an Accept request
type Accepted struct {
	Ballot *Ballot
	Val    interface{} // The associated value
}

// Decided message to broadcast the decided value
type Decided struct {
	Ballot *Ballot
	Val    interface{} // The associated value
}

// Nack is a NotOK acknowledgement
type Nack struct {
	Ballot *Ballot
}
