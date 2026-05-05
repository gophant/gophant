package logger

import "log"

type Logger struct{}

func New(level string) *Logger     { log.Printf("logger initialized, level=%s", level); return &Logger{} }
func (l *Logger) Info(msg string)  { log.Println("INFO:", msg) }
func (l *Logger) Error(msg string) { log.Println("ERROR:", msg) }
