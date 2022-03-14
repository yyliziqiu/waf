package request

import "github.com/gin-gonic/gin"

/**
获取客户端语言
*/
func GetLang(c *gin.Context) string {
	return c.GetString("lang")
}
