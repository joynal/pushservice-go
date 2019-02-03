package main

import (
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
		SiteID:          SiteID,
		VapidPublicKey:  privateKey,
		VapidPrivateKey: publicKey,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	})

	notification := models.GetNotificationObject(SiteID, "Load test - 1", "Ignore please, load testing", "https://joynal.me")

	notificationCollection.Insert(&notification)
}
