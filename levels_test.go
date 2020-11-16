package logs_test

import (
	. "github.com/hx/logs"
	. "github.com/hx/logs/testing"
	"testing"
)

func TestLevelByName(t *testing.T) {
	level, err := LevelByName("DEBUG")
	Ok(t, err)
	Equals(t, Debug, level)
	level, err = LevelByName("Verbose")
	Ok(t, err)
	Equals(t, Verbose, level)
	level, err = LevelByName("")
	Equals(t, InvalidLevelName, err)
	Equals(t, Level(0), level)
}
