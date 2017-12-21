package p2p

import (
	"net"
	"net/rpc"
)

type server struct {
	node *Node
}

func (s *server) serve(l net.Listener, name string) {
	rs := rpc.NewServer()
	rs.RegisterName(name, s)
	rs.Accept(l)
}

// Recv implements RPC callback to recieve a meesage from a peer.
func (s *server) Recv(p *Payload, _ *struct{}) error {
	s.node.recv(p)

	return nil
}
