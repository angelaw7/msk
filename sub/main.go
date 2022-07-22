package main

import (
	"flag"
	"msk-sub/subscriber"
)

func main() {

	masterJSONPtr := flag.String("masterJSONFile", "fetch_shorter.json", "MasterJSON filename")
	channelPtr := flag.String("channel", "channels.*", "Channel to subscribe to")
	allChannelsPtr := flag.String("allChannels", "channels.*", "All channels")
	natsServerPtr := flag.String("natsServer", "nats://localhost:4222", "NATS server connection URL")

	flag.Parse()

	subscriber.SubscriberMain(*masterJSONPtr, *channelPtr, *allChannelsPtr, *natsServerPtr)
}
