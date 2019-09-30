package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

type level int

// Уровни логирования
const (
	FATAL level = iota
	WARNING
	INFO
)

var currentLevel = WARNING
var levelWords = map[level]string{
	FATAL:   "[FATAL]",
	WARNING: "[WARNING]",
	INFO:    "[INFO]",
}

// SetLevel устанавливает уровень логирования
func SetLevel(l level) {
	currentLevel = l
}

// Print выводит сообщение в лог
func Print(l level, values ...interface{}) {
	if l > currentLevel {
		return
	}

	if l > INFO {
		log.SetOutput(os.Stderr)
	} else {
		log.SetOutput(os.Stdout)
	}

	log.Println(append([]interface{}{levelWords[l]}, values...)...)

	if l == FATAL {
		os.Exit(1)
	}
}

// PrintDuration выводит сообщение о длительности операции, приводя
// длительность к наиболее подходящей размерности
func PrintDuration(startTime time.Time) {
	Print(INFO, fmt.Sprintf("Выполнено за %v", time.Since(startTime)))
}
