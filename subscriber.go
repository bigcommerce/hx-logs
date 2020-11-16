package logs

// Subscriber represents anything that can receive Log events.
type Subscriber interface {
	Log(event *Event)
}
