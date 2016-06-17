// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"dsp.iclick/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/showup",
			beego.NSInclude(
				&controllers.ShowupController{},
			),
		),
		beego.NSNamespace("/click",
			beego.NSInclude(
				&controllers.ClickController{},
			),
		),
		beego.NSNamespace("/winner",
			beego.NSInclude(
				&controllers.WinnerController{},
			),
		),
		beego.NSNamespace("/bid",
			beego.NSInclude(
				&controllers.DspController{},
			),
		),
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
