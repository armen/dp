package node

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/armen/dp/link"
)

// Node represents a node in the cluster with its peers.
type Node struct {
	peers    []link.Peer
	isset    map[string]bool
	listener net.Listener

	*Peer
}

// WithDefault can be used to instantiate a new node with default options.
var WithDefault = func(n *Node) {
	l, e := net.Listen("tcp", "127.0.0.1:0") // any available address
	if e != nil {
		log.Fatalf("net.Listen tcp :0: %v", e)
	}
	n.listener = l
	n.addr = l.Addr()
}

// WithListener sets the listener when instantiate a new node.
func WithListener(l net.Listener) func(*Node) {
	return func(n *Node) {
		n.listener = l
		n.addr = l.Addr()
	}
}

// WithID sets the ID when instantiating a new node (e.g. New(WithID("identifier"), ...)).
func WithID(id string) func(*Node) {
	return func(n *Node) { n.id = id }
}

// WithAddr sets the address when instantiating a new node (e.g. New(WithAddr(...), ...)).
func WithAddr(addr net.Addr) func(*Node) {
	return func(n *Node) { n.addr = addr }
}

// WithPeer adds a peer when instantiating a new node (e.g. New(WithPeer(...), ...)).
func WithPeer(p link.Peer) func(*Node) {
	return func(n *Node) { n.AddPeer(p) }
}

// New instantiates a new node.
func New(configs ...func(*Node)) *Node {
	n := &Node{
		isset: make(map[string]bool),
		peers: make([]link.Peer, 0),
		Peer:  &Peer{},
	}

	for _, config := range configs {
		config(n)
	}

	if n.ID() == "" {
		uid := make([]byte, 16)
		io.ReadFull(rand.Reader, uid)
		n.id = fmt.Sprintf("%X", uid)
	}

	return n
}

// Peers returns the list of peers of the node.
func (n *Node) Peers() []link.Peer {
	return n.peers
}

// Members returns all the members including the current node.
func (n *Node) Members() []link.Peer {
	return append(n.peers, n)
}

// Listener returns the network listener.
func (n *Node) Listener() net.Listener {
	return n.listener
}

// AddPeer adds a new peer to the peers list.
func (n *Node) AddPeer(p link.Peer) {
	if n.isset[p.ID()] || p.ID() == n.ID() {
		return
	}
	n.isset[p.ID()] = true
	n.peers = append(n.peers, p)
}

// N returns total number of nodes in the cluster.
func (n *Node) N() int {
	return len(n.peers) + 1
}
