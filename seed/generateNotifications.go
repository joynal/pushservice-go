package main

import (
	"fmt"
	"pushservice/models"
	"pushservice/utils"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func main() {
	session, err := mgo.Dial("127.0.0.1")

	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	privateKey, publicKey, _ := utils.GenerateVapidKeys()
	SiteID := bson.NewObjectId()

	siteCollection := session.DB("mgo-test").C("sites")
	notificationCollection := session.DB("mgo-test").C("notifications")

	// create a site data
	siteCollection.Insert(&models.Site{
		ID:              SiteID,
		VapidPublicKey:  privateKey,
		VapidPrivateKey: publicKey,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	})

	// create notifications
	var notifications []interface{}
	for i := 0; i < 10; i++ {
		notifications = append(notifications, models.GetNotificationObject(SiteID, fmt.Sprintf("Load test - %d", i)))
	}

	notificationCollection.Insert(notifications...)

	// TODO: create subscribers
}
