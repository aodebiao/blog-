package model

import (
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/ent"
	"time"
)

func init()  {

}
func Open() (*ent.Client, error) {
	db, err := sql.Open("mysql", "<mysql-dsn>")
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	// 从db变量中构造一个ent.Driver对象。
	drv := entsql.OpenDB("mysql", db)
	return ent.NewClient(ent.Driver(drv)), nil
}
