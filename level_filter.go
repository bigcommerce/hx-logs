package logs

// LevelFilter wraps a Subscriber, swallowing events that do not match LevelMask.
type LevelFilter struct {
	Subscriber
	LevelMask Level
}

func (l *LevelFilter) Log(event *Event) {
	if event.Level&l.LevelMask != 0 {
		l.Subscriber.Log(event)
	}
}

// NewLevelFilter creates a new LevelFilter that will swallow events with severity
// lower than the given Level.
func NewLevelFilter(level Level, subscriber Subscriber) *LevelFilter {
	return &LevelFilter{
		Subscriber: subscriber,
		LevelMask:  ^(level - 1),
	}
}
