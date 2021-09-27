package logs

import (
	"github.com/hx/golib/ansi"
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

// LevelInfo represents the default properties of a Level.
type LevelInfo struct {
	Name         string
	Abbreviation string
	Color        ansi.EscapeSequence
}

// Levels contains default properties of the set levels, which can be used by formatters.
var Levels = map[Level]LevelInfo{
	Verbose: {"verbose", "VRBSE", ansi.Blue},
	Trace:   {"trace", "TRACE", ansi.Green},
	Debug:   {"debug", "DEBUG", ansi.Cyan},
	Info:    {"info", "INFO ", 0},
	Warn:    {"warn", "WARN ", ansi.Yellow},
	Error:   {"error", "ERROR", ansi.Red},
	Fatal:   {"fatal", "FATAL", ansi.Magenta},
	Silent:  {"silent", "SILNT", ansi.White},
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
func (l Level) Color() ansi.EscapeSequence {
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
