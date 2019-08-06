package main

import (
	"context"
	"fmt"
	"log"
	"pushservice-go/seed"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func main() {
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	defer client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	siteID := seed.GenerateSite(client)

	seed.GenerateNotifications(client, siteID)

	seed.GenerateSubscribers(client, siteID)
}
