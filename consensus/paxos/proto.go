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
}

// Accept request
type Accept struct {
	Ballot *Ballot
}

// Accepted response to an Accept request
type Accepted struct {
	Ballot *Ballot
}

// Decided message to broadcast the decided value
type Decided struct {
	Ballot *Ballot
}

// Nack is a NotOK acknowledgement
type Nack struct {
	Ballot *Ballot
}
