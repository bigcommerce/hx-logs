package logs

// Buffer uses a goroutine to consume events from a buffered channel, ensuring
// publication doesn't block on writing (at least until the buffer is full). A
// zero Buffer will not work. Use NewBuffer instead.
type Buffer struct {
	Subscriber
	buf chan *Event
}

// NewBuffer wraps the given Subscriber in a buffer with the given size.
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

// Close closes the buffer's underlying channel. Calls made to Log after Close
// is called will panic.
func (b *Buffer) Close() {
	close(b.buf)
}

func (b *Buffer) consume() {
	for event := range b.buf {
		b.Subscriber.Log(event)
	}
}
