package link

import (
	"net"
)

// Peer represents a single node in the cluster.
type Peer interface {
	ID() string // Returns the ID of the peer

	Addr() net.Addr // Add returns the network address of the node

	Peers() []Peer  // Returns the list of peers of the node
	AddPeer(p Peer) // Adds a new peer to the peers list
}
