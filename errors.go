package logs

type Err string

func (e Err) Error() string {
	return string(e)
}

const (
	InvalidLevelName Err = "invalid level name"
)
