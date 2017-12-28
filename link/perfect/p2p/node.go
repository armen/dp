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

// Node implements perfect peer-to-peer link.
type Node struct {
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
func New(node link.Node, l net.Listener, keepAlive time.Duration) *Node {
	n := &Node{
		deliver:   func(link.Peer, link.Message) {},
		keepAlive: keepAlive,
		conn:      make(map[string]*rpc.Client),
		listener:  l,
		mux:       make(chan func()),
		Node:      node,
	}

	go n.react()

	return n
}

// Deliver registers the deliver handler.
func (n *Node) Deliver(f func(p link.Peer, m link.Message)) {
	n.deliver = f
}

// react mutually executes events.
func (n *Node) react() {
	n.init()

	for f := range n.mux {
		f()
	}
}
