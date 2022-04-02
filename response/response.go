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
func OK(c *gin.Context) {
	c.String(http.StatusOK, "")
}

/**
响应成功
*/
func Success(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, result)
}

/**
请求终止
*/
func Abort(c *gin.Context, err *errs.Error) {
	c.AbortWithStatusJSON(err.StatusCode, newErrorResponse(err.Code, err.Message))
}

/**
请求终止
*/
func AbortWithError(c *gin.Context, err *errs.Error, err2 error) {
	c.AbortWithStatusJSON(err.StatusCode, newErrorResponse(err.Code, err2.Error()))
}

/**
请求终止
*/
func AbortWithMessage(c *gin.Context, err *errs.Error, message string) {
	c.AbortWithStatusJSON(err.StatusCode, newErrorResponse(err.Code, message))
}

/**
请求终止
*/
func AbortWithMessageF(c *gin.Context, err *errs.Error, message string, a ...interface{}) {
	c.AbortWithStatusJSON(err.StatusCode, newErrorResponse(err.Code, fmt.Sprintf(message, a...)))
}

/**
请求失败
*/
func Fail(c *gin.Context, err *errs.Error) {
	c.JSON(err.StatusCode, newErrorResponse(err.Code, err.Message))
}

/**
请求失败
*/
func FailWithError(c *gin.Context, err *errs.Error, err2 error) {
	c.JSON(err.StatusCode, newErrorResponse(err.Code, err2.Error()))
}

/**
请求失败
*/
func FailWithMessage(c *gin.Context, err *errs.Error, message string) {
	c.JSON(err.StatusCode, newErrorResponse(err.Code, message))
}

/**
请求失败
*/
func FailWithMessageF(c *gin.Context, err *errs.Error, message string, a ...interface{}) {
	c.JSON(err.StatusCode, newErrorResponse(err.Code, fmt.Sprintf(message, a...)))
}

/**
new error response
*/
func newErrorResponse(code int, message string) errorResponse {
	return errorResponse{Code: serviceNo + code, Message: message}
}
