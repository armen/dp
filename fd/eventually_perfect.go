// Package fd implements failure detector abstraction.
package fd

import (
	"github.com/armen/dp/link"
)

// EventuallyPerfect defines the interface and properties of the eventually
// perfect failure detector.
//
// Properties:
// 	EPFD1: Strong completeness
// 		- Eventually, every process that crashes is permanently suspected by
// 		  every correct process.
// 	EPFD2: Eventual strong accuracy:
// 		- Eventually, no correct process is suspected by any correct process.
//
type EventuallyPerfect interface {
	Suspect(func(link.Peer)) // Notifies that process p is suspected to have crashed
	Restore(func(link.Peer)) // Notifies that process p is not suspected anymore
}
