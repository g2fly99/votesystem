package main

import (
	"fmt"
	"votesystem/cache"
	"votesystem/etc"
	"votesystem/models"
	_ "votesystem/routers"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//load config infomation
	err := etc.InitConfig("conf/config.yml")
	if err != nil {
		errmsg := fmt.Sprintf("load config file err: %v", err)
		panic(errmsg)
	}

	logfile := "./logs/votesystem.log"
	if len(etc.LogPath()) != 0 {
		logfile = fmt.Sprintf("%v/votesystem.log", etc.LogPath())
	}
	fmt.Printf("init logfile:%v\n", logfile)
	err = logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"%s","maxdays":30}`, logfile))
	if err != nil {
		errmsg := fmt.Sprintf("create log file err: %v", err)
		panic(errmsg)
	}

	logs.SetLevel(etc.LogLevel())

	//mysql init
	err = models.InitMysql(etc.ConfDetail())
	if err != nil {
		logs.Warn("mysql connect err:%v", err)
	}

	cache.Init(etc.RedisAddr(), etc.RedisPasswd())
	beego.Run()
}
