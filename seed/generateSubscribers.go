package seed

import (
	"context"
	"fmt"
	"log"
	"pushservice/models"
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// GenerateSubscribers will generate subscribers for site
func GenerateSubscribers(client *mongo.Client, siteID primitive.ObjectID) {
	fmt.Println("Creating subscribers ------->", time.Now())
	subscriberCollection := client.Database("pushservice").Collection("subscribers")

	var subscribers []interface{}
	numberOfSubscribers := 50
	for i := 0; i < numberOfSubscribers; i++ {
		subscribers = append(subscribers, models.Subscriber{
			SiteID:         siteID,
			IsActive:       true,
			IsSubscribed:   true,
			Token:          "sdlkjro32wfesclmlsldkf3kefdc",
			PushEndPoint:   "",
			LastActive:     time.Now(),
			FirstSession:   time.Now(),
			SessionCount:   9,
			Timezone:       "",
			Country:        "",
			Language:       "",
			DeviceType:     "",
			Os:             "",
			IP:             "",
			Browser:        "",
			BrowserVersion: "",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		})
	}

	_, err := subscriberCollection.InsertMany(context.TODO(), subscribers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Subscribers created successfully ------->", time.Now())
}
