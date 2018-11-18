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
				n.onPromise(msg.Ballot, msg.Accepted, msg.Val)
			case paxos.Accepted:
				n.onAccepted(msg.Ballot, msg.Val)
			case paxos.Nack:
				n.onNack(msg.Ballot)
			}
		}
	})
}

func (n *Node) onPropose() {
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
		n.onPropose()
	}
}

func (n *Node) onPromise(ballot *paxos.Ballot, accepted *paxos.Ballot, val interface{}) {
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

	b := &paxos.Ballot{n.ts, n.ID()}
	go n.beb.Broadcast(paxos.Accept{b, n.propVal})
}

func (n *Node) onAccepted(ballot *paxos.Ballot, val interface{}) {
	if ballot.Ts != n.ts || ballot.Pid != n.ID() {
		return
	}

	n.numOfAccepts++
	if n.numOfAccepts != (n.N()+1)/2 {
		return
	}

	go n.beb.Broadcast(paxos.Decided{ballot, val})
}

func (n *Node) onNack(ballot *paxos.Ballot) {
	if ballot.Ts == n.ts && ballot.Pid == n.ID() {
		n.onPropose()
	}
}
