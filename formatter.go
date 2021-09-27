package logs

import (
	"fmt"
)

// Formatter can format a given Event for writing to a log device. Formatters should
// normally include a trailing newline.
type Formatter interface {
	Format(event *Event) []byte
}

// FormatterFunc creates a Formatter using the given function.
type FormatterFunc func(event *Event) []byte

func (f FormatterFunc) Format(event *Event) (formatted []byte) {
	if f != nil {
		formatted = f(event)
	}
	return
}

// DefaultPlainTextFormatter formats events as the time (using DefaultTimeFormatter),
// the Level abbreviation, and the message. Tags are not included.
var DefaultPlainTextFormatter Formatter = FormatterFunc(func(event *Event) []byte {
	return append(DefaultTimeFormatter.Format(event), []byte(fmt.Sprintf(
		" %s %s\n",
		event.Level.Abbreviation(),
		event.Message,
	))...)
})

// DefaultColoredTextFormatter is the same as DefaultPlainTextFormatter, but adds
// color to the Level component to help messages of different severities stand out.
var DefaultColoredTextFormatter Formatter = FormatterFunc(func(event *Event) []byte {
	return append(DefaultTimeFormatter.Format(event), []byte(fmt.Sprintf(
		" %s %s\n",
		event.Level.Color().Bold().Wrap(event.Level.Abbreviation()),
		event.Message,
	))...)
})

// MessageOnlyFormatter is a simple formatter that returns an event's message followed
// by a newline.
var MessageOnlyFormatter Formatter = FormatterFunc(func(event *Event) []byte {
	return append([]byte(event.Message.String()), '\n')
})
