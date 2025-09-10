package logger

import (
	"os"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

const (
	miliseconds float32 = 1000000.0
)

var (
	once sync.Once
	log  *Log
)

func GetLogger() Logger {
	once.Do(func() {
		log = NewLog()
	})
	return log
}

type Log struct {
	logger zerolog.Logger
}

func NewLog() *Log {
	return NewLogWithLevel("info")
}

func NewLogWithLevel(level string) *Log {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		lvl = zerolog.InfoLevel
	}
	return &Log{logger: zerolog.New(os.Stdout).Level(lvl).With().Timestamp().Logger()}
}

func (l *Log) Debugf(format string, args ...any) {
	l.logger.Debug().Msgf(format, args...)
}

func (l *Log) Infof(format string, args ...any) {
	l.logger.Info().Msgf(format, args...)
}

func (l *Log) Warningf(moduleName, functionName, format string, args ...any) {
	l.logger.Warn().
		Str("module", moduleName).
		Str("function", functionName).
		Msgf(format, args...)
}

func (l *Log) Error(moduleName, functionName string, err error) {
	l.logger.Error().
		Str("module", moduleName).
		Str("function", functionName).
		Msg(err.Error())
}

func (l *Log) Errorf(moduleName, functionName, format string, args ...any) {
	l.logger.Error().
		Str("module", moduleName).
		Str("function", functionName).
		Msgf(format, args...)
}

func (l *Log) FatalIfError(moduleName, functionName string, errs ...error) {
	var sb strings.Builder
	for _, err := range errs {
		if err != nil {
			sb.WriteString(err.Error())
			sb.WriteString("\t")
		}
	}
	if sb.Len() > 0 {
		l.Fatalf(moduleName, functionName, "%s", sb.String())
	}
}

func (l *Log) Fatal(moduleName, functionName string, err error) {
	l.Fatalf(moduleName, functionName, "%s", err.Error())
}

func (l *Log) Fatalf(moduleName, functionName, format string, args ...any) {
	l.logger.Fatal().
		Str("module", moduleName).
		Str("function", functionName).
		Msgf(format, args...)
}

func (l *Log) Request(status int, method, requestURI string, responseBody, responseError any, start time.Time) {
	if responseError == nil {
		l.logger.Info().
			Str("method", method).
			Int("status", status).
			Str("request", requestURI).
			Interface("responseBody", responseBody).
			Float32("timestamp", timestamp(start)).
			Msg("")
	} else {
		l.logger.Error().
			Str("method", method).
			Int("status", status).
			Str("request", requestURI).
			Interface("responseBody", responseBody).
			Interface("error", responseError).
			Float32("timestamp", timestamp(start)).
			Msg("")
	}
}

func timestamp(start time.Time) float32 {
	return float32(time.Since(start).Nanoseconds()) / miliseconds
}
