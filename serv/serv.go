package serv

import (
	"github.com/gin-gonic/gin"

	"github.com/yyliziqiu/waf/errs"
	"github.com/yyliziqiu/waf/middleware"
	"github.com/yyliziqiu/waf/response"
)

/**
server
*/
var server *APIServer

/**
初始化路由
*/
func InitializeAPIServer(c Config, callbacks ...SetRouteCallback) *APIServer {
	gin.SetMode(c.Env)
	gin.DisableConsoleColor()

	// new server
	server = &APIServer{
		ServiceId: c.ServiceId,
		BasePath:  c.BasePath,
		IP:        c.API.IP,
		Port:      c.API.Port,
		Log:       c.API.Log,
		engine:    gin.New(),
	}
	// register recover middleware
	server.engine.Use(middleware.Recover())
	// register log middleware
	if server.Log {
		server.engine.Use(middleware.Log(server.ServiceId))
	}
	// register error route
	server.engine.NoRoute(func(c *gin.Context) {
		response.Fail(c, errs.NotFound)
	})
	server.engine.NoMethod(func(c *gin.Context) {
		response.Fail(c, errs.MethodNotAllowed)
	})

	// register route
	for _, callback := range callbacks {
		callback(server.engine)
	}

	return server
}
