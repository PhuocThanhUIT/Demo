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

	beego.GlobalControllerRouter["demo/controllers:OtpController"] = append(beego.GlobalControllerRouter["demo/controllers:OtpController"],
		beego.ControllerComments{
			Method:           "GenerateOtp",
			Router:           `/generate-otp`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
	beego.GlobalControllerRouter["demo/controllers:OtpController"] = append(beego.GlobalControllerRouter["demo/controllers:OtpController"],
		beego.ControllerComments{
			Method:           "CheckOtp",
			Router:           `/check-otp`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
}
