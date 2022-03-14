package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/**
连接配置
*/
type Config struct {
	Name                   string
	URI                    string
	ReplicaSet             string
	ConnectTimeout         time.Duration
	ServerSelectionTimeout time.Duration
	SocketTimeout          time.Duration
	ReadPreference         string
	MaxPoolSize            int
	MinPoolSize            int
	MaxConnIdleTime        time.Duration
}

/**
获取优先读取选项
*/
func (c *Config) GetReadPreference() *readpref.ReadPref {
	switch c.ReadPreference {
	case "PrimaryMode":
		return readpref.Primary()
	case "PrimaryPreferredMode":
		return readpref.PrimaryPreferred()
	case "SecondaryMode":
		return readpref.Secondary()
	case "SecondaryPreferredMode":
		return readpref.SecondaryPreferred()
	case "NearestMode":
		return readpref.Nearest()
	default:
		return readpref.PrimaryPreferred()
	}
}

/**
连接
*/
type Connection struct {
	Client *mongo.Client
	Config Config
}

/**
断开连接
*/
func (c *Connection) Disconnect() {
	_ = c.Client.Disconnect(context.Background())
	c.Client = nil
}
