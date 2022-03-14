package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/yyliziqiu/waf/errs"
	"github.com/yyliziqiu/waf/response"
	"github.com/yyliziqiu/waf/util"
	"github.com/yyliziqiu/waf/ylog"
)

/**
IP 过滤
*/
func IPFilter(allow []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(allow) == 0 {
			return
		}

		if match, err := util.IPMMatch(c.ClientIP(), allow); err != nil {
			ylog.WithError(err).Warning("IP limit")
			response.Abort(c, errs.IPLimit)
		} else if !match {
			ylog.Warning("IP limit")
			response.Abort(c, errs.IPLimit)
		}
	}
}
