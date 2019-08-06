package seed

import (
	"context"
	"fmt"
	"log"
	"pushservice-go/models"
	"pushservice-go/utils"
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// GenerateSite will generate site data
func GenerateSite(client *mongo.Client) primitive.ObjectID {
	privateKey, publicKey, _ := utils.GenerateVapidKeys()

	siteCollection := client.Database("pushservice").Collection("sites")

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

	return insertResult.InsertedID.(primitive.ObjectID)
}
