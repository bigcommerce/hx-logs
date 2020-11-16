package logs

type JsonFormatter struct {
	Tags TagSet
}

func NewJsonFormatter(tags TagSet) *JsonFormatter {
	return &JsonFormatter{Tags: tags}
}

func (j *JsonFormatter) Format(event *Event) []byte {
	b, err := Tags{
		{"@timestamp", event.Time},
		{"severity", event.Level.Name()},
		{"message", event.Message.String()},
	}.
		Join(j.Tags.Tags()...).
		Join(event.Tags...).
		MarshalJSON()
	if err != nil {
		b = []byte(err.Error())
	}
	return b
}
