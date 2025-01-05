package p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTCPTransport(t *testing.T) {
	opts := TCPTransportOpts{
		ListenAddr: ":4000",
	}
	tr := NewTCPTransport(opts)

	assert.Equal(t, tr.ListenAddr, ":4000")

	assert.Nil(t, tr.ListenAndAccept())
}
