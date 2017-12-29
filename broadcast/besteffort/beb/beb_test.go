package beb_test

import (
	"testing"
	"time"

	"github.com/armen/dp/broadcast/besteffort/beb"
	"github.com/armen/dp/link"
	"github.com/armen/dp/link/perfect/p2p"
)

// TestValidity tests validity property of best-effort broadcast.
func TestValidity(t *testing.T) {
	var done = make(chan struct{})

	b1 := beb.New(p2p.New(p2p.WithDefault))
	b2 := beb.New(p2p.New(p2p.WithDefault))
	b3 := beb.New(p2p.New(p2p.WithDefault))

	b1.AddPeer(b2)
	b1.AddPeer(b3)

	b2.AddPeer(b1)
	b2.AddPeer(b3)

	b3.AddPeer(b1)
	b3.AddPeer(b2)

	deliver := func(from link.Peer, m link.Message) {
		if m.(string) != "Hello from b1" {
			t.Error("Delivered message doesn't match the sent message")
		}

		if from.ID() != b1.Node.ID() {
			t.Errorf("Expected to receive a message from %.6q but received it from %.6q", b1.Node.ID(), from.ID())
		}

		done <- struct{}{}
	}

	b1.Deliver(deliver)
	b2.Deliver(deliver)
	b3.Deliver(deliver)

	b1.Broadcast("Hello from b1")

	// We're expecting two deliveries, one to the self, and the other to the peer
	for i := 0; i < 3; i++ {
		select {
		case <-done:
		case <-time.After(100 * time.Millisecond):
			t.Error("The message not delivered")
		}
	}
}

func TestConsecutiveBcast(t *testing.T) {
	var done = make(chan struct{})

	b1 := beb.New(p2p.New(p2p.WithDefault))
	b2 := beb.New(p2p.New(p2p.WithDefault))
	b1.AddPeer(b2)
	b2.AddPeer(b1)

	deliver := func(from link.Peer, m link.Message) {
		if m.(string) != "Hello" {
			t.Error("Delivered message doesn't match the sent message")
		}

		done <- struct{}{}
	}

	b1.Deliver(deliver)
	b2.Deliver(deliver)

	b1.Broadcast("Hello")
	b2.Broadcast("Hello")
	b1.Broadcast("Hello")

	// We're expecting two deliveries, one to the self, and the other to the peer
	for i := 0; i < 6; i++ {
		select {
		case <-done:
		case <-time.After(100 * time.Millisecond):
			t.Error("The message not delivered")
		}
	}
}
