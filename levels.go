package logs

import (
	"strconv"
	"strings"
)

type Level uint8

const (
	Verbose Level = 1 << iota
	Trace
	Debug
	Info
	Warn
	Error
	Fatal
)

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

func (c Color) Esc(others ...Color) (seq []byte) {
	seq = strconv.AppendInt([]byte{27, '['}, int64(c), 10)
	for _, other := range others {
		seq = strconv.AppendInt(append(seq, ';'), int64(other), 10)
	}
	seq = append(seq, 'm')
	return
}

type LevelInfo struct {
	Name         string
	Abbreviation string
	Color        Color
}

var Levels = map[Level]LevelInfo{
	Verbose: {"verbose", "VRBSE", Blue},
	Trace:   {"trace", "TRACE", Green},
	Debug:   {"debug", "DEBUG", Cyan},
	Info:    {"info", "INFO ", NoColor},
	Warn:    {"warn", "WARN ", Yellow},
	Error:   {"error", "ERROR", Red},
	Fatal:   {"fatal", "FATAL", Magenta},
}

func (l Level) Name() string {
	return Levels[l].Name
}

func (l Level) Abbreviation() string {
	return Levels[l].Abbreviation
}

func (l Level) Color() Color {
	return Levels[l].Color
}

func LevelByName(name string) (Level, error) {
	name = strings.ToLower(name)
	for level, info := range Levels {
		if info.Name == name {
			return level, nil
		}
	}
	return 0, InvalidLevelName
}
