// Package p2p implements a TCP based perfect peer-to-peer link.
package p2p

import (
	"github.com/armen/dp/link"
	"github.com/armen/dp/link/node"
)

func (n *Node) init() {
	s := &server{node: n}

	go s.serve(n.listener, "p2p")
}

// Send requests to send message m to process q.
func (n *Node) Send(q link.Peer, m link.Message) error {
	result := make(chan error, 1)
	n.mux <- func() {
		c, err := n.connect(q)
		if err != nil {
			result <- err
			return
		}

		// Deliver it to ourselves
		if q.ID() == n.ID() {
			go n.deliver(n, m)

			result <- nil
			return
		}

		result <- c.Call("p2p.Recv", &Payload{n.ID(), m}, nil)
		return
	}
	return <-result
}

func (n *Node) recv(pl *Payload) {
	n.mux <- func() {
		peer := node.New(node.WithID(pl.ID))

		go n.deliver(peer, pl.Message)
	}
}
