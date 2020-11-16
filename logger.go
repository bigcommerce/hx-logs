package logs

import "os"

type Logger struct {
	*Producer
	Tags       Tags
	subscriber Subscriber
	level      Level
}

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

func (l *Logger) SetLevel(level Level) {
	l.level = level
	if level == 0 {
		l.Producer.Subscriber = l
	} else {
		l.Producer.Subscriber = NewLevelFilter(level, l)
	}
}

func (l *Logger) Extend(tags TagSet) (extended *Logger) {
	extended = NewLogger(l.level, l.subscriber)
	extended.Tags = l.Tags.Join(tags.Tags()...)
	return
}

var DefaultLogger = NewLogger(0, NewWriterWithFormat(os.Stdout, DefaultPlainTextFormatter))
