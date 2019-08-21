package main

import (
  "context"
  "encoding/json"
  "fmt"
  "github.com/SherClockHolmes/webpush-go"
  "github.com/Shopify/sarama"
  "github.com/mongodb/mongo-go-driver/bson"
  "github.com/mongodb/mongo-go-driver/mongo"
  "log"
  "pushservice-go/models"
)

func sendPush(
	msg *sarama.ConsumerMessage,
	sess sarama.ConsumerGroupSession,
	coll mongo.Collection,
	ctx context.Context) {
	// commit kafka message
	sess.MarkMessage(msg, "")

	// construct struct from byte
	var push models.PushPayload
	err := json.Unmarshal(msg.Value, &push)
	if err != nil {
		log.Fatal(err)
	}

	// Decode subscription
	s := &webpush.Subscription{}
	err = json.Unmarshal([]byte(push.PushEndpoint), s)

	if err != nil {
		fmt.Println("endpoint err:", err)
	}

	dataStr, err := json.Marshal(push.Data)
	if err != nil {
		fmt.Println("data marshal err:", err)
	}

	// Send Notification
	res, err := webpush.SendNotification(dataStr, s, &webpush.Options{
		Subscriber:      push.Options.VapidDetails.Subject,
		VAPIDPrivateKey: push.Options.VapidDetails.PrivateKey,
		VAPIDPublicKey:  push.Options.VapidDetails.PublicKey,
		TTL:             push.Options.TTL,
	})
	if err != nil {
		fmt.Println("send err:", err)
	}

	fmt.Println("res: ", res)
	
	if res != nil && res.StatusCode == 410 {
		fmt.Println("webpush error:", err)
		_, err = coll.UpdateOne(
			ctx,
			bson.M{"_id": push.SubscriberID},
			bson.M{"$set": bson.M{"status": "unSubscribed"}})

		if err != nil {
			fmt.Println("db update err:", err)
		}

	}
}
