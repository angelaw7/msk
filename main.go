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

	mskCollection := client.Database("msk").Collection("fetchjson")

	// Read from JSON data
	data, err := ioutil.ReadFile("edited_json.json")
	if err != nil {
		log.Fatal(err)
	}
	var JSONFile map[string]interface{}
	err = json.Unmarshal(data, &JSONFile)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	// results := JSONFile["results"].([]interface{})

	/** Insert results from JSON file to Mongo collection **/
	// opts := options.InsertMany().SetOrdered(false)
	// res, err := mskCollection.InsertMany(ctx, results, opts)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("inserted docs with IDS %v\n", res.InsertedIDs)

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

	// Set options and filter (currently hardcoded)
	opts := options.Replace().SetUpsert(true)
	filter := bson.M{"meta-data.dmp_sample_id": 23}

	// Replace document if dmp_sample_id exists, else insert
	result, err := mskCollection.ReplaceOne(ctx, filter, newJSONFile, opts)
	if err != nil {
		log.Fatal(err)
	}

	if result.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")
		return
	}
	if result.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
	}
	// opts := options.Update().SetUpsert(true)
	// filter := bson.M{"results.meta-data.dmp_sample_id": "DUMMYSID_HERE"}
	// result := mskCollection.FindOne(ctx, filter).Decode(&result)
	// fmt.Println(result)

	// // Use the planets collection (db.sample_guides.planets)
	// planetsCollection := client.Database("sample_guides").Collection("planets")
	// id, _ := primitive.ObjectIDFromHex("621ff30d2a3e781873fcb677")
	// opts := options.Update().SetUpsert(true)

	// // Call the collection.UpdateOne() method
	// result, err := planetsCollection.UpdateOne(
	// 	ctx,
	// 	bson.M{"_id": id, "note": "yay"},
	// 	bson.D{
	// 		{"$set", bson.D{{"name", "Pluto + owo"}}},
	// 	},
	// 	opts,
	// )

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Print statements depending on what action was performed
	// if result.ModifiedCount == 1 {
	// 	fmt.Println("Updated 1 Documents")
	// } else if result.MatchedCount == 0 {
	// 	fmt.Println("Inserted 1 document")
	// } else {
	// 	fmt.Println("No documents were inserted/updated")
	// }
}
