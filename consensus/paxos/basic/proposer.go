package basic

import (
	"github.com/armen/dp/consensus/paxos"
	"github.com/armen/dp/link"
)

func (n *Node) initProp() {
	// Setup the pp2p.Deliver event dispatcher
	n.pp2p.Deliver(func(p link.Peer, m link.Message) {
		n.mux <- func() {
			switch msg := m.(type) {
			case paxos.Promise:
				n.promise(msg.Ballot, msg.Accepted)
			case paxos.Accepted:
				n.accepted(msg.Ballot)
			case paxos.Nack:
				n.nack(msg.Ballot)
			}
		}
	})
}

func (n *Node) propose() {
	if n.decided {
		return
	}

	n.ts++
	n.numOfAccepts = 0
	n.promises = nil

	go n.beb.Broadcast(paxos.Prepare{&paxos.Ballot{n.ts, n.ID(), nil}})
}

// Propose proposes value v for consensus
func (n *Node) Propose(v interface{}) {
	n.mux <- func() {
		n.pv = v
		n.propose()
	}
}

func (n *Node) promise(ballot *paxos.Ballot, accepted *paxos.Ballot) {
	if ballot.Ts != n.ts || ballot.Pid != n.ID() {
		return
	}

	n.promises = append(n.promises, accepted)

	if len(n.promises) != (n.N()+1)/2 {
		return
	}

	maxBallot := n.promises[0]
	if maxBallot.Value != nil {
		n.pv = maxBallot.Value
	}

	go n.beb.Broadcast(paxos.Accept{&paxos.Ballot{n.ts, n.ID(), n.pv}})
}

func (n *Node) accepted(ballot *paxos.Ballot) {
	if ballot.Ts != n.ts || ballot.Pid != n.ID() {
		return
	}

	n.numOfAccepts++
	if n.numOfAccepts != (n.N()+1)/2 {
		return
	}

	go n.beb.Broadcast(paxos.Decided{ballot})
}

func (n *Node) nack(ballot *paxos.Ballot) {
	if ballot.Ts == n.ts && ballot.Pid == n.ID() {
		n.propose()
	}
}
