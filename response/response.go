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
func Ok(c *gin.Context) {
	c.String(http.StatusOK, "")
}

/**
响应成功
*/
func Result(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, result)
}

/**
请求终止
*/
func Abort(c *gin.Context, err error) {
	e := prepareError(err)
	c.AbortWithStatusJSON(e.StatusCode, newErrorResponse(e.Code, e.Message))
}

/**
请求终止
*/
func AbortWithError(c *gin.Context, err error, err2 error) {
	e := prepareError(err)
	c.AbortWithStatusJSON(e.StatusCode, newErrorResponse(e.Code, err2.Error()))
}

/**
请求终止
*/
func AbortWithMessage(c *gin.Context, err error, message string) {
	e := prepareError(err)
	c.AbortWithStatusJSON(e.StatusCode, newErrorResponse(e.Code, message))
}

/**
请求终止
*/
func AbortWithMessageF(c *gin.Context, err error, message string, a ...interface{}) {
	e := prepareError(err)
	c.AbortWithStatusJSON(e.StatusCode, newErrorResponse(e.Code, fmt.Sprintf(message, a...)))
}

/**
请求失败
*/
func Fail(c *gin.Context, err error) {
	e := prepareError(err)
	c.JSON(e.StatusCode, newErrorResponse(e.Code, e.Message))
}

/**
请求失败
*/
func FailWithError(c *gin.Context, err error, err2 error) {
	e := prepareError(err)
	c.JSON(e.StatusCode, newErrorResponse(e.Code, err2.Error()))
}

/**
请求失败
*/
func FailWithMessage(c *gin.Context, err error, message string) {
	e := prepareError(err)
	c.JSON(e.StatusCode, newErrorResponse(e.Code, message))
}

/**
请求失败
*/
func FailWithMessageF(c *gin.Context, err error, message string, a ...interface{}) {
	e := prepareError(err)
	c.JSON(e.StatusCode, newErrorResponse(e.Code, fmt.Sprintf(message, a...)))
}

func prepareError(err error) *errs.Error {
	if e, ok := err.(*errs.Error); ok {
		return e
	}
	return errs.BadRequest
}

/**
new error response
*/
func newErrorResponse(code int, message string) errorResponse {
	return errorResponse{Code: serviceNo + code, Message: message}
}
