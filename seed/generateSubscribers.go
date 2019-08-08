package seed

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"pushservice-go/models"
	"pushservice-go/utils"
	"strings"
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// GenerateSubscribers will generate subscribers for site
func GenerateSubscribers(client *mongo.Client, siteID primitive.ObjectID) {
	fmt.Println("Creating subscribers ------->", time.Now())
	start := time.Now()
	subscriberCollection := client.Database("pushservice").Collection("subscribers")

	browsers := []string{"firefox", "chrome"}

	endpoints := make(map[string]string)
	endpoints["firefox"] = "https://updates.push.services.mozilla.com/wpush/v2/"
	endpoints["chrome"] = "https://fcm.googleapis.com/fcm/send/"
	p256dh := "BI4TC8K9l_rfDlnVwNDaHa8dr3_zV5VfRvbPGu7NHaTTBLvvmM_Dbpi1dwK9PLAPNFwNYR-RGZ5cA7iJyTtvXnM"
	auth := "mrLLfPc_dIlwsO521ix1bQ"
	subscriberToken := "ebpSYTIyz5w:APA91bESNu5qsIA484DSFWyuDLEgMHdAJf45IwMua9lknXrhAzQCrLcN-ZWfT8GE-_kxNR6MiCq1tfPr1aKWH8bVFNm6bmtDY-xHug-B76h6IqwemtB9tnlPsTqlr9A8ZcvA3dZzlxMc"

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	batch := 100000
	var subscribers []interface{}
	numberOfSubscribers := 1000000
	for i := 0; i < numberOfSubscribers; i++ {
		index := r.Intn(len(browsers))
		browser := browsers[index]
		token, _ := utils.GenerateRandomString(56)
		newSubscriberToken := strings.Replace(subscriberToken, "APA91bESNu5qsIA484DSFWyuDLEgMHdAJf45IwMua9lknXrhAzQCrLcN", token, -1)
		endpoint := fmt.Sprintf("%s%s", endpoints[browser], newSubscriberToken)
		pushEndpoint, _ := json.Marshal(models.PushEndPoint{
			Endpoint:       endpoint,
			ExpirationTime: nil,
			Keys:           models.Keys{p256dh, auth},
		})

		subscribers = append(subscribers, models.Subscriber{
			SiteID:       siteID,
			Subscribed:   true,
			PushEndpoint: string(pushEndpoint),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		})

		if i == batch {
			_, _ = subscriberCollection.InsertMany(context.TODO(), subscribers)
			subscribers = nil
		}
	}

	_, err := subscriberCollection.InsertMany(context.TODO(), subscribers)
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	log.Printf("Subscriber creation took only %s", elapsed)

	fmt.Println("Subscribers created successfully ------->", time.Now())
}
