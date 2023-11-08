package log

import (
	"strings"
)

type Level int

const (
	unset Level = iota
	DEBUG
	INFO
	ERROR
)

func LevelFromString(lvlString string) Level {
	lvlString = strings.ToLower(lvlString)

	switch lvlString {
	case "debug", "dbug":
		return DEBUG
	case "info":
		return INFO
	case "error", "eror":
		return ERROR
	default:
		return DEBUG
	}
}

type Logger interface {
	// debug level logging
	Debug(msg string, fields ...interface{})

	// info level logging
	Info(msg string, fields ...interface{})

	// error level logging
	Error(err error, msg string, fields ...interface{})

	// get a derived Logger with given fields globally set
	Derive(o ...Option) Logger
}

type Option func(*config)

func WithNamespace(namespace string) Option {
	return func(c *config) {
		c.namespace = namespace
	}
}

func WithLevel(level Level) Option {
	return func(c *config) {
		c.minLevel = level
	}
}

func WithFields(fields ...interface{}) Option {
	return func(c *config) {
		c.fields = fields
	}
}

type config struct {
	namespace string
	minLevel  Level
	fields    []interface{}
}
