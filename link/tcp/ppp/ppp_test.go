package ppp_test

import (
	"net"
	"testing"
	"time"

	"github.com/armen/dp/link/internal/test"
	"github.com/armen/dp/link/node"
	"github.com/armen/dp/link/tcp/ppp"
)

type badAddr struct{}

func (_ *badAddr) Network() string { return "tcp" }
func (_ *badAddr) String() string  { return "bad-addr" }

func TestReliableDelivery(t *testing.T) {
	l, addr := test.ListenTCP()
	p := ppp.New(node.New(node.WithAddr(addr)), l, 1*time.Second)

	l, addr = test.ListenTCP()
	q := ppp.New(node.New(node.WithAddr(addr)), l, 1*time.Second)

	test.ReliableDelivery(p, q, t)
}

func TestUnresolvableAddr(t *testing.T) {
	l, _ := test.ListenTCP()
	p := ppp.New(node.New(node.WithAddr(&badAddr{})), l, 0)
	go p.React()

	err := p.Send(p, []byte("message"))
	if _, ok := err.(*net.AddrError); !ok {
		t.Error("expected net.AddrError")
	}
}

func TestEmptyAddr(t *testing.T) {
	l, _ := test.ListenTCP()
	p := ppp.New(node.New(node.WithAddr(&net.TCPAddr{})), l, 0)
	go p.React()

	err := p.Send(p, []byte("message"))
	if e, ok := err.(*net.OpError); !ok || e.Op != "dial" {
		t.Error("expected dial error")
	}
}
