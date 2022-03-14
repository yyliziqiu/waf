package mysqldb

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/yyliziqiu/waf/ylog"
)

/**
所有 mysql 连接
*/
var connections map[string]*Connection

/**
初始化数据库
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
重新连接数据库
*/
func New(config Config) (*Connection, error) {
	db, err := sql.Open("mysql", config.DSN)
	if err != nil {
		return nil, err
	}

	if config.MaxOpenConns > 0 {
		db.SetMaxOpenConns(config.MaxOpenConns)
	}
	if config.MaxIdleConns > 0 {
		db.SetMaxIdleConns(config.MaxIdleConns)
	}
	if config.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(config.ConnMaxLifetime)
	}

	connection := &Connection{DB: db, Config: config}

	return connection, nil
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
获取连接
*/
func GetConnection(name string) *Connection {
	return connections[name]
}

/**
获取数据库实例
*/
func GetDB(name string) *sql.DB {
	return GetConnection(name).DB
}
