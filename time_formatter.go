package logs

import "time"

// DefaultTimeFormatter is the default TimeFormatter used by various packaged default
// formatters and Logger recipes.
var DefaultTimeFormatter = &TimeFormatter{
	TimeFormat: "02 Jan 2006 15:04:05.000 MST",
	Location:   time.UTC,
}

// TimeFormatter is a Formatter implementation that formats the Time of an Event using
// the time package's formatter.
type TimeFormatter struct {
	TimeFormat string
	Location   *time.Location
}

// NewTimeFormatter creates a new TimeFormatter with the given format and time.Location.
func NewTimeFormatter(format string, location *time.Location) *TimeFormatter {
	return &TimeFormatter{
		Location:   location,
		TimeFormat: format,
	}
}

func (t *TimeFormatter) Format(event *Event) []byte {
	var (
		format = t.TimeFormat
		loc    = t.Location
	)
	if format == "" {
		format = DefaultTimeFormatter.TimeFormat
	}
	if loc == nil {
		loc = DefaultTimeFormatter.Location
	}
	return []byte(event.Time.In(loc).Format(format))
}
