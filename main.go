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
	godotenv.Load()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal(err)
	}

	planetsCollection := client.Database("sample_guides").Collection("planets")
	id, _ := primitive.ObjectIDFromHex("621ff30d2a3e781873fcb621")
	opts := options.Update().SetUpsert(true)

	result, err := planetsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"name", "Uranus + owo"}}},
		},
		opts,
	)

	if err != nil {
		log.Fatal(err)
	}

	if result.ModifiedCount == 1 {
		fmt.Println("Updated 1 Documents")
	} else if result.MatchedCount == 0 {
		fmt.Println("Inserted 1 document")
	} else {
		fmt.Println("No documents were inserted/updated")
	}
}
