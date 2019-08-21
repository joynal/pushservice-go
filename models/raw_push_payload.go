package models

import (
  "github.com/mongodb/mongo-go-driver/bson/primitive"
)

// Push payload raw
type RawPushPayload struct {
  ID     primitive.ObjectID `bson:"_id" json:"_id"`
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
  VapidDetails   VapidDetails `bson:"vapidDetails"`
}

