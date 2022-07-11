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

	mskCollection := client.Database("msk").Collection("testing2")

	// If running for the first time, uncomment this to upload all sample data
	// uploadFetchJSONFile(mskCollection, ctx)

	// Read JSON file with other data
	newData, err := ioutil.ReadFile("testing_small.json")
	if err != nil {
		log.Fatal(err)
	}
	var newJSONFile types.FetchJSON
	err = json.Unmarshal(newData, &newJSONFile)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	results := newJSONFile.Results
	opts := options.FindOne().SetSort(bson.M{"last_modified": -1})

	// Loop through each sample and insert/update the database
	for i := range results {
		thisresult := results[i]

		dmp_sample_id := thisresult.Meta_data.Dmp_sample_id
		// dmp_sample_id := int32(results[i].(map[string]interface{})["dmp_sample_id"].(float64))
		// newSampleData := bson.M(results[i].(map[string]interface{}))
		// fmt.Println(newSampleData)
		// this := bson.M{"hello": "testing"}
		// fmt.Println(this)
		// newSampleData["dmp_sample_id"] = dmp_sample_id

		// filter := bson.M{"meta-data.dmp_sample_id": dmp_sample_id}
		filter := bson.M{"meta_data.dmp_sample_id": dmp_sample_id}
		// filter := bson.M{"dmp_sample_id": dmp_sample_id}

		// Replace document if dmp_sample_id exists, else insert
		var result types.Result
		err = mskCollection.FindOne(ctx, filter, opts).Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				fmt.Printf("No document with dmp_sample_id %s found; inserting new document\n", dmp_sample_id)
				fmt.Println("TODO: Insert new document")
				insertDocument(mskCollection, ctx, thisresult)
				return
			}
			log.Fatal(err)
		}

		// delete(result, "last_modified")
		// delete(result, "_id")

		// if !reflect.DeepEqual(thisresult, result) {
		// 	// Insert here
		// 	fmt.Printf("Document with dmp_sample_id %s found but is different; inserting new version\n", dmp_sample_id)
		// 	insertDocument(mskCollection, ctx, thisresult)
		// } else {
		// 	fmt.Printf("Document with dmp_sample_id %s is the same; skipping\n", dmp_sample_id)
		// }
		thisresult.Last_modified = nil
		result.Last_modified = nil

		fmt.Println(thisresult)
		fmt.Println(result)

		if !reflect.DeepEqual(thisresult, result) {
			// Insert here
			fmt.Printf("Document with dmp_sample_id %s found but is different; inserting new version\n", dmp_sample_id)
			insertDocument(mskCollection, ctx, thisresult)
		} else {
			fmt.Printf("Document with dmp_sample_id %s is the same; skipping\n", dmp_sample_id)
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
func insertDocument(mskCollection *mongo.Collection, ctx context.Context, newSampleData types.Result) {
	tm := time.Now()

	newSampleData.Last_modified = primitive.NewDateTimeFromTime(tm)
	res, err := mskCollection.InsertOne(ctx, newSampleData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted document with ID %v\n", res.InsertedID)

}
