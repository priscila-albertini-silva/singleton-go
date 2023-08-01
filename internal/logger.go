package internal

import (
	"fmt"
	"os"
)

type Logger struct {
	logFile *os.File
}

var instance *Logger

func GetLogger() *Logger {
	if instance == nil {
		// Criação da instância do logger, apenas na primeira chamada
		file, err := os.Create("log.txt")
		if err != nil {
			panic("Falha ao criar arquivo de log")
		}
		instance = &Logger{logFile: file}
	}
	return instance
}

func (l *Logger) Log(message string) {
	fmt.Fprintf(l.logFile, "%s\n", message)
}
