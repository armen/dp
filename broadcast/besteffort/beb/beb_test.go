package beb_test

import (
	"log"
	"net"
	"testing"
	"time"

	"github.com/armen/dp/broadcast"
	"github.com/armen/dp/broadcast/besteffort/beb"
	"github.com/armen/dp/link"
	"github.com/armen/dp/link/node"
	"github.com/armen/dp/link/perfect/p2p"
)

func newBeb() (broadcast.BestEffort, link.Node) {
	n, l := listenTCP()
	pl := p2p.New(n, l, 1*time.Second)

	return beb.New(pl), n
}

// TestValidity tests validity property of best-effort broadcast.
func TestValidity(t *testing.T) {
	var done = make(chan struct{})

	b1, n1 := newBeb()
	b2, n2 := newBeb()

	n1.AddPeer(n2)
	n2.AddPeer(n1)

	deliver := func(from link.Peer, m link.Message) {
		if m.(string) != "Hello from b1" {
			t.Error("Delivered message doesn't match the sent message")
		}

		if from.ID() != n1.ID() {
			t.Errorf("Expected to receive a message from %.6q but received it from %.6q", n1.ID(), from.ID())
		}

		done <- struct{}{}
	}

	b1.Deliver(deliver)
	b2.Deliver(deliver)

	go b1.React()
	go b2.React()

	b1.Broadcast("Hello from b1")

	// We're expecting two deliveries, one to the self, and the other to the peer
	for i := 0; i < 2; i++ {
		select {
		case <-done:
		case <-time.After(100 * time.Millisecond):
			t.Error("The message not delivered")
		}
	}
}

// listenTCP creates a listener and a node for testing.
func listenTCP() (link.Node, net.Listener) {
	l, e := net.Listen("tcp", "127.0.0.1:0") // any available address
	if e != nil {
		log.Fatalf("net.Listen tcp :0: %v", e)
	}

	n := node.New(node.WithAddr(l.Addr()))

	return n, l
}
