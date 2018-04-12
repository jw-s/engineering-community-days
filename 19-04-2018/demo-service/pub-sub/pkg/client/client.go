package client

import (
	"tools.adidas-group.com/whittjoe/engineering-community-days/19-04-2018/demo-service/pub-sub/pkg/pubsub"
)

func Subscribe(server pubsub.PubSubServer) (pubsub.Subscribe, error) {
	subCh := make(chan string)
	return server.SubscribeClient(subCh)
}

func Publish(server pubsub.PubSubServer) pubsub.Publish {
	return server.PublishClient()

}
