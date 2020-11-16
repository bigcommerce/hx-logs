package logs

import "fmt"

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
