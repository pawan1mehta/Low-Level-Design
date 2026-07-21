package main

import "fmt"

type Logger struct {
	name string
}

var loggerInstance *Logger

func GetLogger() *Logger {
	if loggerInstance == nil {
		loggerInstance = &Logger{
			name: "AppLogger",
		}
	}
	return loggerInstance
}

func (l *Logger) Log(message string) {
	fmt.Printf("[%s] %s \n", l.name, message)
}

func main() {
	logger1 := GetLogger()
	logger1.Log("First Log")

	logger2 := GetLogger()
	logger2.Log("Second Log")

	fmt.Printf("Same instanc: %v \n", logger1 == logger2)
}
