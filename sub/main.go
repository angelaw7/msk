package main

import (
	"fmt"
	"msk-sub/server"
	"msk-sub/utils"

	"github.com/nats-io/nats.go"
)

func main() {

	// Get flags
	masterJSONFile, channel, allChannels, natsServer := utils.GetFlags()

	// Check that the file exists and is empty if starting new sub for all channels
	if channel == allChannels {
		utils.CheckFile(masterJSONFile)
	}

	// Create a connection to the NATS server:
	nc, wait := server.ConnectToServer(natsServer)

	// Create hashmap for the sample IDs that have been inserted into master JSON
	idMap := map[string]int{}

	// Subscribe to channel
	nc.Subscribe(channel, func(message *nats.Msg) {

		// Deseralize and print the data
		newStruct := utils.DeserializeAndPrintData(message)

		// Write to master JSON if sub is channels.*
		if channel == allChannels {
			dmpID := newStruct.MetaData.DmpSampleId

			// Read master JSON file
			data := utils.ReadMasterJSON(masterJSONFile)

			// Add sample to JSON or update sample in JSON
			_, sampleExistsInFile := idMap[dmpID]
			data, idMap = utils.AddOrUpdateSample(sampleExistsInFile, idMap, dmpID, data, newStruct)

			// Write the new sample data into the master JSON file
			utils.WriteToMasterJSON(data, masterJSONFile)
		}
	})

	fmt.Println("Subscribed to", channel)
	<-wait
}
