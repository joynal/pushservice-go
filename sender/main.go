package parser

import (
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

  var wg sync.WaitGroup
  maxChan := make(chan bool, maxConcurrency)

  for msg := range claim.Messages() {
    maxChan <- true
    go sendPush(msg, sess, maxChan, &wg)
  }
  wg.Wait()

  return nil
}

func main() {
  utils.LoadConfigs()
  utils.GetNewConsumerGroup("SenderGroup", os.Getenv("TOPIC_PUSH"), consumerGroupHandler{})
}