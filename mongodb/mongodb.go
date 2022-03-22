package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/yyliziqiu/waf/ylog"
)

/**
所有 mongo 连接
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
			ylog.FatalE(err)
		}
		connections[config.Name] = connection
	}
}

/**
构造方法
*/
func New(config Config) (*Connection, error) {
	ops := options.Client()
	ops.ApplyURI(config.URI)
	ops.SetReadPreference(config.GetReadPreference())
	ops.SetRetryReads(false)
	ops.SetRetryWrites(false)
	if config.ReplicaSet != "" {
		ops.SetReplicaSet(config.ReplicaSet)
	}
	if config.ConnectTimeout > 0 {
		ops.SetConnectTimeout(config.ConnectTimeout)
	}
	if config.ServerSelectionTimeout > 0 {
		ops.SetServerSelectionTimeout(config.ServerSelectionTimeout)
	}
	if config.SocketTimeout > 0 {
		ops.SetSocketTimeout(config.SocketTimeout)
	}
	if config.MaxPoolSize > 0 {
		ops.SetMaxPoolSize(uint64(config.MaxPoolSize))
	}
	if config.MinPoolSize > 0 {
		ops.SetMinPoolSize(uint64(config.MinPoolSize))
	}
	if config.MaxConnIdleTime > 0 {
		ops.SetMaxConnIdleTime(config.MaxConnIdleTime)
	}

	cli, err := mongo.Connect(context.Background(), ops)
	if err != nil {
		return nil, err
	}

	return &Connection{Client: cli, Config: config}, nil
}

/**
资源回收
*/
func Finally() {
	for _, collection := range connections {
		collection.Disconnect()
	}
}

/**
获取连接
*/
func GetConnection(name string) *Connection {
	return connections[name]
}

/**
获取 client,context
*/
func GetCliCtx(name string) (*mongo.Client, context.Context) {
	connection := GetConnection(name)
	return connection.Client, context.Background()
}

/**
获取 db,context
*/
func GetDBCtx(name string, opts ...*options.DatabaseOptions) (*mongo.Database, context.Context) {
	connection := GetConnection(name)
	return connection.Client.Database(connection.Config.DefaultDB, opts...), context.Background()
}
