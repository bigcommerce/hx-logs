package logs

import "time"

type Producer struct {
	Subscriber
}

func (p *Producer) Verbose(message string, a ...interface{}) { p.Log(Verbose, message, a) }
func (p *Producer) Trace(message string, a ...interface{})   { p.Log(Trace, message, a) }
func (p *Producer) Debug(message string, a ...interface{})   { p.Log(Debug, message, a) }
func (p *Producer) Info(message string, a ...interface{})    { p.Log(Info, message, a) }
func (p *Producer) Warn(message string, a ...interface{})    { p.Log(Warn, message, a) }
func (p *Producer) Error(message string, a ...interface{})   { p.Log(Error, message, a) }
func (p *Producer) Fatal(message string, a ...interface{})   { p.Log(Fatal, message, a) }

func (p *Producer) Log(level Level, message string, a []interface{}) {
	p.Subscriber.Log(&Event{
		Time:  time.Now(),
		Level: level,
		Message: &LazyMessage{
			Message: message,
			Args:    a,
		},
	})
}
