package logs

import (
	"io"
	"sync"
)

// Writer is a Subscriber implementation that writes events to an io.Writer using a
// Formatter to convert events into byte slices.
type Writer struct {
	Formatter
	Writer io.Writer

	// OnWriteFail is called when the Writer is unable to write an Event to the underlying
	// io.Writer. If OnWriteFail is nil, write failures will cause a Writer to panic.
	OnWriteFail func(error)
	mutex       sync.Mutex
}

func (w *Writer) Log(event *Event) {
	w.mutex.Lock()
	_, err := w.Writer.Write(w.Formatter.Format(event))
	w.mutex.Unlock()
	if err == nil {
		return
	}
	if w.OnWriteFail != nil {
		w.OnWriteFail(err)
		return
	}
	panic(err)
}

// NewWriter makes a new Writer for the given io.Writer, using MessageOnlyFormatter as
// its formatter.
func NewWriter(writer io.Writer) *Writer {
	return NewWriterWithFormat(writer, nil)
}

// NewWriterWithFormat makes a new Writer for the given io.Writer, using the given
// Formatter.
func NewWriterWithFormat(writer io.Writer, format Formatter) *Writer {
	if format == nil {
		format = MessageOnlyFormatter
	}
	return &Writer{
		Writer:    writer,
		Formatter: format,
	}
}
