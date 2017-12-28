// Package consensus implements consensus abstraction.
package consensus

// Regular defines the interface and properties of (regular) consensus.
//
// Properties:
// 	C1: Termination:
// 		- Every correct process eventually decides some value.
// 	C2: Validity:
// 		- If a process decides v, then v was proposed by some process.
// 	C3: Integrity:
// 		- No process decides twice.
// 	C4: Agreement:
// 		- No two correct processes decide differently.
//
type Regular interface {
	Propose(v interface{})      // Proposes value v for consensus
	Decide(func(v interface{})) // Outputs a decided value v of consensus
}
