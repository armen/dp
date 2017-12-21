package node

import (
	"crypto/rand"
	"fmt"
	"io"
	"net"

	"github.com/armen/dp/link"
)

// Node represents a node in the cluster with its peers.
type Node struct {
	id    string
	addr  net.Addr
	peers []link.Peer
	isset map[string]bool
}

// WithID can be used when instantiating a new node (e.g. New(WithID("identifier"), ...))
func WithID(id string) func(*Node) {
	return func(n *Node) { n.id = id }
}

// WithAddr can be used when instantiating a new node (e.g. New(WithAddr(...), ...))
func WithAddr(addr net.Addr) func(*Node) {
	return func(n *Node) { n.addr = addr }
}

// WithPeer can be used when instantiating a new node (e.g. New(WithPeer(...), ...))
func WithPeer(p link.Peer) func(*Node) {
	return func(n *Node) {
		if n.isset[p.ID()] || p.ID() == n.ID() {
			return
		}
		n.isset[p.ID()] = true
		n.peers = append(n.peers, p)
	}
}

// New instantiates a new node.
func New(configs ...func(*Node)) *Node {
	n := &Node{
		isset: make(map[string]bool),
		peers: make([]link.Peer, 0),
	}

	for _, config := range configs {
		config(n)
	}

	if n.id == "" {
		uid := make([]byte, 16)
		io.ReadFull(rand.Reader, uid)
		n.id = fmt.Sprintf("%X", uid)
	}

	return n
}

// ID returns the id of the node.
func (n *Node) ID() string {
	return n.id
}

// Addr returns the network addres of the node.
func (n *Node) Addr() net.Addr {
	return n.addr
}

// Peers returns the list of peers of the node.
func (n *Node) Peers() []link.Peer {
	return n.peers
}
