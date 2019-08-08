package main

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
	"sync"

	"pushservice-go/models"
)

func processPush(msg *sarama.ConsumerMessage, sess sarama.ConsumerGroupSession, maxChan chan bool, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer func(maxChan chan bool) { <-maxChan }(maxChan)

	// construct struct from byte
	var push models.Push
	err := json.Unmarshal(msg.Value, &push)
	if err != nil {
		log.Fatal(err)
	}

	// commit kafka message
	sess.MarkMessage(msg, "")
}
