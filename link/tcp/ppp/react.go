package ppp

import (
	"net"
	"net/rpc"
	"time"

	"github.com/armen/dp/link"
)

// Payload wraps the message and the ID of the sender.
type Payload struct {
	Src     string
	Message link.Message
}

// Ppp implements perfect peer-to-peer link.
type Ppp struct {
	deliver func(p link.Peer, m link.Message) // Deliver handler

	// RPC client keepAlive and connections
	keepAlive time.Duration
	conn      map[string]*rpc.Client

	// RPC server's listener
	listener net.Listener

	mux chan func()

	*link.Node
}

// New instantiates a new TCP based perfect peer-to-peer link.
func New(n *link.Node, l net.Listener, keepAlive time.Duration) *Ppp {
	return &Ppp{
		nil,
		keepAlive,
		make(map[string]*rpc.Client),
		l,
		make(chan func()),
		n,
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
