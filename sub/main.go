package main

import (
	"flag"
	"fmt"
	"msk-sub/subscriber"
	"os"
)

func main() {

	masterJSONPtr := flag.String("masterJSONFile", "fetch_shorter.json", "MasterJSON filename")
	channelPtr := flag.String("channel", "channels.*", "Channel to subscribe to")
	natsServerPtr := flag.String("natsServer", "nats://localhost:4222", "NATS server connection URL")

	flag.Parse()
	for i := range os.Args {
		fmt.Println(os.Args[i])
	}

	subscriber.SubscriberMain(*masterJSONPtr, *channelPtr, *natsServerPtr)
}
