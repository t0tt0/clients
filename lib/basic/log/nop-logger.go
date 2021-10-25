package logger

import (
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"os"
)

type NopLogger struct{}

// NewNopLogger returns a logger that doesn't do anything.
func NewNopLogger() logger.Logger { return &NopLogger{} }

func (*NopLogger) Info(string, ...interface{})                 {}
func (*NopLogger) Debug(string, ...interface{})                {}
func (*NopLogger) Error(string, ...interface{})                {}
func (l *NopLogger) Fatal(msg string, keyvals ...interface{})  { os.Exit(1) }
func (l *NopLogger) Warn(msg string, keyvals ...interface{})   {}
func (l *NopLogger) With(keyvals ...interface{}) logger.Logger { return l }
