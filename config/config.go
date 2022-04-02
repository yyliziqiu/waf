package config

import (
	"github.com/spf13/viper"

	"github.com/yyliziqiu/waf/logs"
)

/**
加载配置
*/
func Initialize(path string, c interface{}) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		logs.Fatal(err)
	}
	if err := viper.Unmarshal(c); err != nil {
		logs.Fatal(err)
	}
}
