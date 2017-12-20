// Package ppp implements a TCP based perfect peer-to-peer link.
package ppp

import (
	"github.com/armen/dp/link"
)

func (p *Ppp) init() {
	s := &server{ppp: p}

	go s.serve(p.listener, "PPP")
}

// Send requests to send message m to process q.
func (p *Ppp) Send(q link.Peer, m link.Message) error {
	result := make(chan error, 1)
	p.mux <- func() {
		c, err := p.connect(q)
		if err != nil {
			result <- err
			return
		}

		result <- c.Call("PPP.Recv", &Payload{p.ID(), m}, nil)
		return
	}
	return <-result
}

func (p *Ppp) recv(pl *Payload) {
	p.mux <- func() {
		peer := link.NewNode(nil)
		peer.SetID(pl.Src)

		go p.deliver(peer, pl.Message)
	}
}
