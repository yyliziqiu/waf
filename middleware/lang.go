package middleware

import "github.com/gin-gonic/gin"

func Lang(defaultLang string) gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.Param("hl")
		if lang == "" {
			lang = defaultLang
		}
		c.Set("lang", lang)
	}
}
