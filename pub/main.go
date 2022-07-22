package main

import (
	"flag"
	"msk-pub/publisher"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	mongoURIPtr := flag.String("mongoURI", os.Getenv("MONGODB_URI"), "MongoDB URI")
	dbPtr := flag.String("databaseName", "msk", "MongoDB Database name")
	collPtr := flag.String("collectionName", "testing", "MongoDB Collection name")
	fetchJSONPtr := flag.String("fetchJSONFile", "fetch_shorter.json", "FetchJSON filename")
	newSamplesChannelPtr := flag.String("newSampleChannel", "channels.newSamples", "Channel for publishing new samples")
	updateSamplesChannelPtr := flag.String("updateSampleChannel", "channels.updateSamples", "Channel for publishing updated samples")
	natsServerPtr := flag.String("natsServer", "nats://localhost:4222", "NATS server connection URL")

	flag.Parse()
	for {
		publisher.PublisherMain(*mongoURIPtr, *dbPtr, *collPtr, *fetchJSONPtr, *newSamplesChannelPtr, *updateSamplesChannelPtr, *natsServerPtr)
		time.Sleep(20 * time.Second)
	}
}
