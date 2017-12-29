package beb

import (
	"github.com/armen/dp/link"
)

// Node implements best-effort broadcast abstraction.
type Node struct {
	// Perfect Point-To-Point link
	pl link.Perfect

	// Deliver handler
	bebDeliver func(link.Peer, link.Message)

	// A multiplexer to run events in a mutually exclusive way
	mux chan func()

	link.Node
}

// New instantiates a basic-broadcast node.
func New(pl link.Perfect) *Node {
	n := &Node{
		pl:         pl,
		bebDeliver: func(link.Peer, link.Message) {},
		mux:        make(chan func()),
		Node:       pl.(link.Node),
	}

	n.init()
	go n.react()

	return n
}

// Deliver registers a bebDeliver event handler
func (n *Node) Deliver(f func(link.Peer, link.Message)) {
	n.bebDeliver = f
}

// react mutually executes events.
func (n *Node) react() {
	for f := range n.mux {
		f()
	}
}
