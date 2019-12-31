//ref: https://github.com/GoogleCloudPlatform/golang-samples/tree/master/pubsub/subscriptions

package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"flag"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	//get arguments
	projectId := flag.String("project_id", "hoge", "GCP project name")
	mode := flag.String("mode", "pull", "PubSub handling mode")
	//FYI: Case of "subscription", it may not need "topic name".
	//topic := flag.String("topic", "", "PubSub topic name")
	subscriptionName := flag.String("sub_name", "", "PubSub subscription name")
	flag.Parse()

	if *mode != "pull" {
		fmt.Printf("mode invalid. mode:%s\n", *mode)
		return
	}

	client, err := pubsub.NewClient(ctx, *projectId)
	if err != nil {
		fmt.Println(err)
		return
	}

	subscription := client.Subscription(*subscriptionName)

	//Receive messages for 10 seconds.
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Create a channel to handle messages to as they come in.
	cm := make(chan *pubsub.Message)
	// Handle individual messages in a goroutine.
	go func() {
		for {
			select {
			case msg := <-cm:
				fmt.Printf("Got message :%q\n", string(msg.Data))
				msg.Ack()
			case <-ctx.Done():
				return
			}
		}
	}()

	// Receive blocks until the context is cancelled or an error occurs.
	err = subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		cm <- msg
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	close(cm)

	return
}
