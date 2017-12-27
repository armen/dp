package link

// Node represents a single node in the cluster, at the same time a node itself
// is a Peer.
type Node interface {
	Members() []Peer // Returns all the peers including the current node
	Peers() []Peer   // Returns the list of peers of the node
	AddPeer(p Peer)  // Adds a new peer to the peers list
	N() int          // Returns total number of nodes in the cluster

	Peer
}
