package basic

import (
	"github.com/armen/dp/consensus/paxos"
	"github.com/armen/dp/link"
)

func (n *Node) initAcc() {
	n.decided = false
	n.promBallot = &paxos.Ballot{}
	n.accBallot = &paxos.Ballot{}
	n.accVal = nil

	// Setup the beb.Deliver event dispatcher
	n.beb.Deliver(func(p link.Peer, m link.Message) {
		n.mux <- func() {
			switch msg := m.(type) {
			case paxos.Prepare:
				n.onPrepare(p, msg.Ballot)
			case paxos.Accept:
				n.onAccept(p, msg.Ballot, msg.Val)
			case paxos.Decided:
				n.onDecide(msg.Ballot, msg.Val)
			}
		}
	})
}

func (n *Node) onPrepare(p link.Peer, ballot *paxos.Ballot) {
	if n.promBallot.Less(ballot) {
		n.promBallot = ballot
		go n.pp2p.Send(p, paxos.Promise{ballot, n.accBallot, n.accVal})

		return
	}

	go n.pp2p.Send(p, paxos.Nack{ballot})
}

func (n *Node) onAccept(p link.Peer, ballot *paxos.Ballot, val interface{}) {
	if n.promBallot.Less(ballot) || n.promBallot.Equals(ballot) {
		n.promBallot = ballot
		n.accBallot = ballot
		n.accVal = val
		go n.pp2p.Send(p, paxos.Accepted{ballot, val})

		return
	}

	go n.pp2p.Send(p, paxos.Nack{ballot})
}

func (n *Node) onDecide(ballot *paxos.Ballot, val interface{}) {
	if n.decided {
		return
	}

	go n.paxosDecide(val)
	n.decided = true
}
