package log

import (
	"github.com/rs/zerolog"
)

// ParseLevel parses a string which represents a log level and returns
// a zerolog.Level.
func ParseLevel(level string, defaultLevel zerolog.Level) zerolog.Level {
	_ = "STUB: not implemented"
	return *new(zerolog.Level)
}
