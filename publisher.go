package logs

// Publisher relays each Event it receives to as many Subscriber implementations as are in its slice.
type Publisher []Subscriber

func (p Publisher) Log(event *Event) {
	for _, subscriber := range p {
		subscriber.Log(event)
	}
}

// NewPublisher creates a new Publisher that relays events to the given Subscriber slice.
func NewPublisher(subscribers ...Subscriber) Publisher {
	return subscribers
}

// Add appends an additional Subscriber to the receiving Publisher.
func (p *Publisher) Add(subscriber Subscriber) {
	*p = append(*p, subscriber)
}
