package service

import (
	"graphqldemo/models"
	"log"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter"
	_ "github.com/go-sql-driver/mysql"
)

var Enforcer *casbin.Enforcer

// 初始化casbin
func CasbinSetup() {
	a := xormadapter.NewAdapterByEngine(models.DbWrite)
	// a := xormadapter.NewAdapter("mysql", "rbac:rbac168!@tcp(119.29.167.27:18168)/rbac_db?charset=utf8")
	// if err != nil {
	// 	log.Printf("连接数据库错误: %v", err)
	// 	return
	// }
	e, err := casbin.NewEnforcer("conf/rbac_models.conf", a)
	if err != nil {
		log.Printf("初始化casbin错误: %v", err)
		return
	}

	Enforcer = e
}
