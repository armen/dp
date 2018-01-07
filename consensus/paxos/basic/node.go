// Package basic implements single decree paxos algorithm.
package basic

import (
	"github.com/armen/dp/broadcast"
	"github.com/armen/dp/consensus/paxos"
	"github.com/armen/dp/link"
)

// Node implements basic-paxos algorithm.
type Node struct {
	// Perfect Point-To-Point link
	pp2p link.Perfect

	// BestEffort Broadcast
	beb broadcast.BestEffort

	// Decide handler
	paxosDecide  func(v interface{})
	decided      bool
	promises     paxos.Ballots
	ts           uint64        // Logical clock fo Paxos rounds
	numOfAccepts int           // Number of accepts
	pv           interface{}   // Promissed value
	promBallot   *paxos.Ballot // Promissed ballot
	accBallot    *paxos.Ballot // Accepted ballot

	// A multiplexer to run events in a mutually exclusive way
	mux chan func()

	link.Node
}

// New Instantiates a basic-paxos node.
func New(pp2p link.Perfect, beb broadcast.BestEffort) *Node {
	n := &Node{
		pp2p:        pp2p,
		beb:         beb,
		paxosDecide: func(interface{}) {},
		mux:         make(chan func()),
		Node:        beb.(link.Node),
	}

	n.initProp()
	n.initAcc()

	go n.react()

	return n
}

// Decide registers the decide handler.
func (n *Node) Decide(f func(interface{})) {
	n.paxosDecide = f
}

// react mutually executes events.
func (n *Node) react() {
	for f := range n.mux {
		f()
	}
}
