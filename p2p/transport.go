package p2p

// Peer is an interface which represents the remote node.
type Peer interface {
	Close() error
}

// Transport is anything that handles communication
// between the nodes in the network.This can be of the
// (TCP, UDP, websockets etc.)
type Transport interface {
	ListAndAccept() error
	Consume() <-chan RPC
}
