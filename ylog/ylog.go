package ylog

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

var logger = newLogger()

func newLogger() *Logger {
	return &Logger{
		level: LevelDebug,
		out:   os.Stdout,
	}
}

func newEntry(logger *Logger) *Entry {
	entry := &Entry{
		logger:  logger,
		Request: RequestField{},
		Log:     LogField{},
	}

	fs, file, line := stack(3)
	entry.Log.Detail.ServiceId = logger.serviceId
	entry.Log.Detail.Name = logger.name
	entry.Log.Detail.Fn = fs
	entry.Log.Detail.File = file
	entry.Log.Detail.Line = line
	entry.Log.Timestamp = time.Now().Unix()

	return entry
}

func stack(skip int) (fs string, file string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		fs = "???"
	} else {
		fs = fn.Name()
	}

	return
}

func SetServiceId(serviceId string) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.serviceId = serviceId
}

func SetName(name string) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.name = name
}

func SetLevelString(level string) {
	SetLevel(ParseLevel(level))
}

func SetLevel(level uint8) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.level = level
}

func SetOutput(out io.Writer) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.out = out
}

func ParseLevel(level string) uint8 {
	switch level {
	case "debug":
		return LevelDebug
	case "info":
		return LevelInfo
	case "warning":
		return LevelWarning
	case "error":
		return LevelError
	case "fatal":
		return LevelFatal
	default:
		return LevelDebug
	}
}

func LevelToString(level uint8) string {
	switch level {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarning:
		return "warning"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	default:
		return "debug"
	}
}

func WithError(err error) *Entry {
	entry := newEntry(logger)
	entry.Log.Detail.Extra = Fields{"error": err.Error()}
	return entry
}

func WithFields(fields Fields) *Entry {
	entry := newEntry(logger)
	entry.Log.Detail.Extra = fields
	return entry
}

func Debug(message string) {
	if LevelDebug < logger.level {
		return
	}
	newEntry(logger).Debug(message)
}

func Info(message string) {
	if LevelInfo < logger.level {
		return
	}
	newEntry(logger).Info(message)
}

func Warning(message string) {
	if LevelWarning < logger.level {
		return
	}
	newEntry(logger).Warning(message)
}

func Error(message string) {
	if LevelError < logger.level {
		return
	}
	newEntry(logger).Error(message)
}

func Fatal(message string) {
	if LevelFatal < logger.level {
		return
	}
	newEntry(logger).Fatal(message)
}

func DebugE(err error) {
	if LevelDebug < logger.level {
		return
	}
	newEntry(logger).DebugE(err)
}

func InfoE(err error) {
	if LevelInfo < logger.level {
		return
	}
	newEntry(logger).InfoE(err)
}

func WarningE(err error) {
	if LevelWarning < logger.level {
		return
	}
	newEntry(logger).WarningE(err)
}

func ErrorE(err error) {
	if LevelError < logger.level {
		return
	}
	newEntry(logger).ErrorE(err)
}

func FatalE(err error) {
	if LevelFatal < logger.level {
		return
	}
	newEntry(logger).FatalE(err)
}

func DebugF(message string, args ...interface{}) {
	if LevelDebug < logger.level {
		return
	}
	newEntry(logger).DebugF(message, args...)
}

func InfoF(message string, args ...interface{}) {
	if LevelInfo < logger.level {
		return
	}
	newEntry(logger).InfoF(message, args...)
}

func WarningF(message string, args ...interface{}) {
	if LevelWarning < logger.level {
		return
	}
	newEntry(logger).WarningF(message, args...)
}

func ErrorF(message string, args ...interface{}) {
	if LevelError < logger.level {
		return
	}
	newEntry(logger).ErrorF(message, args...)
}

func FatalF(message string, args ...interface{}) {
	if LevelFatal < logger.level {
		return
	}
	newEntry(logger).FatalF(message, args...)
}

func DebugR(request RequestField, message string) {
	if LevelDebug < logger.level {
		return
	}
	newEntry(logger).DebugR(request, message)
}

func InfoR(request RequestField, message string) {
	if LevelInfo < logger.level {
		return
	}
	newEntry(logger).InfoR(request, message)
}

func WarningR(request RequestField, message string) {
	if LevelWarning < logger.level {
		return
	}
	newEntry(logger).WarningR(request, message)
}

func ErrorR(request RequestField, message string) {
	if LevelError < logger.level {
		return
	}
	newEntry(logger).ErrorR(request, message)
}

func FatalR(request RequestField, message string) {
	if LevelFatal < logger.level {
		return
	}
	newEntry(logger).FatalR(request, message)
}

func DebugRE(request RequestField, err error) {
	if LevelDebug < logger.level {
		return
	}
	newEntry(logger).DebugRE(request, err)
}

func InfoRE(request RequestField, err error) {
	if LevelInfo < logger.level {
		return
	}
	newEntry(logger).InfoRE(request, err)
}

func WarningRE(request RequestField, err error) {
	if LevelWarning < logger.level {
		return
	}
	newEntry(logger).WarningRE(request, err)
}

func ErrorRE(request RequestField, err error) {
	if LevelError < logger.level {
		return
	}
	newEntry(logger).ErrorRE(request, err)
}

func FatalRE(request RequestField, err error) {
	if LevelFatal < logger.level {
		return
	}
	newEntry(logger).FatalRE(request, err)
}

func DebugRF(request RequestField, message string, args ...interface{}) {
	if LevelDebug < logger.level {
		return
	}
	newEntry(logger).DebugRF(request, message, args...)
}

func InfoRF(request RequestField, message string, args ...interface{}) {
	if LevelInfo < logger.level {
		return
	}
	newEntry(logger).InfoRF(request, message, args...)
}

func WarningRF(request RequestField, message string, args ...interface{}) {
	if LevelWarning < logger.level {
		return
	}
	newEntry(logger).WarningRF(request, message, args...)
}

func ErrorRF(request RequestField, message string, args ...interface{}) {
	if LevelError < logger.level {
		return
	}
	newEntry(logger).ErrorRF(request, message, args...)
}

func FatalRF(request RequestField, message string, args ...interface{}) {
	if LevelFatal < logger.level {
		return
	}
	newEntry(logger).FatalRF(request, message, args...)
}

func Print(message string) {
	_ = logger.write([]byte(time.Now().Format(time.RFC3339) + " " + message))
}

func Println(message string) {
	Print(message + "\n")
}

func PrintE(err error) {
	Println(err.Error())
}

func PrintF(message string, args ...interface{}) {
	Print(fmt.Sprintf(message, args...))
}
