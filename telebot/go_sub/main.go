package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "my-project-1668668140794")
	if err != nil {
		// TODO: Handle error.
	}
	sub := client.Subscription("sub_one")
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		fmt.Printf("%s\n", string(m.Data))
		m.Ack()
	})
	if err != context.Canceled {
		// TODO: Handle error.
	}
}
