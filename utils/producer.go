package utils

import (
  "crypto/tls"
  "log"
  "os"
  "strconv"
  "strings"
  "time"

  "github.com/Shopify/sarama"
)

func GetProducer() (sarama.SyncProducer, error) {
  brokerList := strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
  config := sarama.NewConfig()
  config.Version = sarama.V2_3_0_0
  config.Producer.RequiredAcks = sarama.WaitForAll
  config.Producer.Retry.Max, _ = strconv.Atoi(os.Getenv("PRODUCER_RETRY_MAX"))
  config.Producer.Return.Successes, _ = strconv.ParseBool(os.Getenv("PRODUCER_RETRY_RETURN_SUCCESSES"))

  kafkaSecurity, err := strconv.ParseBool(os.Getenv("KAFKA_SECURITY_ENABLED"))
  if err != nil {
    log.Fatal(err)
  }

  if kafkaSecurity == true {
    config.ClientID = "push-service"
    config.Net.KeepAlive = 1 * time.Hour
    config.Net.TLS.Enable = true
    config.Net.SASL.Enable = true
    config.Net.SASL.Handshake = true
    config.Net.SASL.User = os.Getenv("KAFKA_USERNAME")
    config.Net.SASL.Password = os.Getenv("KAFKA_PASSWORD")
    config.Net.TLS.Config = &tls.Config{InsecureSkipVerify: true}
    config.Net.SASL.Mechanism = sarama.SASLMechanism(sarama.SASLTypeSCRAMSHA256)
    config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA256} }
  }

  return sarama.NewSyncProducer(brokerList, config)
}
