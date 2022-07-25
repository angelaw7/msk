package main

import (
	"fmt"
	"log"
	"msk-pub/database"
	"msk-pub/server"
	"msk-pub/utils"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {

	// Get flags
	mongoURI, databaseName, collectionName, fetchJSONFile, newSamplesChannel, updateSamplesChannel, natsServer := utils.GetFlags()

	for {
		// Connect to MongoDB
		mskCollection, ctx, _ := database.ConnectToMongoDB(mongoURI, databaseName, collectionName)

		// Read JSON data
		newSamples := utils.ReadSampleData(fetchJSONFile)

		// Create a connection to the NATS server:
		nc := server.ConnectToServer(natsServer)

		// Loop through each sample and insert/update the database
		for i := range newSamples {
			newSample := newSamples[i]
			dmpSampleID := newSample.Meta_data.Dmp_sample_id

			// Look for existing sample in database
			err, oldSample := database.FindSample(dmpSampleID, newSample, mskCollection, ctx)

			if err != nil {

				// If no sample exists with dmp_sample_id, then insert new document
				if err == mongo.ErrNoDocuments {
					fmt.Printf("No document with dmp_sample_id %s found; inserting new document\n", dmpSampleID)
					database.InsertDocumentIntoMongoDB(mskCollection, ctx, newSample)
					server.PublishMessage(newSample, nc, newSamplesChannel)
				} else {
					log.Fatal(err)
				}
			} else { // Only insert new document if different from most recent existing document
				oldSample.Last_modified = nil

				// Sample is different from most recent existing document; insert new
				if !reflect.DeepEqual(newSample, oldSample) {
					fmt.Printf("Document with dmp_sample_id %s found but is different; inserting new version\n", dmpSampleID)
					database.InsertDocumentIntoMongoDB(mskCollection, ctx, newSample)
					server.PublishMessage(newSample, nc, updateSamplesChannel)

				} else { // Sample is the same as most recent existing document; skip
					fmt.Printf("Document with dmp_sample_id %s is the same; skipping\n", dmpSampleID)
				}
			}
		}

		// Closes connection to the NATS server
		nc.Drain()

		fmt.Print("Waiting 10 seconds until next cycle...\n\n")
		time.Sleep(10 * time.Second)
	}
}
