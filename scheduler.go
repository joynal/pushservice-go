package main

import (
	"context"
	"fmt"
	"log"

	"pushservice/models"

	"github.com/globalsign/mgo/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func main() {
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	defer client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	siteId, _ := primitive.ObjectIDFromHex("5c5b9ee1752a41a63268c3bb")

	var site models.Site
	idDoc := bson.D{{"_id", siteId}}
	siteCollection := client.Database("pushservice").Collection("sites")
	err = siteCollection.FindOne(context.TODO(), idDoc).Decode(&site)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", site)
}
