package controllers

import (
	"github.com/astaxie/beego"
	logger "github.com/alecthomas/log4go"
	"bytes"
)

type WinnerController struct {
	beego.Controller
}

// http://127.0.0.1:9090/v1/winner?p=1&bid=2&rid=3&imp=4&seat=5&adid=6&cur=CNY
// @Title winning
// @Description winner notice
// @Success 200
// @Failure 403 :objectId is empty
// @router / [get]
func (winner *WinnerController) Winning() {
	winner.Ctx.Request.ParseForm()
	params := winner.Ctx.Request.Form
	var buffer bytes.Buffer
	for key, value := range params {
		buffer.WriteString(",")
		buffer.WriteString(key)
		buffer.WriteString("=")
		buffer.WriteString(value[0])
	}
	logger.Info("%d%s", len(params), buffer.String())
	winner.Ctx.WriteString("success")
}