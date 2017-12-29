package p2p

import (
	"net"
	"net/rpc"
	"time"

	"github.com/armen/dp/link"
	"github.com/armen/dp/link/node"
)

// Payload wraps the message and the ID of the sender.
type Payload struct {
	ID      string
	Addr    net.Addr
	Message link.Message
}

// Node implements perfect peer-to-peer link.
type Node struct {
	// Deliver handler
	deliver func(link.Peer, link.Message)

	// RPC client keepAlive and connections
	keepAlive time.Duration
	conn      map[string]*rpc.Client

	// RPC channel name
	channel string

	// A multiplexer to run events in a mutually exclusive way
	mux chan func()

	link.Node
}

// WithDefault can be used to instantiate a new node with default options.
var WithDefault = func(n *Node) {
	n.Node = node.New(node.WithDefault)
}

// WithChannel sets the RPC channel name when instantiating a new p2p node.
func WithChannel(channel string) func(*Node) {
	return func(n *Node) {
		n.channel = channel
	}
}

// WithNode sets the node when instantiating a new p2p node (e.g.
// New(WithNode(n), ...)).
func WithNode(node link.Node) func(*Node) {
	return func(n *Node) {
		n.Node = node
	}
}

// WithKeepAlive sets the TCP keepalive when instantiating a new p2p node
// (e.g. New(WithKeepAlive(1 * time.Seconds), ...)).
func WithKeepAlive(keepAlive time.Duration) func(*Node) {
	return func(n *Node) {
		n.keepAlive = keepAlive
	}
}

// New instantiates a new TCP based perfect peer-to-peer link.
func New(configs ...func(*Node)) *Node {
	n := &Node{
		keepAlive: 1 * time.Second,
		channel:   "p2p",
		deliver:   func(link.Peer, link.Message) {},
		conn:      make(map[string]*rpc.Client),
		mux:       make(chan func()),
	}

	for _, config := range configs {
		config(n)
	}

	n.init()
	go n.react()

	return n
}

// Deliver registers the deliver handler.
func (n *Node) Deliver(f func(p link.Peer, m link.Message)) {
	n.deliver = f
}

// react mutually executes events.
func (n *Node) react() {
	for f := range n.mux {
		f()
	}
}
