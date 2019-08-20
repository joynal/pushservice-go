package models

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// Site model type
type Site struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Subject    string
	PublicKey  string    `bson:"publicKey"`
	PrivateKey string    `bson:"privateKey"`
	CreatedAt  time.Time `bson:"createdAt"`
	UpdatedAt  time.Time `bson:"updatedAt"`
}

type SitePayload struct {
  Subject    string
  PublicKey  string    `bson:"publicKey"`
  PrivateKey string    `bson:"privateKey"`
}
