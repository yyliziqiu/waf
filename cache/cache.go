package cache

import "time"

var configMap map[string]Config

func Initialize(configs ...Config) {
	configMap = make(map[string]Config, len(configs))
	for _, config := range configs {
		configMap[config.Name] = config
	}
}

func GetCacheConfigDuration(name string) (time.Duration, bool) {
	cf, ok := configMap[name]
	if !ok {
		return 0, false
	}
	return cf.Duration, true
}
