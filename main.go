package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Loads the .env file
	godotenv.Load()

	// Setup
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Connect to database
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal(err)
	}

	mskCollection := client.Database("msk").Collection("sample_data")

	// If running for the first time, uncomment this to upload all sample data
	uploadFetchJSONFile(mskCollection, ctx)

	// Read JSON file with other data
	newData, err := ioutil.ReadFile("new_input.json")
	if err != nil {
		log.Fatal(err)
	}
	var newJSONFile map[string]interface{}
	err = json.Unmarshal(newData, &newJSONFile)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	results := newJSONFile["results"].([]interface{})
	opts := options.Replace().SetUpsert(true)

	totalUpdatedCount := 0
	totalInsertedCount := 0

	// Loop through each sample and insert/update the database
	for i := range results {
		dmp_sample_id := results[i].(map[string]interface{})["meta-data"].(map[string]interface{})["dmp_sample_id"]
		filter := bson.M{"meta-data.dmp_sample_id": dmp_sample_id}
		replacement := results[i]

		// Replace document if dmp_sample_id exists, else insert
		result, err := mskCollection.ReplaceOne(ctx, filter, replacement, opts)
		if err != nil {
			log.Fatal(err)
		}

		if result.MatchedCount != 0 {
			totalUpdatedCount++
		}
		if result.UpsertedCount != 0 {
			totalInsertedCount++
		}
	}

	fmt.Printf("matched and replaced %d existing document(s)\n", totalUpdatedCount)
	fmt.Printf("inserted %d new document(s)\n", totalInsertedCount)
}

// Function for uploading fetchjson.json for the first time
func uploadFetchJSONFile(mskCollection *mongo.Collection, ctx context.Context) {

	// Read data from fetchjson.json
	data, err := ioutil.ReadFile("edited_json.json")
	if err != nil {
		log.Fatal(err)
	}
	var JSONFile map[string]interface{}
	err = json.Unmarshal(data, &JSONFile)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// Insert results from JSON file into MongoDB
	results := JSONFile["results"].([]interface{})
	opts := options.InsertMany().SetOrdered(false)
	res, err := mskCollection.InsertMany(ctx, results, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted docs with IDS %v\n", res.InsertedIDs)
}
