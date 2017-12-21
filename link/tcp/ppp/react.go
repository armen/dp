package ppp

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

// Ppp implements perfect peer-to-peer link.
type Ppp struct {
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
func New(node link.Node, l net.Listener, keepAlive time.Duration) *Ppp {
	return &Ppp{
		keepAlive: keepAlive,
		conn:      make(map[string]*rpc.Client),
		listener:  l,
		mux:       make(chan func()),
		Node:      node,
	}
}

// Deliver registers the deliver handler.
func (p *Ppp) Deliver(f func(p link.Peer, m link.Message)) {
	p.deliver = f
}

// React mutually executes events.
func (p *Ppp) React() {
	p.init()

	for f := range p.mux {
		f()
	}
}
