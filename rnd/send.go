package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "test"

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	for i := 0; i < 1000; i++ {
		err := w.WriteMessages(context.Background(),
			kafka.Message{
				Value: []byte(fmt.Sprintf("notification - %d", i)),
			},
		)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("sent --------->", i)
	}

	w.Close()
}
