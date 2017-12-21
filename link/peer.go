package link

import (
	"net"
)

// Peer represents a peer in the cluster.
type Peer interface {
	ID() string     // Returns the ID of the peer
	Addr() net.Addr // Add returns the network address of the node
}
