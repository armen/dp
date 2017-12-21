package beb

import (
	"github.com/armen/dp/link"
)

// Node implements best-effort broadcast abstraction.
type Node struct {
	pl         link.Perfect
	bebDeliver func(link.Peer, link.Message)

	// A multiplexer to run events in a mutually exclusive way
	mux chan func()

	link.Node
}

// New instantiates
func New(pl link.Perfect) *Node {
	n := &Node{
		pl:         pl,
		bebDeliver: func(link.Peer, link.Message) {},
		mux:        make(chan func()),
		Node:       pl.(link.Node),
	}

	// Register our plDeliver handler
	pl.Deliver(n.plDeliver)

	return n
}

// Deliver registers a bebDeliver event handler
func (n *Node) Deliver(f func(link.Peer, link.Message)) {
	n.bebDeliver = f
}

// React mutually executes events.
func (n *Node) React() {
	go n.pl.React()

	// Mutually exclusive execution of closures
	for f := range n.mux {
		f()
	}
}
