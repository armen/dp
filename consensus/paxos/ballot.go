package paxos

import (
	"fmt"
)

// Ballot is a Paxos ballot containing a logical clock and the process id.
type Ballot struct {
	Ts    uint64      // Timestamp
	Pid   string      // Process id
	Value interface{} // The associated value
}

// String returns string representation of a ballot.
func (b *Ballot) String() string {
	return fmt.Sprintf("(%d, %q)", b.Ts, b.Pid)
}

// Less reports whether the ballot is less than the "from" ballot.
func (b *Ballot) Less(from *Ballot) bool {
	return b.Ts < from.Ts || (b.Ts == from.Ts && b.Pid < from.Pid)
}

// Equals reports whether the ballot is equal "to" ballot.
func (b *Ballot) Equals(to *Ballot) bool {
	return b.Ts == to.Ts && b.Pid == to.Pid
}

// Ballots is a set of ballots.
type Ballots []*Ballot

// Len is the number of ballots in the set.
func (b Ballots) Len() int {
	return len(b)
}

// Swap swaps the ballots with indexes i and j.
func (b Ballots) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// Less reports whether the ballot with index i should sort before the ballot
// with index j.
func (b Ballots) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, ballot so we use greater than here.
	return b[i].Ts > b[j].Ts || (b[i].Ts == b[j].Ts && b[i].Pid > b[j].Pid)
}

// Push pushes a ballot at the end of the ballots heap.
func (b *Ballots) Push(x interface{}) {
	ballot := x.(*Ballot)
	*b = append(*b, ballot)
}

// Pop removes the maximum ballot from the heap.
func (b *Ballots) Pop() interface{} {
	old := *b
	n := len(old)
	ballot := old[n-1]
	*b = old[0 : n-1]

	return ballot
}
