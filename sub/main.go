package main

import (
	"fmt"
	"msk-sub/genome"
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

	// Send an API call to get data from Genome Nexus and write to JSON
	variants := []string{"1:g.182712A>C", "2:g.265023C>T", "3:g.319781del", "19:g.110753dup", "1:g.1385015_1387562del"}
	genomeData := genome.GetGenomeData(variants)
	utils.WriteToGenomeJSON(genomeData, "genomeData.json")

	// Subscribe to channel
	nc.Subscribe(channel, func(message *nats.Msg) {

		// Deseralize and print the data
		sampleResult := utils.DeserializeAndPrintData(message)

		// Write to master JSON if sub is channels.*
		if channel == allChannels {
			dmpID := sampleResult.MetaData.DmpSampleId

			// Read master JSON file
			oldJSONData := utils.ReadMasterJSON(masterJSONFile)

			// Add sample to JSON or update sample in JSON
			var newJSONData []byte
			newJSONData, idMap = utils.AddOrUpdateSample(idMap, dmpID, oldJSONData, sampleResult)

			// Write the new sample data into the master JSON file
			utils.WriteToMasterJSON(newJSONData, masterJSONFile)
		}
	})

	fmt.Println("Subscribed to", channel)
	<-wait
}
