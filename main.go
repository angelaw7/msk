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

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Load the .env file
	godotenv.Load()

	// Setup
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Connect to database
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal(err)
	}

	mskCollection := client.Database("msk").Collection("testing")

	// If running for the first time, uncomment this to upload all sample data
	// uploadFetchJSONFile(mskCollection, ctx)

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
	opts := options.FindOne().SetSort(bson.M{"last-modified": -1})

	// Loop through each sample and insert/update the database
	for i := range results {

		dmp_sample_id := results[i].(map[string]interface{})["meta-data"].(map[string]interface{})["dmp_sample_id"]
		// dmp_sample_id := int32(results[i].(map[string]interface{})["dmp_sample_id"].(float64))
		newSampleData := bson.M(results[i].(map[string]interface{}))
		newSampleData["dmp_sample_id"] = dmp_sample_id

		filter := bson.M{"meta-data.dmp_sample_id": dmp_sample_id}
		// filter := bson.M{"dmp_sample_id": dmp_sample_id}

		// Replace document if dmp_sample_id exists, else insert
		var result bson.M
		err := mskCollection.FindOne(ctx, filter, opts).Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				fmt.Printf("No document with dmp_sample_id %d found; inserting new document\n", dmp_sample_id)
				fmt.Println("TODO: Insert new document")

				// insertDocument(mskCollection, ctx, newSampleData)
				return
			}
			log.Fatal(err)
		}

		delete(result, "last-modified")
		delete(result, "_id")

		if !reflect.DeepEqual(newSampleData, result) {
			// Insert here
			fmt.Printf("Document with dmp_sample_id %d found but has changes; inserting new version\n", dmp_sample_id)
			// insertDocument(mskCollection, ctx, newSampleData)
		} else {
			fmt.Printf("Document with dmp_sample_id %d is the same; skipping\n", dmp_sample_id)
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
	results := JSONFile["results"].([]interface{})
	opts := options.InsertMany().SetOrdered(false)
	res, err := mskCollection.InsertMany(ctx, results, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted docs with IDS %v\n", res.InsertedIDs)
}

// Function for inserting a document to the database
func insertDocument(mskCollection *mongo.Collection, ctx context.Context, newSampleData primitive.M) {
	tm := time.Now()

	newSampleData["last-modified"] = primitive.NewDateTimeFromTime(tm)
	res, err := mskCollection.InsertOne(ctx, newSampleData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted document with ID %v\n", res.InsertedID)

}
