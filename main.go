package main

import (
	"demo/conf"
	_ "demo/routers"
	"demo/utils"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/caarlos0/env"
)

func main() {
	config := conf.AppConfig{}
	_ = env.Parse(&config)
	log.Println(&config)
	utils.GetConstant(config)
	_ = beego.LoadAppConfig("ini", "conf/app.conf")
	log.Println(beego.BConfig.RunMode)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	err := utils.AutoMigration()
	if err != nil {
		fmt.Println(err.Error())
	}
	beego.Run()
}
