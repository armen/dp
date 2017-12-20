package ppp

import (
	"net"
	"net/rpc"

	"github.com/armen/dp/link"
)

func (p *Ppp) connect(q link.Peer) (*rpc.Client, error) {
	// Check if it's already connected
	if _, ok := p.conn[q.ID()]; ok {
		return p.conn[q.ID()], nil
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
	c.SetKeepAlivePeriod(p.keepAlive)

	p.conn[q.ID()] = rpc.NewClient(c)

	return p.conn[q.ID()], nil
}
