package middleware

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/yyliziqiu/waf/errs"
	"github.com/yyliziqiu/waf/response"
	"github.com/yyliziqiu/waf/util"
	"github.com/yyliziqiu/waf/ylog"
)

var DefaultCrosHeaders = CrosHeaders{
	Origin:           []string{"*"},
	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
	AllowHeaders:     []string{"*"},
	ExposeHeaders:    []string{},
	AllowCredentials: false,
	MaxAge:           86400,
}

type CrosHeaders struct {
	Origin           []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	MaxAge           int
}

/**
允许跨域
参考： https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers

Only the 7 CORS-safelisted response headers are exposed: Cache-Control, Content-Language, Content-Length, Content-Type, Expires, Last-Modified, Pragma.

CORS-safelisted request header is one of the following HTTP headers: Accept, Accept-Language, Content-Language, Content-Type.
*/
func Cros(ch CrosHeaders) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestOrigin := c.GetHeader("Origin")
		// _ = c.GetHeader("Access-Control-Request-Method")
		// _ = c.GetHeader("Access-Control-Request-Headers")

		if requestOrigin != "" && len(ch.Origin) > 0 {
			if util.ContainsString(ch.Origin, requestOrigin) {
				c.Header("Access-Control-Allow-Origin", requestOrigin)
			}
			if util.ContainsString(ch.Origin, "*") {
				if ch.AllowCredentials == true {
					ylog.Warning("Allowing credentials for wildcard origins is insecure. Please specify more restrictive origins or set 'credentials' to false in your CORS configuration")
					response.Abort(c, errs.InternalServerError)
					return
				}
				c.Header("Access-Control-Allow-Origin", "*")
			}
		}
		if len(ch.AllowHeaders) > 0 {
			c.Header("Access-Control-Allow-Headers", strings.Join(ch.AllowHeaders, ", "))
		}
		if len(ch.AllowMethods) > 0 {
			c.Header("Access-Control-Allow-Methods", strings.Join(ch.AllowMethods, ", "))
		}
		if len(ch.ExposeHeaders) > 0 {
			c.Header("Access-Control-Expose-Headers", strings.Join(ch.ExposeHeaders, ", "))
		}
		if ch.MaxAge > 0 && c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Max-Age", strconv.Itoa(ch.MaxAge))
		}
		c.Header("Access-Control-Allow-Credentials", strconv.FormatBool(ch.AllowCredentials))

		HandleOptionsRequest(c)
	}
}

func HandleOptionsRequest(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		response.SuccessOK(c)
	}
}
