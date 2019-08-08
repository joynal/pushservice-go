package utils

import (
	"context"
	"crypto/tls"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Shopify/sarama"
)

func GetConsumer(groupId string, topic string, handler sarama.ConsumerGroupHandler) {
	LoadConfigs()

	brokerList := strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	config := sarama.NewConfig()
	config.Version = sarama.V2_3_0_0
	config.Consumer.Return.Errors, _ = strconv.ParseBool(os.Getenv("CONSUMER_RETRY_RETURN_SUCCESSES"))

	kafkaSecurity, err := strconv.ParseBool(os.Getenv("KAFKA_SECURITY_ENABLED"))
	if err != nil {
		log.Fatal(err)
	}

	if kafkaSecurity == true {
		config.Net.TLS.Enable = true
		config.Net.SASL.Enable = true
		config.Net.SASL.Handshake = true
		config.Net.SASL.User = os.Getenv("KAFKA_USERNAME")
		config.Net.SASL.Password = os.Getenv("KAFKA_PASSWORD")
		config.Net.TLS.Config = &tls.Config{InsecureSkipVerify: true}
		config.Net.SASL.Mechanism = sarama.SASLMechanism(sarama.SASLTypeSCRAMSHA256)
		config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA256} }
	}

	client, err := sarama.NewConsumerGroup(brokerList, groupId, config)
	if err != nil {
		log.Fatal(err)
	}

	defer func() { _ = client.Close() }()

	// Track errors
	go TrackGroupErrors(client)

	ctx := context.Background()
	for {
		err = client.Consume(ctx, []string{topic}, handler)
		if err != nil {
			panic(err)
		}
	}
}
