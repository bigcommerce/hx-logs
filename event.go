package logs

import (
	"fmt"
	"time"
)

// Event represents a single log message.
type Event struct {
	Time    time.Time
	Level   Level
	Message fmt.Stringer
	Tags    Tags
}
