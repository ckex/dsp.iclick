package controllers

import "github.com/astaxie/beego"

type ClickController struct {
	beego.Controller
}

// http://127.0.0.1:9090/v1/click?adxclick=${CLICK_URL:URLENCODE}&imp=${AUCTION_IMP_ID}
// Golang Redirect https://github.com/jsix/gouri
// @Title click
// @Description click
// @Success 200
// @Failure 403 :objectId is empty
// @router / [get]
func (click *ClickController) Click() {
	click.Ctx.Request.ParseForm()
	params := click.Ctx.Request.Form
	targetUrl := params["adxclick"][0]
	click.Redirect(targetUrl, 302)
}
