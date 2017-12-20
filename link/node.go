package link

import (
	"crypto/rand"
	"fmt"
	"io"
	"net"
	"sync"
)

// Node represents a node in the cluster with it's peers.
type Node struct {
	id    string
	addr  net.Addr
	isset map[string]bool
	peers []Peer
	mux   sync.Mutex
}

// NewNode instantiates a new node.
func NewNode(addr net.Addr) *Node {
	uid := make([]byte, 16)
	io.ReadFull(rand.Reader, uid)
	id := fmt.Sprintf("%X", uid)

	return &Node{
		id:    id,
		addr:  addr,
		isset: make(map[string]bool),
		peers: make([]Peer, 0),
	}
}

// NewNodeWithID a new node with ID.
func NewNodeWithID(id string) *Node {
	return &Node{
		id:    id,
		addr:  &net.TCPAddr{},
		isset: make(map[string]bool),
		peers: make([]Peer, 0),
	}
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
func (n *Node) Peers() []Peer {
	return n.peers
}

// AddPeer adds a new peer to the peers list.
func (n *Node) AddPeer(p Peer) {
	n.mux.Lock()
	defer n.mux.Unlock()

	if n.isset[p.ID()] || p.ID() == n.ID() {
		return
	}
	n.isset[p.ID()] = true
	n.peers = append(n.peers, p)
}
