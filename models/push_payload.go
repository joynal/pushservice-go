package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type DataPayload struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	LaunchURL          string             `bson:"launchUrl"`
	Priority           string
	Body               string
	Icon               string
	Image              string
	Badge              string
	Vibration          bool
	Renotify           bool
	RequireInteraction bool `bson:"requireInteraction"`
	Dir                string
	Tag                string
  Actions            []Action
}

type PushOption struct {
	VapidDetails VapidDetails `bson:"vapidDetails"`
	TTL          int  `bson:"timeToLive"`
}

// Subscriber push payload
type PushPayload struct {
	SubscriberID primitive.ObjectID `bson:"subscriberId"`
	PushEndpoint string             `bson:"pushEndPoint"`
	Data         DataPayload
	Options      PushOption
}
