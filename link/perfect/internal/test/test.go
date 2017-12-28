package test

import (
	"log"
	"net"
	"testing"
	"time"

	"github.com/armen/dp/link"
)

// ReliableDelivery tests reliable delivery property.
func ReliableDelivery(ppp link.Perfect, qqq link.Perfect, t *testing.T) {
	var done = make(chan struct{})

	p := ppp.(link.Peer)
	q := qqq.(link.Peer)

	qqq.Deliver(func(from link.Peer, m link.Message) {
		if string(m.([]byte)) != "Hello" {
			t.Error("Delivered message doesn't match the sent message")
		}

		if from.ID() != p.ID() {
			t.Errorf("Expected to receive a message from %.6q but received it from %.6q", p.ID(), from.ID())
		}

		done <- struct{}{}
	})

	ppp.Send(q, []byte("Hello"))

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Error("The message not delivered")
	}
}

// SelfDelivery tests reliable delivery property.
func SelfDelivery(ppp link.Perfect, t *testing.T) {
	var done = make(chan struct{})

	p := ppp.(link.Peer)
	ppp.Deliver(func(from link.Peer, m link.Message) {
		if string(m.([]byte)) != "Hello Self" {
			t.Error("Delivered message doesn't match the sent message")
		}

		if from.ID() != p.ID() {
			t.Errorf("Expected to receive a message from %.6q but received it from %.6q", p.ID(), from.ID())
		}

		done <- struct{}{}
	})

	ppp.Send(p, []byte("Hello Self"))

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Error("The message not delivered")
	}
}

// ListenTCP creates a listener for testing.
func ListenTCP() (net.Listener, net.Addr) {
	l, e := net.Listen("tcp", "127.0.0.1:0") // any available address
	if e != nil {
		log.Fatalf("net.Listen tcp :0: %v", e)
	}

	return l, l.Addr()
}
