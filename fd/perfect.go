// Package fd implements failure detector abstraction.
package fd

import (
	"github.com/armen/dp"
)

// Perfect defines the interface and properties of a perfect failure detector.
//
// Properties:
// 	PFD1: Strong completeness
// 		- Eventually, every process that crashes is permanently detected by
// 		  every correct process.
// 	PFD2: Strong accuracy:
// 		- If a process p is detected by any process, then p has crashed.
//
type Perfect interface {
	Crash(func(*dp.Peer)) // Detects that process p has crashed

	dp.Reactor
}
