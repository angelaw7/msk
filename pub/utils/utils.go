package utils

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"msk-pub/types"
	"os"

	"github.com/joho/godotenv"
)

// Get argument flags
func GetFlags() (string, string, string, string, string, string, string) {
	godotenv.Load()

	mongoURIPtr := flag.String("mongoURI", os.Getenv("MONGODB_URI"), "MongoDB URI")
	databaseNamePtr := flag.String("databaseName", "msk", "MongoDB Database name")
	collectionNamePtr := flag.String("collectionName", "testing", "MongoDB Collection name")
	fetchJSONFilePtr := flag.String("fetchJSONFile", "fetch_shorter.json", "FetchJSON filename")
	newSamplesChannelPtr := flag.String("newSampleChannel", "channels.newSamples", "Channel for publishing new samples")
	updateSamplesChannelPtr := flag.String("updateSampleChannel", "channels.updateSamples", "Channel for publishing updated samples")
	natsServerPtr := flag.String("natsServer", "nats://localhost:4222", "NATS server connection URL")

	flag.Parse()

	mongoURI := *mongoURIPtr
	databaseName := *databaseNamePtr
	collectionName := *collectionNamePtr
	fetchJSONFile := *fetchJSONFilePtr
	newSamplesChannel := *newSamplesChannelPtr
	updateSamplesChannel := *updateSamplesChannelPtr
	natsServer := *natsServerPtr

	return mongoURI, databaseName, collectionName, fetchJSONFile, newSamplesChannel, updateSamplesChannel, natsServer
}

// Read sample data from fetchJSON file
func ReadSampleData(fetchJSONFile string) []types.Result {
	// Read JSON file with other data
	newData, err := ioutil.ReadFile(fetchJSONFile)
	if err != nil {
		log.Fatal(err)
	}

	// Read JSON into Go format
	var newJSONFile types.FetchJSON
	err = json.Unmarshal(newData, &newJSONFile)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	newSamples := newJSONFile.Results

	return newSamples
}
