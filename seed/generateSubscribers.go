package main

import (
	"pushservice/models"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// GenerateSubscriber for generating subscriber
func GenerateSubscriber(siteID bson.ObjectId, session mgo.Session) {
	// siteCollection := session.DB("mgo-test").C("sites")
	subscriberCollection := session.DB("mgo-test").C("subscribers")

	// create a site data
	subscriberCollection.Insert(&models.Subscriber{
		SiteID:    siteID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	// create subscribers
	// var subscribers []interface{}
	// for i := 0; i < 10; i++ {
	// }

	// subscriberCollection.Insert(subscribers...)
}
