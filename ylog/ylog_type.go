package ylog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

const (
	LevelDebug uint8 = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

type Logger struct {
	serviceId string
	name      string
	level     uint8
	out       io.Writer
	mu        sync.Mutex
}

func (l *Logger) write(data []byte) error {
	_, err := l.out.Write(data)
	if err != nil {
		return err
	}
	return nil
}

type Entry struct {
	Request RequestField `json:"request"`
	Log     LogField     `json:"log"`

	logger *Logger
	err    error
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
	Fn        string `json:"fn"`
	Message   string `json:"message"`
	Extra     Fields `json:"extra,omitempty"`
}

type Fields map[string]interface{}

func (e *Entry) setRequest(request RequestField) *Entry {
	e.Request = request
	return e
}

func (e *Entry) setLevel(level uint8) *Entry {
	e.Log.Level = LevelToString(level)
	return e
}

func (e *Entry) setMessage(message string) *Entry {
	e.Log.Detail.Message = message
	return e
}

func (e *Entry) GetError() error {
	return e.err
}

func (e *Entry) log(level uint8, request RequestField, message string) {
	if level < e.logger.level {
		return
	}

	e.setLevel(level)
	e.setRequest(request)
	e.setMessage(message)

	encoder := json.NewEncoder(e.logger.out)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(e)
	if err != nil {
		e.err = err
	}
}

func (e *Entry) json() []byte {
	bs, _ := json.Marshal(e)
	return bs
}

func (e *Entry) Debug(message string) {
	e.log(LevelDebug, RequestField{}, message)
}

func (e *Entry) Info(message string) {
	e.log(LevelInfo, RequestField{}, message)
}

func (e *Entry) Warning(message string) {
	e.log(LevelWarning, RequestField{}, message)
}

func (e *Entry) Error(message string) {
	e.log(LevelError, RequestField{}, message)
}

func (e *Entry) Fatal(message string) {
	e.log(LevelFatal, RequestField{}, message)
	os.Exit(1)
}

func (e *Entry) DebugE(err error) {
	e.log(LevelDebug, RequestField{}, err.Error())
}

func (e *Entry) InfoE(err error) {
	e.log(LevelInfo, RequestField{}, err.Error())
}

func (e *Entry) WarningE(err error) {
	e.log(LevelWarning, RequestField{}, err.Error())
}

func (e *Entry) ErrorE(err error) {
	e.log(LevelError, RequestField{}, err.Error())
}

func (e *Entry) FatalE(err error) {
	e.log(LevelFatal, RequestField{}, err.Error())
	os.Exit(1)
}

func (e *Entry) DebugF(message string, args ...interface{}) {
	e.log(LevelDebug, RequestField{}, fmt.Sprintf(message, args...))
}

func (e *Entry) InfoF(message string, args ...interface{}) {
	e.log(LevelInfo, RequestField{}, fmt.Sprintf(message, args...))
}

func (e *Entry) WarningF(message string, args ...interface{}) {
	e.log(LevelWarning, RequestField{}, fmt.Sprintf(message, args...))
}

func (e *Entry) ErrorF(message string, args ...interface{}) {
	e.log(LevelError, RequestField{}, fmt.Sprintf(message, args...))
}

func (e *Entry) FatalF(message string, args ...interface{}) {
	e.log(LevelFatal, RequestField{}, fmt.Sprintf(message, args...))
	os.Exit(1)
}

func (e *Entry) DebugR(request RequestField, message string) {
	e.log(LevelDebug, request, message)
}

func (e *Entry) InfoR(request RequestField, message string) {
	e.log(LevelInfo, request, message)
}

func (e *Entry) WarningR(request RequestField, message string) {
	e.log(LevelWarning, request, message)
}

func (e *Entry) ErrorR(request RequestField, message string) {
	e.log(LevelError, request, message)
}

func (e *Entry) FatalR(request RequestField, message string) {
	e.log(LevelFatal, request, message)
	os.Exit(1)
}

func (e *Entry) DebugRE(request RequestField, err error) {
	e.log(LevelDebug, request, err.Error())
}

func (e *Entry) InfoRE(request RequestField, err error) {
	e.log(LevelInfo, request, err.Error())
}

func (e *Entry) WarningRE(request RequestField, err error) {
	e.log(LevelWarning, request, err.Error())
}

func (e *Entry) ErrorRE(request RequestField, err error) {
	e.log(LevelError, request, err.Error())
}

func (e *Entry) FatalRE(request RequestField, err error) {
	e.log(LevelFatal, request, err.Error())
	os.Exit(1)
}

func (e *Entry) DebugRF(request RequestField, message string, args ...interface{}) {
	e.log(LevelDebug, request, fmt.Sprintf(message, args...))
}

func (e *Entry) InfoRF(request RequestField, message string, args ...interface{}) {
	e.log(LevelInfo, request, fmt.Sprintf(message, args...))
}

func (e *Entry) WarningRF(request RequestField, message string, args ...interface{}) {
	e.log(LevelWarning, request, fmt.Sprintf(message, args...))
}

func (e *Entry) ErrorRF(request RequestField, message string, args ...interface{}) {
	e.log(LevelError, request, fmt.Sprintf(message, args...))
}

func (e *Entry) FatalRF(request RequestField, message string, args ...interface{}) {
	e.log(LevelFatal, request, fmt.Sprintf(message, args...))
	os.Exit(1)
}
