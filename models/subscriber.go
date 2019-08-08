package models

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// PUsh keys
type Keys struct {
	P256Dh string `json:"p256dh"`
	Auth   string `json:"auth"`
}

// push endpoint object
type PushEndPoint struct {
	Endpoint       string      `json:"endpoint"`
	ExpirationTime interface{} `json:"expirationTime"`
	Keys           Keys        `json:"keys"`
}

// Subscriber model
type Subscriber struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	SiteID       primitive.ObjectID `bson:"siteId"`
	Subscribed   bool
	PushEndpoint string    `bson:"pushEndPoint"`
	CreatedAt    time.Time `bson:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt"`
}
