package main

import (
	"fmt"
	"github.com/mkashifaslam/distr-fs/p2p"
	"log"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":4000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("Received message: %+v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
