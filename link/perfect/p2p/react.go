package p2p

import (
	"net"
	"net/rpc"
	"time"

	"github.com/armen/dp/link"
)

// Payload wraps the message and the ID of the sender.
type Payload struct {
	ID      string
	Message link.Message
}

// P2p implements perfect peer-to-peer link.
type P2p struct {
	// Deliver handler
	deliver func(link.Peer, link.Message)

	// RPC client keepAlive and connections
	keepAlive time.Duration
	conn      map[string]*rpc.Client

	// RPC server's listener
	listener net.Listener

	// A multiplexer to run events in a mutually exclusive way
	mux chan func()

	link.Node
}

// New instantiates a new TCP based perfect peer-to-peer link.
func New(node link.Node, l net.Listener, keepAlive time.Duration) *P2p {
	return &P2p{
		keepAlive: keepAlive,
		conn:      make(map[string]*rpc.Client),
		listener:  l,
		mux:       make(chan func()),
		Node:      node,
	}
}

// Deliver registers the deliver handler.
func (p *P2p) Deliver(f func(p link.Peer, m link.Message)) {
	p.deliver = f
}

// React mutually executes events.
func (p *P2p) React() {
	p.init()

	for f := range p.mux {
		f()
	}
}
