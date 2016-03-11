package loggo

import (
	"log"
)

const (
	DebugLevel = 0
	ErrorLevel = 1
	InfoLevel  = 2
)

const (
	DarkColor      = "\x1b[0m"
	GreenColor     = "\x1b[32;1m"
	TurquoiseColor = "\x1b[36;1m"
	GrayColor      = "\x1b[30;1m"
	RedColor       = "\x1b[31;1m"
	YellowColor    = "\x1b[33;1m"
	BlueColor      = "\x1b[34;1m"
	PinkColor      = "\x1b[35;1m"
	BoldColor      = "\x1b[37;1m"
)

type Logger struct {
	Level int
}

var l = &Logger{InfoLevel}

func SetLevel(level int) {
	l.Level = level
}

func Debug(message string, v ...interface{}) {
	if l.Level <= DebugLevel {
		log.Printf("["+TurquoiseColor+"DEBUG"+DarkColor+"] "+message, v...)
	}
}

func Error(message string, v ...interface{}) {
	if l.Level <= InfoLevel {
		log.Printf("["+RedColor+"ERROR"+DarkColor+"] "+message, v...)
	}
}

func Info(message string, v ...interface{}) {
	if l.Level <= InfoLevel {
		log.Printf("["+GrayColor+"INFO"+DarkColor+"] "+message, v...)
	}
}
