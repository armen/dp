// Package p2p implements a TCP based perfect peer-to-peer link.
package p2p

import (
	"github.com/armen/dp/link"
	"github.com/armen/dp/link/node"
)

func (p *P2p) init() {
	s := &server{p2p: p}

	go s.serve(p.listener, "P2P")
}

// Send requests to send message m to process q.
func (p *P2p) Send(q link.Peer, m link.Message) error {
	result := make(chan error, 1)
	p.mux <- func() {
		c, err := p.connect(q)
		if err != nil {
			result <- err
			return
		}

		// Deliver it to ourselves
		if q.ID() == p.ID() {
			go p.deliver(p, m)

			result <- nil
			return
		}

		result <- c.Call("P2P.Recv", &Payload{p.ID(), m}, nil)
		return
	}
	return <-result
}

func (p *P2p) recv(pl *Payload) {
	p.mux <- func() {
		peer := node.New(node.WithID(pl.ID))

		go p.deliver(peer, pl.Message)
	}
}
