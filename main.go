package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"time"

	"msk-mongo/types"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	collectionName := "testing"
	fetchJSONFile := "fetch_shorter.json"

	// Setup connection with MongoDB
	ctx, mskCollection := mongoSetup(collectionName)

	// Read JSON file with other data
	newData, err := ioutil.ReadFile(fetchJSONFile)
	if err != nil {
		log.Fatal(err)
	}
	var newJSONFile types.FetchJSON
	err = json.Unmarshal(newData, &newJSONFile)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	newSamples := newJSONFile.Results
	opts := options.FindOne().SetSort(bson.M{"last_modified": -1})

	// Loop through each sample and insert/update the database
	for i := range newSamples {
		newSample := newSamples[i]
		dmp_sample_id := newSample.Meta_data.Dmp_sample_id
		filter := bson.M{"meta_data.dmp_sample_id": dmp_sample_id}

		// Search database for existing sample using dmp_sample_id
		var oldSample types.Result
		err = mskCollection.FindOne(ctx, filter, opts).Decode(&oldSample)
		if err != nil {

			// If no sample exists with dmp_sample_id, then insert new document
			if err == mongo.ErrNoDocuments {
				fmt.Printf("No document with dmp_sample_id %s found; inserting new document\n", dmp_sample_id)
				insertDocument(mskCollection, ctx, newSample)
			} else {
				log.Fatal(err)
			}
		} else { // Only insert new document if different from most recent existing document
			oldSample.Last_modified = nil

			// Sample is different from most recent existing document; insert new
			if !reflect.DeepEqual(newSample, oldSample) {
				fmt.Printf("Document with dmp_sample_id %s found but is different; inserting new version\n", dmp_sample_id)
				insertDocument(mskCollection, ctx, newSample)
			} else { // Sample is the same as most recent existing document; skip
				fmt.Printf("Document with dmp_sample_id %s is the same; skipping\n", dmp_sample_id)
			}
		}
	}
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
	newSamples := JSONFile["results"].([]interface{})
	opts := options.InsertMany().SetOrdered(false)
	res, err := mskCollection.InsertMany(ctx, newSamples, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted docs with IDS %v\n", res.InsertedIDs)
}

// Function for inserting a document to the database
func insertDocument(mskCollection *mongo.Collection, ctx context.Context, newSampleData types.Result) {
	tm := time.Now()

	newSampleData.Last_modified = primitive.NewDateTimeFromTime(tm)
	res, err := mskCollection.InsertOne(ctx, newSampleData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted document with ID %v\n", res.InsertedID)
}

// Setup connection with MongoDB
func mongoSetup(collectionName string) (context.Context, *mongo.Collection) {

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

	mskCollection := client.Database("msk").Collection(collectionName)

	return ctx, mskCollection
}
