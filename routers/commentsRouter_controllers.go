package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["demo/controllers:UserController"] = append(beego.GlobalControllerRouter["demo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetUser",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["demo/controllers:LoginController"] = append(beego.GlobalControllerRouter["demo/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["demo/controllers:LoginController"] = append(beego.GlobalControllerRouter["demo/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "Facebook",
			Router:           `/facebook`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["demo/controllers:LoginController"] = append(beego.GlobalControllerRouter["demo/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "FacebookCallback",
			Router:           `/facebookcallback`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["demo/controllers:LoginController"] = append(beego.GlobalControllerRouter["demo/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "Google",
			Router:           `/google`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["demo/controllers:LoginController"] = append(beego.GlobalControllerRouter["demo/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "GoogleCallback",
			Router:           `/googlecallback`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
