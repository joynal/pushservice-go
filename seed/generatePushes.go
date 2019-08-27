package seed

import (
  "context"
  "fmt"
  "log"
  "pushservice-go/models"
  "time"

  "github.com/mongodb/mongo-go-driver/bson/primitive"
  "github.com/mongodb/mongo-go-driver/mongo"
)

// GenerateNotifications will generate notifications for site
func GenerateNotifications(client *mongo.Client, siteID primitive.ObjectID) {
  notificationCollection := client.Database("pushservice").Collection("pushes")

  fmt.Println("Creating notifications ------->", time.Now())

  var notifications []interface{}
  for i := 0; i < 10; i++ {
    notifications = append(notifications, models.GetPushObject(siteID, fmt.Sprintf("Load test - %d", i)))
  }

  insertManyResult, err := notificationCollection.InsertMany(context.TODO(), notifications)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Notifications ids: ", insertManyResult.InsertedIDs)
}
