package models

import (
  "time"

  "github.com/mongodb/mongo-go-driver/bson/primitive"
)

// Action - Browser push action configs
type Action struct {
  Action string
  Title  string
  icon   string
  URL    string
}

// Option - Push message
type Option struct {
  Body               string
  Icon               string
  Image              string
  Badge              string
  Vibration          bool
  Renotify           bool
  RequireInteraction bool `bson:"requireInteraction"`
  Dir                string
  Tag                string
  Actions            Action
}

// Push model
type Push struct {
  ID     primitive.ObjectID `bson:"_id,omitempty"`
  SiteID primitive.ObjectID `bson:"siteId"`
  Status string

  // details
  Title   string
  Options Option

  // others
  LaunchURL  string `bson:"launchUrl"`
  Priority   string
  TimeToLive int `bson:"timeToLive"`

  // stats
  TotalSent    int `bson:"totalSent"`
  TotalDeliver int `bson:"totalDeliver"`
  TotalClick   int `bson:"totalClick"`
  TotalClose   int `bson:"totalClose"`

  CreatedAt time.Time `bson:"createdAt"`
  UpdatedAt time.Time `bson:"updatedAt"`
}

// GetPushObject Create a notification instance
func GetPushObject(SiteID primitive.ObjectID, title string) Push {
  return Push{
    SiteID:       SiteID,
    Status:       "pending",
    Title:        title,
    Options:      Option{Body: "Ignore please, load testing"},
    LaunchURL:    "https://joynal.me",
    Priority:     "high",
    TimeToLive:   259200,
    TotalSent:    0,
    TotalDeliver: 0,
    TotalClick:   0,
    TotalClose:   0,
    CreatedAt:    time.Now(),
    UpdatedAt:    time.Now(),
  }
}
