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
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	SiteID         primitive.ObjectID `bson:"siteId"`
	IsActive       bool               `bson:"isActive"`
	IsSubscribed   bool               `bson:"isSubscribed"`
	Token          string
	PushEndpoint   string    `bson:"pushEndPoint"`
	LastActive     time.Time `bson:"lastActive"`
	FirstSession   time.Time `bson:"firstSession"`
	SessionCount   int       `bson:"sessionCount"`
	Timezone       string
	Country        string
	Language       string
	DeviceType     string `bson:"deviceType"`
	Os             string
	IP             string
	Browser        string
	BrowserVersion string               `bson:"browserVersion"`
	Segments       []primitive.ObjectID `bson:"segments"`
	CreatedAt      time.Time            `bson:"createdAt"`
	UpdatedAt      time.Time            `bson:"updatedAt"`
}
