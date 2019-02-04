package models

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// Site model type
type Site struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	VapidPublicKey  string             `bson:"vapidPublicKey"`
	VapidPrivateKey string             `bson:"vapidPrivateKey"`
	CreatedAt       time.Time          `bson:"createdAt"`
	UpdatedAt       time.Time          `bson:"updatedAt"`
}
