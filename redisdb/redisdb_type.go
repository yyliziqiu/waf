package redisdb

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

/**
连接配置
*/
type Config struct {
	Name               string
	IsCluster          bool
	ReadPreference     string
	Addrs              []string
	Addr               string
	Username           string
	Password           string
	DB                 int
	MaxRetries         int
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	MinIdleConns       int
	MaxConnAge         time.Duration
	PoolSize           int
	PoolTimeout        time.Duration
	IdleTimeout        time.Duration
	IdleCheckFrequency time.Duration
}

/**
连接
*/
type Connection struct {
	Client        *redis.Client
	ClusterClient *redis.ClusterClient
	Config        Config
}

/**
断开连接
*/
func (c *Connection) Close() {
	if c.Config.IsCluster {
		if c.ClusterClient != nil {
			_ = c.ClusterClient.Close()
			c.ClusterClient = nil
		}
	} else {
		if c.Client != nil {
			_ = c.Client.Close()
			c.Client = nil
		}
	}
}

/**
ping
*/
func (c *Connection) Ping() error {
	if c.Config.IsCluster {
		if _, err := c.ClusterClient.Ping(context.Background()).Result(); err != nil {
			return err
		}
	} else {
		if _, err := c.Client.Ping(context.Background()).Result(); err != nil {
			return err
		}
	}
	return nil
}
