package tools

// A data broker that can be published to, and eny entity that has subscribed will get a copy of the message.
// https://stackoverflow.com/questions/36417199/how-to-broadcast-message-using-channel
// https://stackoverflow.com/a/49877632
type Broker[T any] struct {
	stopCh    chan struct{}
	publishCh chan T
	subCh     chan chan T
	unsubCh   chan chan T
}

// Creates a new Broker and start it running
func NewBroker[T any]() *Broker[T] {
	ret := &Broker[T]{
		stopCh:    make(chan struct{}),
		publishCh: make(chan T, 100),
		subCh:     make(chan chan T, 100),
		unsubCh:   make(chan chan T, 100),
	}

	go ret.start()
	return ret
}

// Should be started as a go routine
func (b *Broker[T]) start() {
	subs := map[chan T]struct{}{}
	for {
		select {
		case <-b.stopCh:
			for msgCh := range subs {
				// drain the messages
				close(msgCh)
				for range msgCh {
					// do nothing
				}
			}
			return
		case msgCh := <-b.subCh:
			subs[msgCh] = struct{}{}
		case msgCh := <-b.unsubCh:
			// drain the messages
			close(msgCh)
			for range msgCh {
				// do nothing
			}
			delete(subs, msgCh)
		case msg := <-b.publishCh:
			for msgCh := range subs {
				// msgCh is buffered, use non-blocking send to protect the broker:
				select {
				case msgCh <- msg:
				default:
				}
			}
		}
	}
}

func (b *Broker[T]) Stop() {
	close(b.stopCh)
}

func (b *Broker[T]) Subscribe() chan T {
	msgCh := make(chan T, 100)
	b.subCh <- msgCh
	return msgCh
}

func (b *Broker[T]) Unsubscribe(msgCh chan T) {
	b.unsubCh <- msgCh
}

func (b *Broker[T]) Publish(msg T) {
	b.publishCh <- msg
}
