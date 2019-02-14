package main

import (
	"context"
	"fmt"
	"log"
	"pushservice/models"

	"github.com/globalsign/mgo/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	defer client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	siteID, _ := primitive.ObjectIDFromHex("5c5b9ee1752a41a63268c3bb")

	// var site models.Site
	// filter := bson.D{{"_id", siteID}}

	// siteCollection := client.Database("pushservice").Collection("sites")
	// err = siteCollection.FindOne(context.TODO(), filter).Decode(&site)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Found a single document: %+v\n", site)

	// Finding multiple documents returns a cursor
	notificationCollection := client.Database("pushservice").Collection("notifications")
	findOptions := options.Find()
	findOptions.SetLimit(2)

	var results []*models.Site

	cur, err := notificationCollection.Find(context.TODO(), bson.D{{"siteId", siteID}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var elem models.Site
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
}
