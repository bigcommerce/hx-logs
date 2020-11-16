package logs

type Buffer struct {
	Subscriber
	buf chan *Event
}

func NewBuffer(size int, subscriber Subscriber) (b *Buffer) {
	b = &Buffer{
		Subscriber: subscriber,
		buf:        make(chan *Event, size),
	}
	go b.consume()
	return
}

func (b *Buffer) Log(event *Event) {
	b.buf <- event
}

func (b *Buffer) Close() {
	close(b.buf)
}

func (b *Buffer) consume() {
	for event := range b.buf {
		b.Subscriber.Log(event)
	}
}
