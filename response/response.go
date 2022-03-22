package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yyliziqiu/waf/errs"
)

/**
服务编号
*/
var serviceNo = 0

/**
设置服务编号。serviceNo * 1000 + code
*/
func SetServiceNo(no int) {
	serviceNo = no * 1000
}

/**
响应成功
*/
func Success(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, result)
}

/**
响应成功
*/
func SuccessOK(c *gin.Context) {
	c.String(http.StatusOK, "")
}

/**
响应成功
*/
func OK(c *gin.Context) {
	c.String(http.StatusOK, "")
}

/**
请求终止
*/
func Abort(c *gin.Context, err error) {
	e := prepare(err)
	c.AbortWithStatusJSON(e.StatusCode, newErrorDto(e.Code, e.Message))
}

/**
请求终止
*/
func AbortE(c *gin.Context, err error, err2 error) {
	e := prepare(err)
	c.AbortWithStatusJSON(e.StatusCode, newErrorDto(e.Code, err2.Error()))
}

/**
请求终止
*/
func AbortM(c *gin.Context, err error, message string) {
	e := prepare(err)
	c.AbortWithStatusJSON(e.StatusCode, newErrorDto(e.Code, message))
}

/**
请求终止
*/
func AbortMF(c *gin.Context, err error, message string, a ...interface{}) {
	e := prepare(err)
	c.AbortWithStatusJSON(e.StatusCode, newErrorDto(e.Code, fmt.Sprintf(message, a...)))
}

/**
请求失败
*/
func Fail(c *gin.Context, err error) {
	e := prepare(err)
	c.JSON(e.StatusCode, newErrorDto(e.Code, e.Message))
}

/**
请求失败
*/
func FailE(c *gin.Context, err error, err2 error) {
	e := prepare(err)
	c.JSON(e.StatusCode, newErrorDto(e.Code, err2.Error()))
}

/**
请求失败
*/
func FailM(c *gin.Context, err error, message string) {
	e := prepare(err)
	c.JSON(e.StatusCode, newErrorDto(e.Code, message))
}

/**
请求失败
*/
func FailMF(c *gin.Context, err error, message string, a ...interface{}) {
	e := prepare(err)
	c.JSON(e.StatusCode, newErrorDto(e.Code, fmt.Sprintf(message, a...)))
}

/**
预处理 err
*/
func prepare(err error) *errs.Error {
	if e, ok := err.(*errs.Error); ok {
		return e
	}
	return errs.BadRequest
}

/**
new error dto
*/
func newErrorDto(code int, message string) errorDto {
	return errorDto{Code: serviceNo + code, Message: message}
}
