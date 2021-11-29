package ui

import (
	"fmt"
	"os"
)

// Logger : interface declaration
type Logger struct{}

var successColor = "\033[1;32m%s\033[0m"
var infoColor = "\033[0;34m%s\033[0m"
var warningColor = "\033[1;33m%s\033[0m"
var errorColor = "\033[1;31m%s\033[0m"

// Info : informative messages
func (l *Logger) Info(message string) {
	userInterface.Info(fmt.Sprintf(infoColor, message))
}

// Success : success messages
func (l *Logger) Success(message string) {
	userInterface.Output(fmt.Sprintf(successColor, message))
}

// Warn : warning messages
func (l *Logger) Warn(message string) {
	userInterface.Warn(fmt.Sprintf(warningColor, message))
}

// Error : error/fatal messages
func (l *Logger) Error(message string) {
	userInterface.Error(fmt.Sprintf(errorColor, message))
}

// Output : generic messages
func (l *Logger) Output(message string) {
	userInterface.Output(message)
}

// Debug : debugging messages
func (l *Logger) Debug(message string) {
	if os.Getenv("ODIN_DEBUG") == "yes" {
		userInterface.Output(fmt.Sprintf("[ DEBUG ] %s", message))
	}
}