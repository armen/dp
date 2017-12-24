// Package consensus implements consensus abstraction.
package consensus

// Uniform defines the interface and properties of uniform consensus.
//
// Properties:
// 	UC1-UC3: Same as properties C1-C3 in (regular) consensus.
// 	UC4: Uniform agreement:
// 		- No two processes decide differently.
//
type Uniform interface {
	Regular
}
