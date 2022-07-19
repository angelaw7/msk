package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	msk_protobuf "msk-mongo/protobuf"
	"os"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func main() {

	allChannels := "channels.*"
	filename := "sub.json"

	// Gets the channel to subscribe to
	channel := os.Args[1]

	// Check that the file exists
	err := checkFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Set up connection to NATS server
	wait := make(chan bool)
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println(err)
	}

	// Create hashmap for the sample IDs that have been inserted into master JSON
	idMap := map[string]int{}

	// Subscribe to channel
	nc.Subscribe(channel, func(m *nats.Msg) {

		// Read amd print the data
		newStruct := &msk_protobuf.Result{}
		err = proto.Unmarshal(m.Data, newStruct)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(newStruct)

		// Write to master JSON if sub is channels.*
		if channel == allChannels {

			// Read master JSON
			file, err := ioutil.ReadFile(filename)
			if err != nil {
				log.Fatal(err)
			}
			data := []msk_protobuf.Result{}
			json.Unmarshal(file, &data)

			// Sample data ID and check if in map
			dmpID := newStruct.MetaData.DmpSampleId
			_, check := idMap[dmpID]

			// Check whether another version of the sample is already in master JSON file
			if check {
				fmt.Println("Another version of this sample was uploaded already in current JSON; rewriting old one...")
				indexOfExistingSample := idMap[dmpID]
				data = append(data[:indexOfExistingSample], data[indexOfExistingSample+1:]...)
			}
			data = append(data, *newStruct)
			idMap[dmpID] = len(data) - 1

			// Write the new sample data into the master JSON file
			dataBytes, err := json.Marshal(data)
			if err != nil {
				log.Fatal(err)
			}
			err = ioutil.WriteFile(filename, dataBytes, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}

	})

	fmt.Println("Subscribed to", channel)
	<-wait
}

// Function to check that the master JSON file name exits; creates new file if not
func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}
