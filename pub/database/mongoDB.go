package database

import (
	"context"
	"fmt"
	"log"
	"msk-pub/types"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect to MongoDB database
func ConnectToMongoDB(mongoURI string, databaseName string, collectionName string) (*mongo.Collection, context.Context, context.CancelFunc) {
	// Setup connection with MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	// Connect to database
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	mskCollection := client.Database(databaseName).Collection(collectionName)

	return mskCollection, ctx, cancel
}

// Insert a document to the MongoDB database
func InsertDocumentIntoMongoDB(mskCollection *mongo.Collection, ctx context.Context, newSampleData types.Result) {

	// Add last_modified field to the current time
	tm := time.Now()
	newSampleData.Last_modified = primitive.NewDateTimeFromTime(tm)

	// Insert document into databse
	res, err := mskCollection.InsertOne(ctx, newSampleData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\t- Inserted document with ID %v\n", res.InsertedID)
}

// Look for existing sample in database
func FindSample(dmp_sample_id string, newSample types.Result, mskCollection *mongo.Collection, ctx context.Context) (error, types.Result) {
	// Search database for existing sample using dmp_sample_id
	filter := bson.M{"meta_data.dmp_sample_id": dmp_sample_id}
	sortOpts := options.FindOne().SetSort(bson.M{"last_modified": -1})
	var oldSample types.Result
	err := mskCollection.FindOne(ctx, filter, sortOpts).Decode(&oldSample)

	return err, oldSample
}
