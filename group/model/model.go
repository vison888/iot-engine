package model

import (
	"github.com/vison888/go-vkit/mysqlx"
	"github.com/vison888/iot-engine/group/app"
)

// InitTable 初始化数据库表
func InitTable() {
	// 自动建表
	app.Mysql.GetDB().AutoMigrate(&CategoryModel{})
}

func getTx(tx *mysqlx.MysqlClient) *mysqlx.MysqlClient {
	if tx == nil {
		return app.Mysql
	}
	return tx
}
