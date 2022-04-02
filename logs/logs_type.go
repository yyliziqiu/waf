package logs

import (
	"fmt"
	"os"
)

const (
	LevelDebug uint8 = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

type Config struct {
	Name      string
	Level     string
	ServiceId string
}

type Entry struct {
	Request RequestField `json:"request"`
	Log     LogField     `json:"log"`
}

type RequestField struct {
	XRequestId string `json:"x-request-id,omitempty"`
	Method     string `json:"method,omitempty"`
	Path       string `json:"path,omitempty"`
}

type LogField struct {
	Level     string      `json:"level"`
	Timestamp int64       `json:"timestamp"`
	Detail    DetailField `json:"detail"`
}

type DetailField struct {
	ServiceId string `json:"service_id,omitempty"`
	Name      string `json:"name,omitempty"`
	File      string `json:"file"`
	Line      int    `json:"line"`
	Func      string `json:"func"`
	Message   string `json:"message"`
}

func (rf RequestField) Debug(value interface{}) {
	log(LevelDebug, value, rf)
}

func (rf RequestField) Info(value interface{}) {
	log(LevelInfo, value, rf)
}

func (rf RequestField) Warn(value interface{}) {
	log(LevelWarn, value, rf)
}

func (rf RequestField) Error(value interface{}) {
	log(LevelError, value, rf)
}

func (rf RequestField) Fatal(value interface{}) {
	log(LevelFatal, value, rf)
	os.Exit(1)
}

func (rf RequestField) Debugf(message string, args ...interface{}) {
	log(LevelDebug, fmt.Sprintf(message, args...), rf)
}

func (rf RequestField) Infof(message string, args ...interface{}) {
	log(LevelInfo, fmt.Sprintf(message, args...), rf)
}

func (rf RequestField) Warnf(message string, args ...interface{}) {
	log(LevelWarn, fmt.Sprintf(message, args...), rf)
}

func (rf RequestField) Errorf(message string, args ...interface{}) {
	log(LevelError, fmt.Sprintf(message, args...), rf)
}

func (rf RequestField) Fatalf(message string, args ...interface{}) {
	log(LevelFatal, fmt.Sprintf(message, args...), rf)
	os.Exit(1)
}
