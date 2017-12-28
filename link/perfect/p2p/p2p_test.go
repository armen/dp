package p2p_test

import (
	"net"
	"testing"
	"time"

	"github.com/armen/dp/link/node"
	"github.com/armen/dp/link/perfect/internal/test"
	"github.com/armen/dp/link/perfect/p2p"
)

type badAddr struct{}

func (_ *badAddr) Network() string { return "tcp" }
func (_ *badAddr) String() string  { return "bad-addr" }

func TestReliableDelivery(t *testing.T) {
	l, addr := test.ListenTCP()
	p := p2p.New(node.New(node.WithAddr(addr)), l, 1*time.Second)

	l, addr = test.ListenTCP()
	q := p2p.New(node.New(node.WithAddr(addr)), l, 1*time.Second)

	test.ReliableDelivery(p, q, t)
}

func TestSelfDelivery(t *testing.T) {
	l, addr := test.ListenTCP()
	p := p2p.New(node.New(node.WithAddr(addr)), l, 1*time.Second)

	test.SelfDelivery(p, t)
}

func TestUnresolvableAddr(t *testing.T) {
	l, _ := test.ListenTCP()
	p := p2p.New(node.New(node.WithAddr(&badAddr{})), l, 0)

	err := p.Send(p, []byte("message"))
	if _, ok := err.(*net.AddrError); !ok {
		t.Error("expected net.AddrError")
	}
}

func TestEmptyAddr(t *testing.T) {
	l, _ := test.ListenTCP()
	p := p2p.New(node.New(node.WithAddr(&net.TCPAddr{})), l, 0)

	err := p.Send(p, []byte("message"))
	if e, ok := err.(*net.OpError); !ok || e.Op != "dial" {
		t.Error("expected dial error")
	}
}
