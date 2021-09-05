package logging

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

var (
	logger *zerolog.Logger
)

func ConfigureDefaultLogging() {
	ConfigureLogging(viper.GetString("loglevel"), viper.GetString("logfile"))
}

// Configures the root logger (see .L()) with the specified global log level and the given file.
func ConfigureLogging(globalLevel, logfile string) {
	level, err := zerolog.ParseLevel(globalLevel)
	if err != nil {
		level = zerolog.WarnLevel
	}
	zerolog.SetGlobalLevel(level)
	var writer io.Writer
	if logfile == "-" {
		writer = zerolog.ConsoleWriter{Out: os.Stderr}
	} else {
		writer, err = os.OpenFile(logfile, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
		if err != nil {
			panic(err)
		}
	}
	l := zerolog.New(writer).With().Timestamp().Logger()
	logger = &l
}

func L() *zerolog.Logger {
	if logger == nil {
		nop := zerolog.Nop()
		return &nop
	}
	return logger
}
