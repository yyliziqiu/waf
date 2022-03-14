package mysqldb

import (
	"database/sql"
	"time"
)

type Config struct {
	Name            string
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type Connection struct {
	DB     *sql.DB
	Config Config
}

/**
关闭连接
*/
func (c *Connection) Close() {
	if c.DB != nil {
		_ = c.DB.Close()
	}
}
