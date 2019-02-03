package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// Message - Notification message
type Message struct {
	Title    string
	Message  string
	Language string
}

// Action - Browser push action configs
type Action struct {
	Action string
	Title  string
	URL    string
}

// Browser push configs
type Browser struct {
	BrowserName        string `bson:"browserName"`
	IconURL            string `bson:"iconURL"`
	ImageURL           string `bson:"imageURL"`
	Badge              string
	Vibration          bool
	IsActive           bool `bson:"isActive"`
	IsEnabledCTAButton bool `bson:"isEnabledCTAButton"`
	Actions            []Action
}

// HideRule when notification popup will be closed.
type HideRule struct {
	Type  string
	Value int
}

// Notification model
type Notification struct {
	ID           bson.ObjectId   `bson:"_id,omitempty"`
	SiteID       bson.ObjectId   `bson:"siteId"`
	Status       string          `bson:"status"`
	Message      Message         `bson:"message"`
	SendToAll    bool            `bson:"sendToAll"`
	Segments     []bson.ObjectId `bson:"segments"`
	Browsers     []Browser       `bson:"browsers"`
	HideRule     HideRule        `bson:"hideRule"`
	LaunchURL    string          `bson:"launchUrl"`
	Priority     string          `bson:"priority"`
	TTL          int             `bson:"ttl"`
	TotalSent    int             `bson:"totalSent"`
	TotalDeliver int             `bson:"totalDeliver"`
	TotalShow    int             `bson:"totalShow"`
	TotalError   int             `bson:"totalError"`
	TotalClick   int             `bson:"totalClick"`
	TotalClose   int             `bson:"totalClose"`
	SentAt       time.Time       `bson:"sentAt"`
	CreatedAt    time.Time       `bson:"createdAt"`
	UpdatedAt    time.Time       `bson:"updatedAt"`
}

// GetNotificationObject Create a notification instance
func GetNotificationObject(SiteID bson.ObjectId, title string) *Notification {
	return &Notification{
		SiteID:       SiteID,
		Status:       "pending",
		Message:      Message{title, "Ignore please, load testing", "en"},
		SendToAll:    true,
		HideRule:     HideRule{"delay", 30},
		LaunchURL:    "https://joynal.me",
		Priority:     "high",
		TTL:          259200,
		TotalSent:    0,
		TotalDeliver: 0,
		TotalShow:    0,
		TotalError:   0,
		TotalClick:   0,
		TotalClose:   0,
		SentAt:       time.Now(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
