package logs

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

var (
	logName      = ""
	logLevel     = LevelDebug
	logServiceId = ""

	out = os.Stdout
	mux = sync.Mutex{}
)

func Initialize(c Config) {
	logName = c.Name
	logLevel = parseLevelStr(c.Level)
	logServiceId = c.ServiceId
}

func log(level uint8, value interface{}, requestField RequestField) {
	if level < logLevel {
		return
	}

	fn, file, line := stack(3)
	var message string
	switch value.(type) {
	case string:
		message = value.(string)
	case error:
		message = value.(error).Error()
	default:
		message = fmt.Sprintf("%#v", value)
	}

	encoder := json.NewEncoder(out)
	encoder.SetEscapeHTML(false)

	_ = encoder.Encode(Entry{
		Request: requestField,
		Log: LogField{
			Level:     parseLevel(level),
			Timestamp: time.Now().Unix(),
			Detail: DetailField{
				Name:      logName,
				ServiceId: logServiceId,
				File:      file,
				Line:      line,
				Func:      fn,
				Message:   message,
			},
		},
	})
}

func parseLevelStr(levelStr string) uint8 {
	switch levelStr {
	case "info":
		return LevelInfo
	case "warning":
		return LevelWarn
	case "error":
		return LevelError
	case "fatal":
		return LevelFatal
	default:
		return LevelDebug
	}
}

func parseLevel(level uint8) string {
	switch level {
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warning"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	default:
		return "debug"
	}
}

func stack(skip int) (fn string, file string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}

	fnp := runtime.FuncForPC(pc)
	if fnp == nil {
		fn = "???"
	} else {
		fn = fnp.Name()
	}

	return
}

func Debug(value interface{}) {
	log(LevelDebug, value, RequestField{})
}

func Info(value interface{}) {
	log(LevelInfo, value, RequestField{})
}

func Warn(value interface{}) {
	log(LevelWarn, value, RequestField{})
}

func Error(value interface{}) {
	log(LevelError, value, RequestField{})
}

func Fatal(value interface{}) {
	log(LevelFatal, value, RequestField{})
	os.Exit(1)
}

func Debugf(message string, args ...interface{}) {
	log(LevelDebug, fmt.Sprintf(message, args...), RequestField{})
}

func Infof(message string, args ...interface{}) {
	log(LevelInfo, fmt.Sprintf(message, args...), RequestField{})
}

func Warnf(message string, args ...interface{}) {
	log(LevelWarn, fmt.Sprintf(message, args...), RequestField{})
}

func Errorf(message string, args ...interface{}) {
	log(LevelError, fmt.Sprintf(message, args...), RequestField{})
}

func Fatalf(message string, args ...interface{}) {
	log(LevelFatal, fmt.Sprintf(message, args...), RequestField{})
	os.Exit(1)
}
