package pubsub

import (
	"errors"
	"fmt"
)

type PubSubServer interface {
	Disconnect()
	SubscribeClient(chan string) (Subscribe, error)
	PublishClient() Publish
}

type server struct {
	publishCh   chan string
	subscribers []*subscribe
	doneCh      chan struct{}
}

func NewServer() PubSubServer {
	srv := &server{
		publishCh: make(chan string),
		doneCh:    make(chan struct{}),
	}
	go srv.handleMessages()
	return srv
}

func (s *server) SubscribeClient(subCh chan string) (Subscribe, error) {

	if subCh == nil {
		return nil, errors.New("channel can't be nil")
	}

	subscribe := &subscribe{
		subscribeCh: subCh,
	}
	s.subscribers = append(s.subscribers, subscribe)

	return subscribe, nil
}

func (s *server) PublishClient() Publish {
	return &publish{
		messageCh: s.publishCh,
	}
}

func (s *server) handleMessages() {
	for {
		select {
		case msg := <-s.publishCh:
			for _, subscriber := range s.subscribers {

				select {
				case subscriber.subscribeCh <- msg:
				}
			}
		case <-s.doneCh:
			return
		}
	}
}

func (s *server) Disconnect() {
	fmt.Println("exiting")
	close(s.doneCh)
}
