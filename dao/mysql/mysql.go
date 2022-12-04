package mysql

import (
	"fmt"
	"go-web/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		settings.Conf.Mysql.User,
		settings.Conf.Mysql.Password,
		settings.Conf.Mysql.Host,
		settings.Conf.Mysql.Port,
		settings.Conf.Mysql.Name,
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}

	db.SetMaxOpenConns(settings.Conf.Mysql.MaxOpenConnections)
	db.SetMaxIdleConns(settings.Conf.Mysql.MaxIdleConnections)
	return
}

func Close() {
	_ = db.Close()
}
