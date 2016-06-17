package controllers

import (
	"github.com/astaxie/beego"
	"bytes"
	logger "github.com/alecthomas/log4go"
)

type ShowupController struct {
	beego.Controller
}

// http://127.0.0.1:9090/v1/showup?bid=${AUCTION_ID}&imp=${AUCTION_IMP_ID}
// @Title showup
// @Description showup
// @Success 200
// @Failure 403 :objectId is empty
// @router / [get]
func (showup *ShowupController) Showup() {
	showup.Ctx.Request.ParseForm()
	params := showup.Ctx.Request.Form
	var buffer bytes.Buffer
	for key, value := range params {
		buffer.WriteString(",")
		buffer.WriteString(key)
		buffer.WriteString("=")
		buffer.WriteString(value[0])
	}
	logger.Info("%d%s", len(params), buffer.String())
}

