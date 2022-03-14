package middleware

import (
	"bytes"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"

	"github.com/yyliziqiu/waf/errs"
	"github.com/yyliziqiu/waf/response"
	"github.com/yyliziqiu/waf/ylog"
)

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

type errStack struct {
	Error interface{} `json:"error"`
	Stack []string    `json:"stack"`
}

func Recover() gin.HandlerFunc {
	return gin.CustomRecoveryWithWriter(nil, func(c *gin.Context, err interface{}) {
		ylog.WithFields(ylog.Fields{
			"error": err,
			"stack": errStack{
				Error: err,
				Stack: stack(3),
			},
		}).Error("panic")
		response.Abort(c, errs.InternalServerError)
	})
}

func stack(skip int) []string {
	var es []string
	for i := skip; ; i++ { // Skip the expected number of frames
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		es = append(es, fmt.Sprintf("%s:%d %s", file, line, string(function(pc))))
	}
	return es
}

func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	if lastSlash := bytes.LastIndex(name, slash); lastSlash >= 0 {
		name = name[lastSlash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}
