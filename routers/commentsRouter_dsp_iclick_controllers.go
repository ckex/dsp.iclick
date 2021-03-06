package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["dsp.iclick/controllers:ClickController"] = append(beego.GlobalControllerRouter["dsp.iclick/controllers:ClickController"],
		beego.ControllerComments{
			"Click",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dsp.iclick/controllers:DspController"] = append(beego.GlobalControllerRouter["dsp.iclick/controllers:DspController"],
		beego.ControllerComments{
			"Head",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dsp.iclick/controllers:DspController"] = append(beego.GlobalControllerRouter["dsp.iclick/controllers:DspController"],
		beego.ControllerComments{
			"Head",
			`/`,
			[]string{"head"},
			nil})

	beego.GlobalControllerRouter["dsp.iclick/controllers:DspController"] = append(beego.GlobalControllerRouter["dsp.iclick/controllers:DspController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["dsp.iclick/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dsp.iclick/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["dsp.iclick/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dsp.iclick/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dsp.iclick/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dsp.iclick/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dsp.iclick/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dsp.iclick/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["dsp.iclick/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dsp.iclick/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["dsp.iclick/controllers:ShowupController"] = append(beego.GlobalControllerRouter["dsp.iclick/controllers:ShowupController"],
		beego.ControllerComments{
			"Showup",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dsp.iclick/controllers:UserController"] = append(beego.GlobalControllerRouter["dsp.iclick/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["dsp.iclick/controllers:UserController"] = append(beego.GlobalControllerRouter["dsp.iclick/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dsp.iclick/controllers:UserController"] = append(beego.GlobalControllerRouter["dsp.iclick/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dsp.iclick/controllers:UserController"] = append(beego.GlobalControllerRouter["dsp.iclick/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["dsp.iclick/controllers:UserController"] = append(beego.GlobalControllerRouter["dsp.iclick/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["dsp.iclick/controllers:UserController"] = append(beego.GlobalControllerRouter["dsp.iclick/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dsp.iclick/controllers:UserController"] = append(beego.GlobalControllerRouter["dsp.iclick/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dsp.iclick/controllers:WinnerController"] = append(beego.GlobalControllerRouter["dsp.iclick/controllers:WinnerController"],
		beego.ControllerComments{
			"Winning",
			`/`,
			[]string{"get"},
			nil})

}
