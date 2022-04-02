package request

import (
	"github.com/gin-gonic/gin"

	"github.com/yyliziqiu/waf/auth"
	"github.com/yyliziqiu/waf/errs"
	"github.com/yyliziqiu/waf/logs"
	"github.com/yyliziqiu/waf/response"
)

/**
解析 token info
*/
func ParseTokenInfo(c *gin.Context) {
	identify, err := auth.ParseTokenInfo(c.GetHeader("Token-Info"))
	if err != nil {
		logs.Warn("Unauthorized")
		response.Abort(c, errs.Unauthorized)
		return
	}

	SetIdentify(c, identify)
}

/**
设置验证信息
*/
func SetIdentify(c *gin.Context, identify auth.Identify) {
	c.Set("identify", identify)
}

/**
获取认证信息
*/
func GetIdentify(c *gin.Context) auth.Identify {
	identify, ok := c.Get("identify")
	if ok {
		return identify.(auth.Identify)
	}
	return auth.Identify{}
}

/**
获取认证信息
*/
func GetAccId(c *gin.Context) string {
	return GetIdentify(c).GetAccId()
}

/**
获取认证信息
*/
func GetAccType(c *gin.Context) string {
	return GetIdentify(c).GetAccType()
}

/**
获取认证信息
*/
func GetDeviceId(c *gin.Context) string {
	return GetIdentify(c).GetDeviceId()
}

/**
获取认证信息
*/
func GetDeviceType(c *gin.Context) string {
	return GetIdentify(c).GetDeviceType()
}

/**
是否为登录用户
*/
func IsConsumer(c *gin.Context) bool {
	return GetIdentify(c).IsConsumer()
}

/**
是否为游客
*/
func IsVisitor(c *gin.Context) bool {
	return GetIdentify(c).IsVisitor()
}
