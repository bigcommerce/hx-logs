package logs

import "io"

type Writer struct {
	Formatter
	Writer      io.Writer
	OnWriteFail func(error)
}

func (w *Writer) Log(event *Event) {
	_, err := w.Writer.Write(w.Formatter.Format(event))
	if err == nil {
		return
	}
	if w.OnWriteFail != nil {
		w.OnWriteFail(err)
		return
	}
	panic(err)
}

func NewWriter(writer io.Writer) *Writer {
	return NewWriterWithFormat(writer, nil)
}

func NewWriterWithFormat(writer io.Writer, format Formatter) *Writer {
	if format == nil {
		format = MessageOnlyFormatter
	}
	return &Writer{
		Writer:    writer,
		Formatter: format,
	}
}
