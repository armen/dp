// Package p2p implements a TCP based perfect peer-to-peer link.
package p2p

import (
	"github.com/armen/dp/link"
	"github.com/armen/dp/link/node"
)

func (n *Node) init() {
	s := &server{n}

	go s.serve(n.channel)
}

// Send requests to send message m to process q.
func (n *Node) Send(q link.Peer, m link.Message) error {
	result := make(chan error, 1)
	n.mux <- func() {
		// Deliver it to ourselves
		if q.ID() == n.ID() {
			go n.deliver(n.Node.(link.Peer), m)

			result <- nil
			return
		}

		c, err := n.connect(q)
		if err != nil {
			result <- err
			return
		}

		result <- c.Call(n.channel+".Recv", &Payload{n.ID(), n.Addr(), m}, nil)
		return
	}
	return <-result
}

func (n *Node) recv(pl *Payload) {
	n.mux <- func() {
		peer := node.NewPeer(pl.ID, pl.Addr)
		go n.deliver(peer, pl.Message)
	}
}
