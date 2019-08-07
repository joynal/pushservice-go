package main

import (
  "github.com/Shopify/sarama"
  "os"
  "pushservice-go/utils"
  "sync"
)

const (
  maxConcurrency = 100
)

type consumerGroupHandler struct{}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
  utils.LoadConfigs()

  var wg sync.WaitGroup
  maxChan := make(chan bool, maxConcurrency)

  for msg := range claim.Messages() {
    maxChan <- true
    go processPush(msg, sess, maxChan, &wg)
  }
  wg.Wait()

  return nil
}

func main() {
  utils.LoadConfigs()
  utils.GetConsumer("ParserGroup", os.Getenv("TOPIC_RAW_PUSH"), consumerGroupHandler{})
}
