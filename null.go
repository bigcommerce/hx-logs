package logs

import "io"

// NullSubscriber is an implementation of Subscriber that discards messages. It is used chiefly by NullLogger.
type NullSubscriber struct{}

func (n *NullSubscriber) Log(*Event) {}

// NullLogger is a logger that discards events. You can use it as a placeholder in apps that have logging disabled and
// want calls to Info, Warn, Error etc to no-op.
var NullLogger = NewLogger(0, new(NullSubscriber)) // 0 level avoids making a LevelFilter

// NullFormatter is an implementation of Formatter that formats every message as an empty byte slice.
type NullFormatter struct{}

func (n *NullFormatter) Format(*Event) []byte { return nil }

// NullWriter is a writer that does not write events to anywhere.
var NullWriter = NewWriterWithFormat(io.Discard, new(NullFormatter))

// NullProducer is a producer that discards all events.
var NullProducer = &Producer{Subscriber: new(NullSubscriber)}
