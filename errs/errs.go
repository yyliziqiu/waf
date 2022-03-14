package errs

/**
1~7×× 客户端错误
8××   服务端错误
9××   第三方服务错误

1×× / 80× / 90× http 错误
*/
var (
	// http
	BadRequest           = &Error{StatusCode: 400, Code: 100, Message: "Bad Request"}
	Unauthorized         = &Error{StatusCode: 401, Code: 101, Message: "Unauthorized"}
	PermissionForbidden  = &Error{StatusCode: 403, Code: 102, Message: "Permission Forbidden"}
	NotFound             = &Error{StatusCode: 404, Code: 103, Message: "Not Found"}
	MethodNotAllowed     = &Error{StatusCode: 405, Code: 104, Message: "Method Not Allowed"}
	InternalServerError  = &Error{StatusCode: 500, Code: 800, Message: "Internal Server Error"}
	ServiceResponseError = &Error{StatusCode: 504, Code: 900, Message: "Service Response Error"}

	// client error
	Retry           = &Error{StatusCode: 400, Code: 210, Message: "Retry"}
	IPLimit         = &Error{StatusCode: 400, Code: 211, Message: "IP Limit"}
	ParamsError     = &Error{StatusCode: 400, Code: 212, Message: "Params Error"}
	NotExist        = &Error{StatusCode: 400, Code: 213, Message: "Not Exist"}
	HasExisted      = &Error{StatusCode: 400, Code: 214, Message: "Has Existed"}
	HasDeleted      = &Error{StatusCode: 400, Code: 215, Message: "Has Deleted"}
	HasCanceled     = &Error{StatusCode: 400, Code: 216, Message: "Has Cancel"}
	NotStart        = &Error{StatusCode: 400, Code: 217, Message: "Not Start"}
	HasStopped      = &Error{StatusCode: 400, Code: 218, Message: "Has Stop"}
	ForbiddenAction = &Error{StatusCode: 400, Code: 219, Message: "Forbidden Action"}

	// server error

	// third service error

)
