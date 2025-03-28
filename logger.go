package viper

import (
	"fmt"
	"log"
)

// Logger is a unified interface for various logging use cases and practices, including:
// 		- leveled logging
// 		- structured logging
type Logger interface {
	// Trace logs a Trace event.
	//
	// Even more fine-grained information than Debug events.
	// Loggers not supporting this level should fall back to Debug.
	Trace(msg string, keyvals ...interface{})

	// Debug logs a Debug event.
	//
	// A verbose series of information events.
	// They are useful when debugging the system.
	Debug(msg string, keyvals ...interface{})

	// Info logs an Info event.
	//
	// General information about what's happening inside the system.
	Info(msg string, keyvals ...interface{})

	// Warn logs a Warn(ing) event.
	//
	// Non-critical events that should be looked at.
	Warn(msg string, keyvals ...interface{})

	// Error logs an Error event.
	//
	// Critical events that require immediate attention.
	// Loggers commonly provide Fatal and Panic levels above Error level,
	// but exiting and panicing is out of scope for a logging library.
	Error(msg string, keyvals ...interface{})
}

type jwwLogger struct{}

func (jwwLogger) Trace(msg string, keyvals ...interface{}) {
	printf("TRACE", jwwLogMessage(msg, keyvals...))
}

func (jwwLogger) Debug(msg string, keyvals ...interface{}) {
	printf("DEBUG", jwwLogMessage(msg, keyvals...))
}

func (jwwLogger) Info(msg string, keyvals ...interface{}) {
	printf("TRACE", jwwLogMessage(msg, keyvals...))
}

func (jwwLogger) Warn(msg string, keyvals ...interface{}) {
	printf("WARN", jwwLogMessage(msg, keyvals...))
}

func (jwwLogger) Error(msg string, keyvals ...interface{}) {
	printf("ERROR", jwwLogMessage(msg, keyvals...))
}

var defaultLogger Logger

func init() {
	defaultLogger = jwwLogger{}
}

func logger() Logger {
	return defaultLogger
}

func SetLogger(logger Logger) {
	if logger == nil {
		return
	}
	defaultLogger = logger
}

func printf(level, msg string) {
	log.Printf("%s %s", level, msg)
}

func jwwLogMessage(msg string, keyvals ...interface{}) string {
	out := msg

	if len(keyvals) > 0 && len(keyvals)%2 == 1 {
		keyvals = append(keyvals, nil)
	}

	for i := 0; i <= len(keyvals)-2; i += 2 {
		out = fmt.Sprintf("%s %v=%v", out, keyvals[i], keyvals[i+1])
	}

	return out
}
