package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/yyliziqiu/waf/errs"
	"github.com/yyliziqiu/waf/logs"
	"github.com/yyliziqiu/waf/request"
	"github.com/yyliziqiu/waf/response"
)

func PermitConsumer() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !request.IsConsumer(c) {
			logs.Warn("permission forbidden, only consumer allow")
			response.Abort(c, errs.PermissionForbidden)
			return
		}
	}
}

func PermitVisitor() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !request.IsVisitor(c) {
			logs.Warn("permission forbidden, only visitor allow")
			response.Abort(c, errs.PermissionForbidden)
			return
		}
	}
}

func PermitAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !request.IsConsumer(c) && !request.IsConsumer(c) {
			logs.Warn("permission forbidden")
			response.Abort(c, errs.PermissionForbidden)
			return
		}
	}
}
