package logs

// Logger is a full-featured logging utility.
type Logger struct {
	*Producer
	Tags       Tags
	subscriber Subscriber
	level      Level
}

// NewLogger creates a new Logger with the given Level, which publishes synchronously
// to the given Subscriber.
func NewLogger(level Level, subscriber Subscriber) (logger *Logger) {
	logger = &Logger{
		Producer:   new(Producer),
		subscriber: subscriber,
	}
	logger.SetLevel(level)
	return
}

func (l *Logger) Log(event *Event) {
	event.Tags = l.Tags
	l.subscriber.Log(event)
}

// SetLevel changes the Level of the receiving Logger. The function is not concurrency-safe,
// so care should be taken not to race any log calls.
func (l *Logger) SetLevel(level Level) {
	l.level = level
	if level == 0 {
		l.Producer.Subscriber = l
	} else {
		l.Producer.Subscriber = NewLevelFilter(level, l)
	}
}

// Extend returns a copy of the receiving Logger that has the given TagSet added to its
// Tags list. The TagSet is evaluated at call time, so Loggers depending on changing
// TagSet implementations should be discarded after use.
func (l *Logger) Extend(tags TagSet) (extended *Logger) {
	extended = NewLogger(l.level, l.subscriber)
	extended.Tags = l.Tags.Join(tags.Tags()...)
	return
}

// DefaultLogger is an all-levels, TTY-sensing logger that writes to STDOUT.
var DefaultLogger = NewStdoutLogger(0)
