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
				n.promise(msg.Ballot, msg.Accepted, msg.Val)
			case paxos.Accepted:
				n.accepted(msg.Ballot, msg.Val)
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
	n.vals = make(map[*paxos.Ballot]interface{})

	go n.beb.Broadcast(paxos.Prepare{&paxos.Ballot{n.ts, n.ID()}})
}

// Propose proposes value v for consensus
func (n *Node) Propose(v interface{}) {
	n.mux <- func() {
		n.propVal = v
		n.propose()
	}
}

func (n *Node) promise(ballot *paxos.Ballot, accepted *paxos.Ballot, val interface{}) {
	if ballot.Ts != n.ts || ballot.Pid != n.ID() {
		return
	}

	n.vals[accepted] = val
	n.promises = append(n.promises, accepted)

	if len(n.promises) != (n.N()+1)/2 {
		return
	}

	maxBallot := n.promises[0]
	if val, ok := n.vals[maxBallot]; ok && val != nil {
		n.propVal = val
	}

	go n.beb.Broadcast(paxos.Accept{&paxos.Ballot{n.ts, n.ID()}, n.propVal})
}

func (n *Node) accepted(ballot *paxos.Ballot, val interface{}) {
	if ballot.Ts != n.ts || ballot.Pid != n.ID() {
		return
	}

	n.numOfAccepts++
	if n.numOfAccepts != (n.N()+1)/2 {
		return
	}

	go n.beb.Broadcast(paxos.Decided{ballot, val})
}

func (n *Node) nack(ballot *paxos.Ballot) {
	if ballot.Ts == n.ts && ballot.Pid == n.ID() {
		n.propose()
	}
}
