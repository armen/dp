// Package beb implements the broadcast.BestEffort interface (Basic Broadcast)
package beb

import (
	"github.com/armen/dp/link"
)

func (n *Node) init() {
	// Register our plDeliver handler
	n.pl.Deliver(n.plDeliver)
}

// Broadcast broadcasts a message m to all processes
func (n *Node) Broadcast(m link.Message) {
	n.mux <- func() {
		for _, q := range n.Members() {
			n.pl.Send(q, m)
		}
	}
}

// Upon plDeliver, delivers the message
func (n *Node) plDeliver(p link.Peer, m link.Message) {
	n.mux <- func() {
		go n.bebDeliver(p, m)
	}
}
