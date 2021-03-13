package console

// Level is a 16-bit set holding the enabled log levels.
type Level uint16

const (
	LevelDebug Level = 1 << iota
	LevelInfo
	LevelError
	LevelWarn
	LevelSuccess
	LevelFatal
)

func (l *Level) set(newLevel Level) {
	*l = newLevel
}

func (l Level) permits(level Level) bool {
	return l <= level
}
