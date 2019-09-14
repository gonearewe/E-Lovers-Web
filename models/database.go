package models

import (
	"fmt"

	"github.com/gonearewe/E-Lovers-Web/tools"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

/*
1. RegisterDriver
2. RegisterDataBase
3. RegisterModel
*/
func init() {
	log := tools.NewLogger()
	defer log.Close()

	driverName := beego.AppConfig.String("driverName")
	fmt.Println(beego.AppConfig.String("appname"))
	orm.RegisterDriver(driverName, orm.DRSqlite)

	//数据库连接
	dbname := beego.AppConfig.String("dbname")

	err := orm.RegisterDataBase("default", driverName, dbname)
	if err != nil {
		log.Critical("连接数据库出错")
		return
	}
	log.Informational("连接数据库成功")

	orm.RegisterModel(new(user))

	// create table if not exist
	orm.RunSyncdb("default", false, true)
}
