package main

import (
	"context"
	"fmt"
	"log"
	"pushservice/models"
	"pushservice/utils"
	"time"

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

	privateKey, publicKey, _ := utils.GenerateVapidKeys()

	siteCollection := client.Database("pushservice").Collection("sites")
	notificationCollection := client.Database("pushservice").Collection("notifications")

	insertResult, err := siteCollection.InsertOne(context.TODO(), models.Site{
		VapidPublicKey:  privateKey,
		VapidPrivateKey: publicKey,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Site created successfully: ", insertResult.InsertedID)

	siteID := insertResult.InsertedID

	// create notifications
	fmt.Println("Creating notifications ------->")
	var notifications []interface{}
	for i := 0; i < 10; i++ {
		notifications = append(notifications, models.GetNotificationObject(siteID.(primitive.ObjectID), fmt.Sprintf("Load test - %d", i)))
	}

	insertManyResult, err := notificationCollection.InsertMany(context.TODO(), notifications)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}
