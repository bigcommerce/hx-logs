package logs

import (
	"strconv"
	"strings"
)

// Level represents log message severity. Levels are preset as the seven constants Verbose,
// Trace, Debug, Info, Warn, Error, and Fatal.
type Level uint8

const (
	Verbose Level = 1 << iota
	Trace
	Debug
	Info
	Warn
	Error
	Fatal
	Silent
)

// Color represents an ANSI color code that can be used to enhance log messages meant for
// display on a TTY.
type Color int

const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

const (
	NoColor Color = iota
	Bold
)

// Esc returns an escape sequence for the receiving Color, plus any additional modifiers.
// You can, for example, pass Bold as an argument to get a brighter version of a color.
func (c Color) Esc(others ...Color) (seq []byte) {
	seq = strconv.AppendInt([]byte{27, '['}, int64(c), 10)
	for _, other := range others {
		seq = strconv.AppendInt(append(seq, ';'), int64(other), 10)
	}
	seq = append(seq, 'm')
	return
}

// LevelInfo represents the default properties of a Level.
type LevelInfo struct {
	Name         string
	Abbreviation string
	Color        Color
}

// Levels contains default properties of the set levels, which can be used by formatters.
var Levels = map[Level]LevelInfo{
	Verbose: {"verbose", "VRBSE", Blue},
	Trace:   {"trace", "TRACE", Green},
	Debug:   {"debug", "DEBUG", Cyan},
	Info:    {"info", "INFO ", NoColor},
	Warn:    {"warn", "WARN ", Yellow},
	Error:   {"error", "ERROR", Red},
	Fatal:   {"fatal", "FATAL", Magenta},
	Silent:  {"silent", "SIL", White},
}

// Name returns the lowercase name of a Level.
func (l Level) Name() string {
	return Levels[l].Name
}

// Abbreviation returns a five-character uppercase string representing the receiving Level.
func (l Level) Abbreviation() string {
	return Levels[l].Abbreviation
}

// Color returns the Color of the receiving Level, based on the Levels map.
func (l Level) Color() Color {
	return Levels[l].Color
}

// LevelByName returns a Level if one case-insensitively matches the given name. If none
// match, InvalidLevelName is returned.
func LevelByName(name string) (Level, error) {
	name = strings.ToLower(name)
	for level, info := range Levels {
		if info.Name == name {
			return level, nil
		}
	}
	return 0, InvalidLevelName
}
