package models

import (
	"Blog/pkg/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type Model struct {

	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int64 `json:"created_on"`
	ModifiedOn int64 `json:"modified_on"`
}


//数据库连接的初始化
func init() {
	var (
		err error
		dbName, user, password, host string
		 //dbType,tablePrefix string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	//获取配置,转化为string形式
	//dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	//tablePrefix = sec.Key("TABLE_PREFIX").String()
	dsn:= fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	//gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
	//	return tablePrefix + defaultTableName;
	//}
	//
	//db.SingularTable(true)
	//db.LogMode(true)
	//db.DB().SetMaxIdleConns(10)
	//db.DB().SetMaxOpenConns(100)
}

//func CloseDB() {
//	defer db.Close()
//}