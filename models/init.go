package models

import (
	"fmt"
	"graphqldemo/service/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/core"
	"xorm.io/xorm"
)

// 数据库连接,读写分离
// 读
var DbRead *xorm.Engine

// 写
var DbWrite *xorm.Engine

// 初始化连接
func Setup() {
	var err error
	//数据库链接,读
	DbRead, err = xorm.NewEngine(config.DbReadSetting.Type, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local",
		config.DbReadSetting.User,
		config.DbReadSetting.Password,
		config.DbReadSetting.Host,
		config.DbReadSetting.Port,
		config.DbReadSetting.Name,
		config.DbReadSetting.Encoding,
	))

	if err != nil {
		log.Printf("创建DbRead连接错误: %v\n", err)
		panic(err)
	}

	// 控制台打印SQL语句
	DbRead.ShowSQL(true)
	DbRead.Logger().SetLevel(core.LOG_WARNING)

	//数据库链接,写
	DbWrite, err = xorm.NewEngine(config.DbWriteSetting.Type, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local",
		config.DbWriteSetting.User,
		config.DbWriteSetting.Password,
		config.DbWriteSetting.Host,
		config.DbWriteSetting.Port,
		config.DbWriteSetting.Name,
		config.DbWriteSetting.Encoding,
	))

	if err != nil {
		log.Printf("创建DbWrite连接错误: %v\n", err)
		panic(err)
	}

	// 控制台打印SQL语句
	DbWrite.ShowSQL(true)
	DbWrite.Logger().SetLevel(core.LOG_WARNING)
}
