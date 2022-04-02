package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/yyliziqiu/waf/errs"
	"github.com/yyliziqiu/waf/logs"
	"github.com/yyliziqiu/waf/response"
	"github.com/yyliziqiu/waf/util"
)

/**
IP 过滤
*/
func IPFilter(allow []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(allow) == 0 {
			return
		}

		match, err := util.IPMMatch(c.ClientIP(), allow)
		if err != nil {
			logs.Warn(err)
			response.Abort(c, errs.IPLimit)
		} else if !match {
			logs.Warn("IP Limit: " + c.ClientIP())
			response.Abort(c, errs.IPLimit)
		}
	}
}
