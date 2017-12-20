package ppp_test

import (
	"testing"
	"time"

	"github.com/armen/dp/link"
	"github.com/armen/dp/link/internal/test"
	"github.com/armen/dp/link/tcp/ppp"
)

func TestReliableDelivery(t *testing.T) {
	l, addr := test.ListenTCP()
	p := ppp.New(link.NewNode(addr), l, 1*time.Second)

	l, addr = test.ListenTCP()
	q := ppp.New(link.NewNode(addr), l, 1*time.Second)

	test.ReliableDelivery(p, q, t)
}
