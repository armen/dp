package dp

import (
	"net"
	"strconv"
)

// Peer represents a single node in the cluster.
type Peer struct {
	ID        string
	Addr      net.IP
	Port      uint16
	Incorrect bool
}

// Address returns the host:port form of a peer's address.
func (n *Peer) Address() string {
	return net.JoinHostPort(n.Addr.String(), strconv.FormatUint(uint64(n.Port), 10))
}

// String returns the peer id.
func (n *Peer) String() string {
	return n.ID
}

// Correct returns true if the peer is alive.
func (n *Peer) Correct() bool {
	return !n.Incorrect
}
