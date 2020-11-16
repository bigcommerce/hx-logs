package logs

import "fmt"

// LazyMessage is used by Producer to delay formatting of messages, in case
// an event is to be discarded by a LevelFilter.
type LazyMessage struct {
	Message string
	Args    []interface{}
}

func (m *LazyMessage) String() string {
	if len(m.Args) > 0 {
		m.Message = fmt.Sprintf(m.Message, m.Args...)
		m.Args = nil
	}
	return m.Message
}
