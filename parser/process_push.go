package main

import (
  "context"
  "encoding/json"
  "fmt"
  "github.com/Shopify/sarama"
  "github.com/mongodb/mongo-go-driver/bson"
  "github.com/mongodb/mongo-go-driver/mongo"
  "log"
  "os"
  "pushservice-go/models"
  "sync"
  "time"
)

func processPush(
  msg *sarama.ConsumerMessage,
  sess sarama.ConsumerGroupSession,
  maxChan chan bool,
  db mongo.Database,
  ctx context.Context,
  producer sarama.SyncProducer) {
  defer func(maxChan chan bool) { <-maxChan }(maxChan)

  // commit kafka message
  sess.MarkMessage(msg, "")

  fmt.Print("record: ", string(msg.Value))

  // construct struct from byte
  var push models.RawPushPayload
  err := json.Unmarshal(msg.Value, &push)
  if err != nil {
    log.Fatal(err)
  }

  // Lets prepare subscriber query
  query := bson.M{
    "siteId": push.SiteID,
    "status": "subscribed",
  }

  // find subscribers and
  cur, err := db.Collection("subscribers").Find(ctx, query)
  if err != nil {
    log.Fatal(err)
  }
  // Close the cursor once finished
  defer func() { _ = cur.Close(ctx) }()

  // wait group for finishing all goroutines
  var wg sync.WaitGroup
  counter := 0

  // Iterate through the cursor
  for cur.Next(ctx) {
    var elem models.Subscriber
    err := cur.Decode(&elem)
    if err != nil {
      log.Fatalln("encode err:", err)
    }

    fmt.Println("subscriberId: ", elem.ID)

    wg.Add(1)
    go sendToTopic(models.PushPayload{
      SubscriberID: elem.ID,
      PushEndpoint: elem.PushEndpoint,
      Data: models.DataPayload{
        ID:                 push.ID,
        LaunchURL:          push.LaunchURL,
        Priority:           push.Priority,
        Body:               push.Options.Body,
        Icon:               push.Options.Icon,
        Image:              push.Options.Image,
        Badge:              push.Options.Badge,
        Vibration:          push.Options.Vibration,
        Renotify:           push.Options.Renotify,
        RequireInteraction: push.Options.RequireInteraction,
        Dir:                push.Options.Dir,
        Tag:                push.Options.Tag,
        Actions:            push.Options.Actions,
      },
      Options: models.PushOption{
        VapidDetails: push.VapidDetails,
        TTL:          push.TimeToLive,
      },
    }, producer, &wg)
    counter++
  }

  wg.Wait()

  if err := cur.Err(); err != nil {
    log.Fatal(err)
  }

  // update notification stats
  updateQuery := bson.M{
    "status": "done",
    "updatedAt": time.Now(),
    "totalSent": push.TotalSent + counter,
  }

  _, _ = db.Collection("pushes").UpdateOne(ctx, bson.M{"_id": push.ID}, bson.M{"$set": updateQuery})
}

func sendToTopic(data models.PushPayload, producer sarama.SyncProducer, wg *sync.WaitGroup) {
  defer wg.Done()
  jsonData, _ := json.Marshal(data)
  msg := &sarama.ProducerMessage{
    Topic: os.Getenv("TOPIC_PUSH"),
    Value: sarama.StringEncoder(string(jsonData)),
  }
  partition, offset, err := producer.SendMessage(msg)
  if err != nil {
    log.Println("send error ------------->", err)
  }

  log.Printf("sent at -------------> %s/partition - %d/offset - %d\n", os.Getenv("TOPIC_PUSH"), partition, offset)
}
