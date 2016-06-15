package main

import (
	_ "dsp.iclick/docs"
	_ "dsp.iclick/routers"
	logger "github.com/alecthomas/log4go"
	"github.com/astaxie/beego"
)

func init() {
	logger.LoadConfiguration("./conf/log4go.xml")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
