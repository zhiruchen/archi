package infrastructure

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/zhiruchen/archi/conf"
)

var Db *sql.DB

func InitMysql(c *conf.MySQLConfig) {
	pool, err := sql.Open("mysql", c.Dsn)
	if err != nil {
		panic(err)
	}
	if c.MaxIdle > 0 {
		pool.SetMaxIdleConns(c.MaxIdle)
	}
	if c.MaxConn > 0 {
		pool.SetMaxOpenConns(c.MaxConn)
	}
	Db = pool
}
