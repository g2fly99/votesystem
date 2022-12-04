package models

import (
	"fmt"
	"time"
	"votesystem/etc"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
)

type TimeModel struct {
	Created time.Time `orm:"auto_now_add; type(datetime)"`
	Updated time.Time `orm:"auto_now; type(datetime)"`
}

func registerAllModels() {
	registerCondidate()
	registerAuthority()
	registerVotes()
	registerRole()
	registerUser()
}

// 初始化mysql
func InitMysql(config etc.ConfigT) error {

	orm.Debug = config.Mysql.Debug

	registerAllModels()

	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {

		return err
	}

	alias := "default"
	auth := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&loc=Local",
		config.Mysql.UserName, config.Mysql.Password, config.Mysql.IpAddr, config.Mysql.DbName)

	logs.Debug(auth)
	err = orm.RegisterDataBase(alias, "mysql", auth)
	if err != nil {
		logs.Error("register dataBase failed:%v,auth:%v", err, auth)
		return err
	}

	err = orm.RunSyncdb(alias, false, true)
	if err != nil {
		logs.Error("RunSyncdb failed:%v", err)
		return err
	}

	//init base roles
	InitBaseRole()

	return nil
}
