package logs

import (
	"fmt"
)

type Formatter interface {
	Format(event *Event) []byte
}

type FormatterFunc func(event *Event) []byte

func (f FormatterFunc) Format(event *Event) (formatted []byte) {
	if f != nil {
		formatted = f(event)
	}
	return
}

var DefaultPlainTextFormatter Formatter = FormatterFunc(func(event *Event) []byte {
	return append(DefaultTimeFormatter.Format(event), []byte(fmt.Sprintf(
		" %s %s\n",
		event.Level.Abbreviation(),
		event.Message,
	))...)
})

var DefaultColoredTextFormatter Formatter = FormatterFunc(func(event *Event) []byte {
	return append(DefaultTimeFormatter.Format(event), []byte(fmt.Sprintf(
		" %s%s%s %s\n",
		event.Level.Color().Esc(Bold),
		event.Level.Abbreviation(),
		NoColor.Esc(),
		event.Message,
	))...)
})

var MessageOnlyFormatter Formatter = FormatterFunc(func(event *Event) []byte {
	return append([]byte(event.Message.String()), '\n')
})
