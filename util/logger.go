package util

import (
	"log"
)

// Log is a daa structure to populate logger instance.
type Log struct {
	namespace string
}

// Logger is an interface that exposses methods associates with log.
type Logger interface {
	Debug(err error)
	Error(err error)
	Info(message interface{})
}

// NewLogger initiates new logger.
func NewLogger(ns string) Logger {
	return &Log{
		namespace: ns,
	}
}

// Debug
func (l *Log) Debug(err error) {
	log.Printf("[ %s ]: %v : %v", l.namespace, err.Error(), err)
}

// Error
func (l *Log) Error(err error) {
	log.Fatalf("[ %s ]: %v : %v", l.namespace, err.Error(), err)
}

// Info
func (l *Log) Info(message interface{}) {
	log.Printf("[ %s ]: %v", l.namespace, message)
}
