package logs

import (
	"io"
	"os"
)

// NewPlainTextLogger creates a new Logger with the given Level for the given io.Writer,
// using DefaultPlainTextFormatter as its formatter.
func NewPlainTextLogger(level Level, writer io.Writer) *Logger {
	return NewLogger(level, NewWriterWithFormat(writer, DefaultPlainTextFormatter))
}

// NewColoredTextLogger creates a new Logger with the given Level for the given io.Writer,
// using DefaultColoredTextFormatter as its formatter.
func NewColoredTextLogger(level Level, writer io.Writer) *Logger {
	return NewLogger(level, NewWriterWithFormat(writer, DefaultColoredTextFormatter))
}

// NewFormattedLogger creates a new Logger with the given Level for the given io.Writer,
// using the given Formatter as its formatter.
func NewFormattedLogger(level Level, writer io.Writer, formatter Formatter) *Logger {
	return NewLogger(level, NewWriterWithFormat(writer, formatter))
}

// NewStdoutLogger creates a new Logger with the given Level for STDOUT, using either
// DefaultPlainTextFormatter or DefaultColoredTextFormatter as its formatter, depending
// on whether it senses a TTY.
func NewStdoutLogger(level Level) *Logger {
	return newSenseTTYLogger(level, os.Stdout)
}

// NewStderrLogger creates a new Logger with the given Level for STDERR, using either
// DefaultPlainTextFormatter or DefaultColoredTextFormatter as its formatter, depending
// on whether it senses a TTY.
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
