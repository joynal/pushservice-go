package main

import (
  "context"
  "fmt"
  "github.com/Shopify/sarama"
  "github.com/mongodb/mongo-go-driver/mongo"
  "log"
  "os"
  "pushservice-go/utils"
)

func init() {
  sarama.Logger = log.New(os.Stdout, "[Sarama] ", log.LstdFlags)
}

func main() {
  utils.LoadConfigs()

  dbUrl := os.Getenv("MONGODB_URL")
  dbName := os.Getenv("DB_NAME")

  // Db connection stuff
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()

  ctx = context.WithValue(ctx, utils.DbURL, dbUrl)
  db, err := utils.ConfigDB(ctx, dbName)
  if err != nil {
    log.Fatalf("database configuration failed: %v", err)
  }

  fmt.Println("Connected to MongoDB!")
  coll := db.Collection("subscribers")

  consumer := Consumer{
    coll: *coll,
    ctx:  ctx,
  }

  // consuming
  utils.GetConsumer("SenderGroup", os.Getenv("TOPIC_PUSH"), consumer)
}

type Consumer struct {
  coll mongo.Collection
  ctx  context.Context
}

func (consumer Consumer) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumer Consumer) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (consumer Consumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
  for msg := range claim.Messages() {
    sendPush(msg, sess, consumer.coll, consumer.ctx)
  }

  return nil
}
