package orm

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

type orm struct {
	Db *gorm.DB
}

var db *orm
var once sync.Once

var (
	user     = beego.AppConfig.String("mysqluser")
	pass     = beego.AppConfig.String("mysqlpass")
	urls     = beego.AppConfig.String("mysqlurls")
	database = beego.AppConfig.String("mysqldb")
)

func open() *orm {
	var orm orm
	orm.Db, _ = gorm.Open("mysql", user+":"+pass+"@/"+database+"?charset=utf8&parseTime=True&loc=Local")
	return &orm
}

func GetDb() *gorm.DB {
	once.Do(func() {
		db = open()
	})
	return db.Db
}
