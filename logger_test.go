package logs_test

import (
	"bytes"
	. "github.com/hx/logs"
	. "github.com/hx/logs/testing"
	"testing"
	"time"
)

func TestLogger_Log(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := NewPlainTextLogger(0, buf)
	logger.Clock = func() time.Time { return time.Unix(0, 0) }
	logger.Debug("Hey %s", "you")
	Equals(t, "01 Jan 1970 00:00:00.000 UTC DEBUG Hey you\n", buf.String())
}

var TagsOnlyFormatter = FormatterFunc(func(event *Event) []byte {
	return append(event.Tags.QueryEncode(), '\n')
})

func TestLogger_Extend(t *testing.T) {
	var (
		buf    = new(bytes.Buffer)
		logger = NewFormattedLogger(0, buf, TagsOnlyFormatter)
	)
	logger.Tags = Tags{{"foo", 123}}
	logger.Extend(Tags{{"bar", 456}}).Info("")
	logger.Debug("")
	Equals(t, "foo=123&bar=456\nfoo=123\n", buf.String())
}
