package p2p

// HandshakeFunc ...?
type HandshakeFunc func(*TCPPeer) error

func NOPHandshakeFunc(TCPPeer) error {
	return nil
}
