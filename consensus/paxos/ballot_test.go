package paxos_test

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/armen/dp/consensus/paxos"
)

var lesstests = []struct {
	b1     *paxos.Ballot
	b2     *paxos.Ballot
	expect bool
}{
	{&paxos.Ballot{1, "abc"}, &paxos.Ballot{2, "def"}, true},
	{&paxos.Ballot{2, "abc"}, &paxos.Ballot{2, "def"}, true},
	{&paxos.Ballot{2, "def"}, &paxos.Ballot{2, "abc"}, false},
	{&paxos.Ballot{1, "abc"}, &paxos.Ballot{1, "abc"}, false},
}

func TestLess(t *testing.T) {
	for _, tt := range lesstests {
		if got := tt.b1.Less(tt.b2); got != tt.expect {
			t.Errorf("%s < %s => %v, want %v", tt.b1, tt.b2, got, tt.expect)
		}
	}
}

func TestEquals(t *testing.T) {
	b1 := &paxos.Ballot{1, "abc"}
	b2 := &paxos.Ballot{1, "abc"}

	if !b1.Equals(b2) {
		t.Errorf("want %s == %s", b1, b2)
	}
}

func Example() {
	ballots := paxos.Ballots{
		&paxos.Ballot{2, "abc"},
		&paxos.Ballot{1, "abc"},
		&paxos.Ballot{2, "def"},
	}

	heap.Init(&ballots)

	fmt.Println("Max ballot:", ballots[0])

	for ballots.Len() > 0 {
		b := heap.Pop(&ballots).(*paxos.Ballot)
		fmt.Println(b)
	}

	// Output:
	// Max ballot: (2, "def")
	// (2, "def")
	// (2, "abc")
	// (1, "abc")
}
