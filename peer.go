package dp

import (
	"net"
)

// Peer represents a single node in the cluster.
type Peer struct {
	Addr net.Addr // Addr is the peer address
}
