package logs_test

import (
	"bytes"
	. "github.com/hx/logs"
	. "github.com/hx/logs/testing"
	"testing"
)

var LevelOnlyFormatter Formatter = FormatterFunc(func(event *Event) []byte {
	return append([]byte(event.Level.Name()), '\n')
})

func TestLevelFilter_Log(t *testing.T) {
	var (
		buf    = new(bytes.Buffer)
		filter = &LevelFilter{
			Subscriber: NewWriterWithFormat(buf, LevelOnlyFormatter),
			LevelMask:  Fatal | Info,
		}
		prod = &Producer{Subscriber: filter}
	)
	prod.Debug("")
	prod.Info("")
	prod.Warn("")
	prod.Fatal("")
	prod.Error("")
	Equals(t, "info\nfatal\n", buf.String())
}

func TestNewLevelFilter(t *testing.T) {
	var (
		buf    = new(bytes.Buffer)
		filter = NewLevelFilter(Warn, NewWriterWithFormat(buf, LevelOnlyFormatter))
		prod   = &Producer{Subscriber: filter}
	)
	prod.Debug("")
	prod.Info("")
	prod.Warn("")
	prod.Error("")
	prod.Fatal("")
	Equals(t, "warn\nerror\nfatal\n", buf.String())
}
