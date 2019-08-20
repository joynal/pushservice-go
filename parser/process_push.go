package main

import (
  "context"
  "encoding/json"
  "github.com/Shopify/sarama"
  "github.com/mongodb/mongo-go-driver/mongo"
  "log"
  "pushservice-go/models"
)

func processPush(
  msg *sarama.ConsumerMessage,
  sess sarama.ConsumerGroupSession,
  maxChan chan bool,
  coll mongo.Collection,
  ctx context.Context) {
	defer func(maxChan chan bool) { <-maxChan }(maxChan)

  // commit kafka message
  sess.MarkMessage(msg, "")

	// construct struct from byte
	var push models.RawPushPayload
	err := json.Unmarshal(msg.Value, &push)
	if err != nil {
		log.Fatal(err)
	}
}
