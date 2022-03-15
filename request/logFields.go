package request

import (
	"github.com/gin-gonic/gin"

	"github.com/yyliziqiu/waf/ylog"
)

func GetLogFields(c *gin.Context) ylog.RequestField {
	return ylog.RequestField{
		XRequestId: c.GetHeader("X-Request-Id"),
		Method:     c.Request.Method,
		Path:       c.Request.RequestURI,
	}
}
