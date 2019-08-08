package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"pushservice-go/utils"

	"github.com/Shopify/sarama"
)

const (
	maxConcurrency = 1000
)

type consumerGroupHandler struct{}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
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

	var wg sync.WaitGroup
	maxChan := make(chan bool, maxConcurrency)

	for msg := range claim.Messages() {
		maxChan <- true
		go sendPush(msg, sess, maxChan, &wg, *coll, ctx)
	}
	wg.Wait()

	return nil
}

func init() {
	sarama.Logger = log.New(os.Stdout, "[Sarama] ", log.LstdFlags)
}

func main() {
	utils.LoadConfigs()
	utils.GetConsumer("SenderGroup", os.Getenv("TOPIC_PUSH"), consumerGroupHandler{})
}
