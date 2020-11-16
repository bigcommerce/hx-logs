package logs

import (
	"fmt"
	"time"
)

type Event struct {
	Time    time.Time
	Level   Level
	Message fmt.Stringer
	Tags    Tags
}
