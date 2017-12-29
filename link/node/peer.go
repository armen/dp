package node

import (
	"net"
)

// Peer represents a peer in the cluster.
type Peer struct {
	id   string
	addr net.Addr
}

// NewPeer instantiates a peer.
func NewPeer(id string, addr net.Addr) *Peer {
	return &Peer{
		id:   id,
		addr: addr,
	}
}

// ID returns the id of the node.
func (p *Peer) ID() string {
	return p.id
}

// Addr returns the network addres of the node.
func (p *Peer) Addr() net.Addr {
	return p.addr
}
