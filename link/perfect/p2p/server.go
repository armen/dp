package p2p

import (
	"encoding/gob"
	"log"
	"net"
	"net/rpc"
	"sync"
)

var (
	rpcChannels    = make(map[string]*rpc.Server)
	rpcChannelsMux sync.Mutex
)

func init() {
	gob.Register(&net.TCPAddr{})
}

type server struct {
	Node *Node
}

func (s *server) serve(name string) {
	rpcChannelsMux.Lock()
	defer rpcChannelsMux.Unlock()

	l := s.Node.Listener()
	addr := l.Addr().String()
	if _, ok := rpcChannels[addr]; !ok {
		rpcChannels[addr] = rpc.NewServer()
		go rpcChannels[addr].Accept(l)
	}

	err := rpcChannels[addr].RegisterName(name, s)
	if err != nil {
		log.Fatal(err)
	}
}

// Recv implements RPC callback to recieve a meesage from a peer.
func (s *server) Recv(p *Payload, _ *struct{}) error {
	go s.Node.recv(p)

	return nil
}
