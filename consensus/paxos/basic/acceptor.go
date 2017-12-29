package basic

import (
	"github.com/armen/dp/consensus/paxos"
	"github.com/armen/dp/link"
)

func (n *Node) initAcc() {
	n.decided = false
	n.promBallot = &paxos.Ballot{}
	n.accBallot = &paxos.Ballot{}

	// Setup the beb.Deliver event dispatcher
	n.beb.Deliver(func(p link.Peer, m link.Message) {
		n.mux <- func() {
			switch msg := m.(type) {
			case paxos.Prepare:
				n.prepare(p, msg.Ballot)
			case paxos.Accept:
				n.accept(p, msg.Ballot)
			case paxos.Decided:
				n.decide(msg.Ballot)
			}
		}
	})
}

func (n *Node) prepare(p link.Peer, ballot *paxos.Ballot) {
	if n.promBallot.Less(ballot) {
		n.promBallot = ballot
		go n.pp2p.Send(p, paxos.Promise{ballot, n.accBallot})

		return
	}

	go n.pp2p.Send(p, paxos.Nack{ballot})
}

func (n *Node) accept(p link.Peer, ballot *paxos.Ballot) {
	if n.promBallot.Less(ballot) || n.promBallot.Equals(ballot) {
		n.promBallot = ballot
		n.accBallot = ballot
		go n.pp2p.Send(p, paxos.Accepted{ballot})

		return
	}

	go n.pp2p.Send(p, paxos.Nack{ballot})
}

func (n *Node) decide(ballot *paxos.Ballot) {
	if n.decided {
		return
	}

	go n.paxosDecide(ballot.Value)
	n.decided = true
}
