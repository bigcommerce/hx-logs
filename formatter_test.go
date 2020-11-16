package logs_test

import (
	. "github.com/hx/logs"
	. "github.com/hx/logs/testing"
	"testing"
	"time"
)

var FormattingTestEvent = &Event{
	Time:  time.Unix(1234567890, 123e6),
	Level: Debug,
	Message: &LazyMessage{
		Message: "Let's %s",
		Args:    []interface{}{"party"},
	},
}

// Note these depend on DefaultTimeFormatter as well.

func TestDefaultColoredTextFormatter(t *testing.T) {
	actual := DefaultColoredTextFormatter.Format(FormattingTestEvent)
	Equals(t, "13 Feb 2009 23:31:30.123 UTC \033[36;1mDEBUG\033[0m Let's party\n", string(actual))
}

func TestDefaultPlainTextFormatter(t *testing.T) {
	actual := DefaultPlainTextFormatter.Format(FormattingTestEvent)
	Equals(t, "13 Feb 2009 23:31:30.123 UTC DEBUG Let's party\n", string(actual))
}
