package logs_test

import (
	"bytes"
	. "github.com/bigcommerce/logs"
	. "github.com/bigcommerce/logs/testing"
	"testing"
)

func TestLogger_Log(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := NewFormattedLogger(0, buf, nil)
	logger.Debug("Hey %s", "you")
	Equals(t, "Hey you\n", buf.String())
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
