package serv

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/yyliziqiu/waf/ylog"
)

/**
config
*/
type Config struct {
	ServiceId string
	BasePath  string
	Env       string
	API       APIConfig
	CMD       CMDConfig
}

type APIConfig struct {
	IP   string
	Port string
	Log  bool
}

type CMDConfig struct {
}

/**
server
*/
type APIServer struct {
	ServiceId string
	BasePath  string
	IP        string
	Port      string
	Log       bool
	engine    *gin.Engine
}

/**
启动服务
*/
func (s *APIServer) Run() {
	if err := s.engine.Run(fmt.Sprintf("%s:%s", s.IP, s.Port)); err != nil {
		ylog.FatalE(err)
	}
}

/**
注册路由回调
*/
type SetRouteCallback func(engine *gin.Engine)
