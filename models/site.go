package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// Site model type
type Site struct {
	ID              bson.ObjectId `bson:"_id,omitempty"`
	SiteID          bson.ObjectId `bson:"siteId"`
	VapidPublicKey  string        `bson:"vapidPublicKey"`
	VapidPrivateKey string        `bson:"vapidPrivateKey"`
	CreatedAt       time.Time     `bson:"createdAt"`
	UpdatedAt       time.Time     `bson:"updatedAt"`
}
