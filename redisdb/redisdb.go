package redisdb

import (
	"context"

	"github.com/go-redis/redis/v8"

	"github.com/yyliziqiu/waf/logs"
)

/**
所有 redis 连接
*/
var connections map[string]*Connection

/**
初始化
*/
func Initialize(configs ...Config) {
	connections = make(map[string]*Connection, len(configs))
	for _, config := range configs {
		connection, err := New(config)
		if err != nil {
			Finally()
			logs.Fatal(err)
		}
		connections[config.Name] = connection
	}
}

/**
构造方法
*/
func New(config Config) (*Connection, error) {
	var (
		connection *Connection
		err        error
	)

	if config.IsCluster {
		connection, err = NewClusterClient(config)
		if err != nil {
			return nil, err
		}
	} else {
		connection, err = NewClient(config)
		if err != nil {
			return nil, err
		}
	}

	if err := connection.Ping(); err != nil {
		return nil, err
	}

	return connection, nil
}

/**
集群
*/
func NewClusterClient(config Config) (*Connection, error) {
	ops := &redis.ClusterOptions{
		Addrs: config.Addrs,
	}
	switch config.ReadPreference {
	case "ReadOnly":
		ops.ReadOnly = true
	case "RouteByLatency":
		ops.RouteByLatency = true
	case "RouteRandomly":
		ops.RouteRandomly = true
	}
	if config.Username != "" {
		ops.Username = config.Username
	}
	if config.Password != "" {
		ops.Password = config.Password
	}
	if config.MaxRetries != 0 {
		ops.MaxRetries = config.MaxRetries
	}
	if config.DialTimeout > 0 {
		ops.DialTimeout = config.DialTimeout
	}
	if config.ReadTimeout > 0 {
		ops.ReadTimeout = config.ReadTimeout
	}
	if config.WriteTimeout > 0 {
		ops.WriteTimeout = config.WriteTimeout
	}
	if config.MinIdleConns > 0 {
		ops.MinIdleConns = config.MinIdleConns
	}
	if config.MaxConnAge > 0 {
		ops.MaxConnAge = config.MaxConnAge
	}
	if config.PoolSize > 0 {
		ops.PoolSize = config.PoolSize
	}
	if config.PoolTimeout > 0 {
		ops.PoolTimeout = config.PoolTimeout
	}
	if config.IdleTimeout > 0 {
		ops.IdleTimeout = config.IdleTimeout
	}
	if config.IdleCheckFrequency > 0 {
		ops.IdleCheckFrequency = config.IdleCheckFrequency
	}

	cli := redis.NewClusterClient(ops)

	return &Connection{ClusterClient: cli, Config: config}, nil
}

/**
单机
*/
func NewClient(config Config) (*Connection, error) {
	ops := &redis.Options{
		Addr: config.Addr,
		DB:   config.DB,
	}
	if config.Username != "" {
		ops.Username = config.Username
	}
	if config.Password != "" {
		ops.Password = config.Password
	}
	if config.MaxRetries != 0 {
		ops.MaxRetries = config.MaxRetries
	}
	if config.DialTimeout > 0 {
		ops.DialTimeout = config.DialTimeout
	}
	if config.ReadTimeout > 0 {
		ops.ReadTimeout = config.ReadTimeout
	}
	if config.WriteTimeout > 0 {
		ops.WriteTimeout = config.WriteTimeout
	}
	if config.MinIdleConns > 0 {
		ops.MinIdleConns = config.MinIdleConns
	}
	if config.MaxConnAge > 0 {
		ops.MaxConnAge = config.MaxConnAge
	}
	if config.PoolSize > 0 {
		ops.PoolSize = config.PoolSize
	}
	if config.PoolTimeout > 0 {
		ops.PoolTimeout = config.PoolTimeout
	}
	if config.IdleTimeout > 0 {
		ops.IdleTimeout = config.IdleTimeout
	}
	if config.IdleCheckFrequency > 0 {
		ops.IdleCheckFrequency = config.IdleCheckFrequency
	}

	cli := redis.NewClient(ops)

	return &Connection{Client: cli, Config: config}, nil
}

/**
资源回收
*/
func Finally() {
	for _, collection := range connections {
		collection.Close()
	}
}

/**
获取 redis 连接
*/
func GetConnection(name string) *Connection {
	return connections[name]
}

/**
获取 client,context
*/
func GetCliCtx(name string) (*redis.Client, context.Context) {
	connection := GetConnection(name)
	return connection.Client, context.Background()
}

/**
获取 clusterClient,context
*/
func GetCluCtx(name string) (*redis.ClusterClient, context.Context) {
	connection := GetConnection(name)
	return connection.ClusterClient, context.Background()
}
