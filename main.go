package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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

	// Use the planets collection (db.sample_guides.planets)
	planetsCollection := client.Database("sample_guides").Collection("planets")
	id, _ := primitive.ObjectIDFromHex("621ff30d2a3e781873fcb677")
	opts := options.Update().SetUpsert(true)

	// Call the collection.UpdateOne() method
	result, err := planetsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id, "note": "yay"},
		bson.D{
			{"$set", bson.D{{"name", "Pluto + owo"}}},
		},
		opts,
	)

	if err != nil {
		log.Fatal(err)
	}

	// Print statements depending on what action was performed
	if result.ModifiedCount == 1 {
		fmt.Println("Updated 1 Documents")
	} else if result.MatchedCount == 0 {
		fmt.Println("Inserted 1 document")
	} else {
		fmt.Println("No documents were inserted/updated")
	}
}
