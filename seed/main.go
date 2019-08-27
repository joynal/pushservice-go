package seed

import (
  "context"
  "fmt"
  "github.com/mongodb/mongo-go-driver/mongo"
  "log"
)

func main() {
  client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
  defer client.Disconnect(context.TODO())

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Connected to MongoDB!")

  siteID := GenerateSite(client)

  GenerateNotifications(client, siteID)

  GenerateSubscribers(client, siteID)
}
