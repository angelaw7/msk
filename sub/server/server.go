package server

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

// Connect to NATS server
func ConnectToServer(natsServer string) (*nats.Conn, chan bool) {
	// Set up connection to NATS server
	wait := make(chan bool)
	var nc *nats.Conn
	var err error
	for {
		nc, err = nats.Connect(natsServer)
		if err != nil {
			fmt.Println("Attempting to connect to server...")
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}

	return nc, wait
}
