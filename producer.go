package logs

import "time"

// Producer provides the per-level functions of Logger, but can also be wrapped around any
// Subscriber when composing custom loggers.
type Producer struct {
	Subscriber
	Clock func() time.Time // Function used to timestamp events created by the receiver
}

func (p *Producer) Verbose(message string, a ...interface{}) { p.Log(Verbose, message, a) }
func (p *Producer) Trace(message string, a ...interface{})   { p.Log(Trace, message, a) }
func (p *Producer) Debug(message string, a ...interface{})   { p.Log(Debug, message, a) }
func (p *Producer) Info(message string, a ...interface{})    { p.Log(Info, message, a) }
func (p *Producer) Warn(message string, a ...interface{})    { p.Log(Warn, message, a) }
func (p *Producer) Error(message string, a ...interface{})   { p.Log(Error, message, a) }
func (p *Producer) Fatal(message string, a ...interface{})   { p.Log(Fatal, message, a) }

func (p *Producer) Log(level Level, message string, a []interface{}) {
	event := Event{
		Level: level,
		Message: &LazyMessage{
			Message: message,
			Args:    a,
		},
	}
	if p.Clock == nil {
		event.Time = time.Now()
	} else {
		event.Time = p.Clock()
	}
	p.Subscriber.Log(&event)
}
