package link

// Node represents a single node in the cluster, at the same time a node itself
// is a Peer.
type Node interface {
	Peers() []Peer   // Returns the list of peers of the node
	Members() []Peer // Returns all the peers including the current node

	Peer
}
