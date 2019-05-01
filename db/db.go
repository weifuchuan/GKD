package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"log"
)

var Engine *xorm.Engine

func init() {
	var err error
	Engine, err = xorm.NewEngine("mysql", "root:root@/gkd?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	Engine.SetMapper(core.SameMapper{})
	err = Engine.CreateTables(new(Account), new(Order), new(Car))
	if err != nil {
		log.Fatal(err)
	}
}
