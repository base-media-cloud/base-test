package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func New() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	return zerolog.New(os.Stderr).With().Timestamp().Logger()
}
