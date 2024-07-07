package olog

import (
	"os"

	"github.com/fatih/color"
)

// Fatal prints a message in red and exits the program
func Fatal(content string) {
	d := color.New(color.FgRed, color.BgWhite, color.Bold)
	d.Println("         " + content + "         ")
	os.Exit(0)
}

// Error prints a message in red
func Error(content string) {
	d := color.New(color.FgRed, color.BgWhite, color.Bold)
	d.Println("         " + content + "         ")
}

// Info prints a message in blue
func Info(content string) {
	d := color.New(color.FgHiBlue, color.BgWhite, color.Bold)
	d.Println("         " + content + "         ")
}

// Success prints a message in green
func Success(content string) {
	d := color.New(color.FgGreen, color.BgWhite, color.Bold)
	d.Println("         " + content + "         ")
}
