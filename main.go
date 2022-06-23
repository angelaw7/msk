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
	id, _ := primitive.ObjectIDFromHex("621ff30d2a3e781873fcb661")
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

}
