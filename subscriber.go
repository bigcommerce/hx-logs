package logs

type Subscriber interface {
	Log(event *Event)
}
