package p2p_test

import (
	"net"
	"testing"

	"github.com/armen/dp/link/node"
	"github.com/armen/dp/link/perfect/internal/test"
	"github.com/armen/dp/link/perfect/p2p"
)

type badAddr struct{}

func (_ *badAddr) Network() string { return "tcp" }
func (_ *badAddr) String() string  { return "bad-addr" }

func TestReliableDelivery(t *testing.T) {
	p := p2p.New(p2p.WithDefault)
	q := p2p.New(p2p.WithDefault)

	test.ReliableDelivery(p, q, t)
}

func TestSelfDelivery(t *testing.T) {
	p := p2p.New(p2p.WithDefault)

	test.SelfDelivery(p, t)
}

func TestUnresolvableAddr(t *testing.T) {
	n := node.New(node.WithDefault)
	p := p2p.New(p2p.WithNode(n))
	q := node.NewPeer("bad", &badAddr{})

	err := p.Send(q, []byte("message"))
	if _, ok := err.(*net.AddrError); !ok {
		t.Error("expected net.AddrError", err, n.Addr())
	}
}

func TestEmptyAddr(t *testing.T) {
	n := node.New(node.WithDefault)
	p := p2p.New(p2p.WithNode(n))
	q := node.NewPeer("bad", &net.TCPAddr{})

	err := p.Send(q, []byte("message"))
	if e, ok := err.(*net.OpError); !ok || e.Op != "dial" {
		t.Error("expected dial error")
	}
}
