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

const (
	maxConcurrency = 10000
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

  // new producer for sending data
  producer, err := utils.GetProducer()
  if err != nil {
    log.Fatal(err)
  }
  defer func() { _ = producer.Close() }()

  // buffered channel for concurrency control
  consumer := Consumer{
    db: *db,
    ctx: ctx,
    maxChan: make(chan bool, maxConcurrency),
    producer: producer,
  }

  // start consuming
	utils.GetConsumer("ParserGroup", os.Getenv("TOPIC_RAW_PUSH"), consumer)
}

type Consumer struct{
  maxChan chan bool
  db mongo.Database
  ctx context.Context
  producer sarama.SyncProducer
}

func (consumer Consumer) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumer Consumer) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (consumer Consumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
  for msg := range claim.Messages() {
    consumer.maxChan <- true
    processPush(msg, sess, consumer.maxChan, consumer.db, consumer.ctx, consumer.producer)
  }

  return nil
}
