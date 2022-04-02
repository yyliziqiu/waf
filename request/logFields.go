package request

import (
	"github.com/gin-gonic/gin"

	"github.com/yyliziqiu/waf/logs"
)

func GetLogRequestField(c *gin.Context) logs.RequestField {
	return logs.RequestField{
		XRequestId: c.GetHeader("X-Request-Id"),
		Method:     c.Request.Method,
		Path:       c.Request.RequestURI,
	}
}
