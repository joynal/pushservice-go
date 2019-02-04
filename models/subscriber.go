package models

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// Subscriber model
type Subscriber struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	SiteID         primitive.ObjectID `bson:"siteId"`
	IsActive       bool               `bson:"isActive"`
	IsSubscribed   bool               `bson:"isSubscribed"`
	Token          string
	PushEndPoint   string    `bson:"pushEndPoint"`
	LastActive     time.Time `bson:"lastActive"`
	FirstSession   time.Time `bson:"firstSession"`
	SessionCount   int       `bson:"sessionCount"`
	Timezone       string
	Country        string
	Language       string
	DeviceType     string `bson:"deviceType"`
	os             string
	IP             string
	Browser        string
	BrowserVersion string               `bson:"browserVersion"`
	Segments       []primitive.ObjectID `bson:"segments"`
	CreatedAt      time.Time            `bson:"createdAt"`
	UpdatedAt      time.Time            `bson:"updatedAt"`
}
