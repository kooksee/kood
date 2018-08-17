package kfs

import (
	"github.com/kooksee/kdb"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
)

var log *logs.BeeLogger

func InitLog() {
	log = beego.BeeLogger
}

var db kdb.IKDB

func InitDb(db1 kdb.IKDB) {
	db = db1
}
