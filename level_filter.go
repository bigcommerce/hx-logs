package logs

type LevelFilter struct {
	Subscriber
	LevelMask Level
}

func (l *LevelFilter) Log(event *Event) {
	if event.Level&l.LevelMask != 0 {
		l.Subscriber.Log(event)
	}
}

func NewLevelFilter(level Level, subscriber Subscriber) *LevelFilter {
	return &LevelFilter{
		Subscriber: subscriber,
		LevelMask:  ^(level - 1),
	}
}
