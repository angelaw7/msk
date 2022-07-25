package server

import (
	"encoding/json"
	"fmt"
	"log"
	msk_protobuf "msk-pub/protobuf"
	"msk-pub/types"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// Connect to NATS server
func ConnectToServer(natsServer string) *nats.Conn {
	// Create a connection to the NATS server:
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

	return nc
}

// Publish a message through the NATS server
func PublishMessage(newSample types.Result, nc *nats.Conn, channel string) {

	// Result struct -> JSON
	newSampleBytes, err := json.Marshal(newSample)
	if err != nil {
		log.Fatal(err)
	}

	// JSON -> ProtoMessage
	protoJSON := &msk_protobuf.Result{}
	err = protojson.Unmarshal(newSampleBytes, protoJSON)
	if err != nil {
		log.Fatal(err)
	}

	// ProtoMessage -> []byte
	protoData, err := proto.Marshal(protoJSON)
	if err != nil {
		log.Fatal(err)
	}

	// Send bytes through channel
	nc.Publish(channel, protoData)
}
