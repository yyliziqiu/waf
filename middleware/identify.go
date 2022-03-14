package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/yyliziqiu/waf/request"
)

func Identify() gin.HandlerFunc {
	return request.ParseTokenInfo
}
