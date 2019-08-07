package utils

import (
	"log"

	"github.com/Shopify/sarama"
)

func TrackGroupErrors(client sarama.ConsumerGroup) {
	for err := range client.Errors() {
		log.Fatal(err)
	}
}
