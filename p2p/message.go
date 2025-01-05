package p2p

import (
	"fmt"
	"net"
)

// RPC holds any arbitrary data that is being sent over
// each transport between two nodes in the network.
type RPC struct {
	From    net.Addr
	Payload []byte
}

func NewRPC(from net.Addr, payload []byte) *RPC {
	return &RPC{
		From:    from,
		Payload: payload,
	}
}

func (m *RPC) Print() {
	fmt.Printf("From: %s\nPayload: %b\n", m.From.String(), m.Payload)
}
