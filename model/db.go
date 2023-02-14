package model

import (
	"aodebiao/conf"
	"aodebiao/ent"
	"aodebiao/log"
	"context"
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"go.uber.org/zap"
	"time"
)

var MysqlDsn string
var Client *ent.Client
func init() {
	MysqlDsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", conf.Setting.Db.UserName, conf.Setting.Db.Password, conf.Setting.Db.Host, conf.Setting.Db.Port, conf.Setting.Db.Name)
}
func InitDb() (*ent.Client, error) {
	db, err := sql.Open("mysql", MysqlDsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	// 从db变量中构造一个ent.Driver对象。
	drv := entsql.OpenDB("mysql", db)
	Client = ent.NewClient(ent.Driver(drv))
	// 根据scheme自动建表
	if err := Client.Schema.Create(context.TODO());err != nil {
		log.Error("数据库连接失败",zap.Error(err))
		panic(err)
	}
	log.Info("client 初始化成功！")
	return Client, nil
}
