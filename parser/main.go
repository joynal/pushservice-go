package parser

import (
  "context"
  "log"
  "os"
  "strconv"
  "strings"
  "sync"

  "go-kafka-example/utils"

  "github.com/Shopify/sarama"
  "github.com/joho/godotenv"

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
  err := godotenv.Load()
  if err != nil {
    log.Fatal(err)
  }

  brokerList := strings.Split(os.Getenv("KAFKA_SERVER_URL"), ",")
  groupId := "ParserGroup"

  config := sarama.NewConfig()
  config.Version = sarama.V2_1_0_0
  config.Consumer.Return.Errors, _ = strconv.ParseBool(os.Getenv("CONSUMER_RETRY_RETURN_SUCCESSES"))
  group, err := sarama.NewConsumerGroup(brokerList, groupId, config)
  if err != nil {
    log.Fatal(err)
  }
  defer group.Close()

  // Track errors
  go func() {
    for err := range group.Errors() {
      log.Fatal(err)
    }
  }()

  ctx := context.Background()
  for {
    err = group.Consume(ctx, []string{os.Getenv("TOPIC_RAW_PUSH")}, consumerGroupHandler{})
    if err != nil {
      panic(err)
    }
  }
}
