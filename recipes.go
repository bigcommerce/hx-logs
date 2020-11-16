package logs

import (
	"io"
	"os"
)

func NewPlainTextLogger(level Level, writer io.Writer) *Logger {
	return NewLogger(level, NewWriterWithFormat(writer, DefaultPlainTextFormatter))
}

func NewColoredTextLogger(level Level, writer io.Writer) *Logger {
	return NewLogger(level, NewWriterWithFormat(writer, DefaultColoredTextFormatter))
}

func NewFormattedLogger(level Level, writer io.Writer, formatter Formatter) *Logger {
	return NewLogger(level, NewWriterWithFormat(writer, formatter))
}

func NewStdoutLogger(level Level) *Logger {
	return newSenseTTYLogger(level, os.Stdout)
}

func NewStderrLogger(level Level) *Logger {
	return newSenseTTYLogger(level, os.Stderr)
}

func newSenseTTYLogger(level Level, file *os.File) *Logger {
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return NewColoredTextLogger(level, file)
	}
	return NewPlainTextLogger(level, file)
}
