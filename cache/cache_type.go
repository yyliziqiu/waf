package cache

import "time"

type Config struct {
	Name     string
	Duration time.Duration
}

type Cache struct {
	Prefix   string
	Duration time.Duration
}

/**
获取缓存键名，时效
*/
func (c *Cache) KeyDuration(suffix string) (string, time.Duration) {
	return c.Prefix + suffix, c.Duration
}

/**
获取缓存键名
*/
func (c *Cache) Key(suffix string) string {
	key, _ := c.KeyDuration(suffix)
	return key
}

/**
获取缓存键名，时效（多）
*/
func (c *Cache) MKeyDuration(suffixes []string) ([]string, time.Duration) {
	keys := make([]string, len(suffixes))
	for i := range suffixes {
		keys[i] = c.Prefix + suffixes[i]
	}

	return keys, c.Duration
}

/**
获取缓存键名（多）
*/
func (c *Cache) MKey(suffixes []string) []string {
	keys, _ := c.MKeyDuration(suffixes)
	return keys
}
