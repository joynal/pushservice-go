package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"pushservice/models"
	"pushservice/utils"

	"github.com/joho/godotenv"
	"github.com/mongodb/mongo-go-driver/bson"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx = context.WithValue(ctx, utils.DbURL, os.Getenv("MONGODB_URL"))
	db, err := utils.ConfigDB(ctx, "pushservice")
	if err != nil {
		log.Fatalf("database configuration failed: %v", err)
	}

	fmt.Println("Connected to MongoDB!")

	var site models.Site

	err = db.Collection("sites").FindOne(ctx, bson.D{}).Decode(&site)

	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	cur, err := db.Collection("notifications").Find(ctx, bson.D{{"siteId", site.ID}})
	if err != nil {
		log.Fatal(err)
	}
	// Close the cursor once finished
	defer cur.Close(ctx)

	// Iterate through the cursor
	for cur.Next(ctx) {
		var elem models.Notification
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(elem.Message.Title)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
