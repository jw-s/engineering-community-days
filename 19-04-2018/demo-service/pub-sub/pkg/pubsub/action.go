package pubsub

type Publish interface {
	Send(string)
}

type Subscribe interface {
	Messages() <-chan string
}

type publish struct {
	messageCh chan string
}

func (p *publish) Send(msg string) {
	p.messageCh <- msg
}

type subscribe struct {
	subscribeCh  chan string
	disconnected bool
}

func (s *subscribe) Messages() <-chan string {
	return s.subscribeCh
}
