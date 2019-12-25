package clioutput

import (
	"github.com/gookit/color"
)

// Info theme output
func Info(message string) {
	color.Info.Tips(message)
}

// Success theme output
func Success(message string) {
	color.Success.Tips(message)
}

// Error theme output
func Error(message string) {
	color.Error.Tips(message)
}
