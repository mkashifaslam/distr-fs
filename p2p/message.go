package p2p

import (
	"fmt"
	"net"
)

// Message holds any arbitrary data that is being sent over
// each transport between two nodes in the network.
type Message struct {
	From    net.Addr
	Payload []byte
}

func NewMessage(from net.Addr, payload []byte) *Message {
	return &Message{
		From:    from,
		Payload: payload,
	}
}

func (m *Message) Print() {
	fmt.Printf("From: %s\nPayload: %b\n", m.From.String(), m.Payload)
}
