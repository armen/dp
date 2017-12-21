package p2p

import (
	"net"
	"net/rpc"

	"github.com/armen/dp/link"
)

func (n *Node) connect(q link.Peer) (*rpc.Client, error) {
	// Check if it's already connected
	if _, ok := n.conn[q.ID()]; ok {
		return n.conn[q.ID()], nil
	}

	addr, err := net.ResolveTCPAddr(q.Addr().Network(), q.Addr().String())
	if err != nil {
		return nil, err
	}

	c, err := net.DialTCP(q.Addr().Network(), nil, addr)
	if err != nil {
		return nil, err
	}

	c.SetKeepAlive(true)
	c.SetKeepAlivePeriod(n.keepAlive)

	n.conn[q.ID()] = rpc.NewClient(c)

	return n.conn[q.ID()], nil
}
