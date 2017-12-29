package basic_test

import (
	"testing"
	"time"

	"github.com/armen/dp/broadcast/besteffort/beb"
	"github.com/armen/dp/consensus/paxos/basic"
	"github.com/armen/dp/link/node"
	"github.com/armen/dp/link/perfect/p2p"
)

func newPaxos(id string) *basic.Node {
	n := node.New(node.WithDefault, node.WithID(id))
	pp2p := p2p.New(p2p.WithNode(n), p2p.WithChannel("beb"))
	beb := beb.New(pp2p)

	return basic.New(p2p.New(p2p.WithNode(n)), beb)
}

func TestSingleDecree(t *testing.T) {
	s1 := newPaxos("1")
	s2 := newPaxos("2")
	s3 := newPaxos("3")

	s1.AddPeer(s2)
	s1.AddPeer(s3)

	s2.AddPeer(s1)
	s2.AddPeer(s3)

	s3.AddPeer(s1)
	s3.AddPeer(s2)

	decided := make(chan bool)
	s3.Decide(func(v interface{}) {
		if val, ok := v.(string); !ok || val != "Single Decree Paxos" {
			t.Errorf("Decided incorrect value %q", val)
		}
		decided <- true
	})

	s2.Propose("Single Decree Paxos")
	s3.Propose("Single Decree Paxos")
	s1.Propose("Single Decree Paxos")

	select {
	case <-decided:
	case <-time.After(3 * time.Second):
		t.Error("Not decided")
	}
}
