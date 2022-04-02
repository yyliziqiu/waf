package response

/**
错误信息响应结构
*/
type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
