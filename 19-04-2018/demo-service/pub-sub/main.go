package main

import (
	"fmt"
	"time"
	"tools.adidas-group.com/whittjoe/engineering-community-days/19-04-2018/demo-service/pub-sub/pkg/pubsub"
)

func main() {

	srv := pubsub.NewServer()

	sub1, err := srv.SubscribeClient(make(chan string))

	if err != nil {
		panic(err)
	}

	go func() {
		for message := range sub1.Messages() {
			fmt.Println("sub 1")
			fmt.Println(message)
		}
	}()

	sub2, err := srv.SubscribeClient(make(chan string))

	if err != nil {
		panic(err)
	}

	go func() {
		for message := range sub2.Messages() {
			fmt.Println("???? sub 2 ?????")
			fmt.Println(message)

		}
	}()

	pub1 := srv.PublishClient()

	pub1.Send("I am exiting")

	srv.PublishClient().Send("Hello from Joel")

	srv.Disconnect()

	time.Sleep(10 * time.Minute)

}
