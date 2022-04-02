package cache

import (
	"time"

	"github.com/yyliziqiu/waf/logs"
)

var configMap map[string]Config

func Initialize(configs ...Config) {
	configMap = make(map[string]Config, len(configs))
	for _, config := range configs {
		configMap[config.Name] = config
	}
}

func GetCacheDuration(name string) time.Duration {
	cf, ok := configMap[name]
	if !ok {
		logs.Fatalf("未发现 { %s } 缓存配置", name)
	}
	return cf.Duration
}
