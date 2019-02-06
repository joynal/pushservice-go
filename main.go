package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"pushservice/utils"
	"strings"
	"time"
)

type Keys struct {
	P256Dh string `json:"p256dh"`
	Auth   string `json:"auth"`
}

type PushEndPoint struct {
	Endpoint       string      `json:"endpoint"`
	ExpirationTime interface{} `json:"expirationTime"`
	Keys           Keys        `json:"keys"`
}

func main() {
	browsers := []string{"firefox", "chrome"}

	endpoints := make(map[string]string)
	endpoints["firefox"] = "https://updates.push.services.mozilla.com/wpush/v2/"
	endpoints["chrome"] = "https://fcm.googleapis.com/fcm/send/"
	p256dh := "BI4TC8K9l_rfDlnVwNDaHa8dr3_zV5VfRvbPGu7NHaTTBLvvmM_Dbpi1dwK9PLAPNFwNYR-RGZ5cA7iJyTtvXnM"
	auth := "mrLLfPc_dIlwsO521ix1bQ"
	subscriberId := "ebpSYTIyz5w:APA91bESNu5qsIA484DSFWyuDLEgMHdAJf45IwMua9lknXrhAzQCrLcN-ZWfT8GE-_kxNR6MiCq1tfPr1aKWH8bVFNm6bmtDY-xHug-B76h6IqwemtB9tnlPsTqlr9A8ZcvA3dZzlxMc"

	token, _ := utils.GenerateRandomString(56)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator

	for i := 0; i < 2; i++ {
		index := r.Intn(len(browsers))
		browser := browsers[index]
		endpoint := fmt.Sprintf("%s%s", endpoints[browser], strings.Replace(subscriberId, "APA91bESNu5qsIA484DSFWyuDLEgMHdAJf45IwMua9lknXrhAzQCrLcN", token, -1))

		fmt.Println(json.Marshal(PushEndPoint{
			Endpoint:       endpoint,
			ExpirationTime: nil,
			Keys:           Keys{p256dh, auth},
		}))
	}

	// client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	// defer client.Disconnect(context.TODO())

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Connected to MongoDB!")

	// siteID := seed.GenerateSite(client)

	// seed.GenerateNotifications(client, siteID)

	// seed.GenerateSubscribers(client, siteID)
}
