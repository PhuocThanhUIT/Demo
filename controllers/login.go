package controllers

import (
	"demo/helper"
	"demo/models"
	"demo/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/astaxie/beego"
)

// Operations about LoginController
type LoginController struct {
	beego.Controller
}

func (c *LoginController) URLMapping() {
	c.Mapping("Login", c.Login)
	c.Mapping("Google", c.Google)
	c.Mapping("Facebook", c.Facebook)
	c.Mapping("FacebookCallback", c.FacebookCallback)
	c.Mapping("GoogleCallback", c.GoogleCallback)
}

// Post ...
// @Title Login
// @Description Logs user into the system
// @Param	body	body 	models.LoginInput	true		"email and password"
// @Success 200 {string} login success
// @Success 200 {string} token and platform info
// @Failure 403 Wrong Email or Password
// @Failure 500 error message
// @router / [post]
func (u *LoginController) Login() {
	in := models.LoginInput{}
	log.Println(in)
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &in)
	if err != nil {
		fmt.Println(err.Error())
	}
	cus, err, statusCode := models.CheckLogin(in)
	if err == nil {
		if !reflect.DeepEqual(cus, models.User{}) {
			//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTkwMjE4NTEsImp0aSI6ImRhX3VNb1NIY0pYc1VqdVhYOGlxOTJWU0VpX2tBT01BYnJyRG03ZHdHcjFfeUxfZGZ3by1BVk02bjFuZXJLNFlPaWpQYUE9PSJ9.b8WflRxFPUCuc8b0BK0_HQyBvtBtbMX4QZMZbdNossQ"
			ts, err := models.CreateToken(cus.ID)
			if err != nil {
				//c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}
			tokens := map[string]string{
				"access_token":  ts.AccessToken,
				"refresh_token": ts.RefreshToken,
			}
			u.Data["json"] = tokens
			/*res["data"] = helper.LoginResponseData{
				Token:        token,
				RefreshToken: "refreshToken",
				CustomerInfo: cus,
			}*/
			u.Ctx.Output.Status = http.StatusOK
			//}
			//}
		} else {
			res["code"] = "failed"
			res["message"] = "Wrong Email or Password"
			u.Data["json"] = res
			u.Ctx.Output.Status = http.StatusForbidden
		}
	} else {
		res["code"] = "failed"
		res["message"] = err.Error()
		u.Data["json"] = res
		u.Ctx.Output.Status = statusCode
	}
	u.ServeJSON()
	//u.Data["json"] = token
	//u.ServeJSON()
	//cus.JSON(http.StatusOK, token)
}

// @Title Google
// @Description login google
// @Success 200 redirect to login google URL
// @router /google / [get]
func (u *LoginController) Google() {
	urlRes := helper.LoginInWithGoogle()
	u.Ctx.Redirect(http.StatusTemporaryRedirect, urlRes)
	//u.Data["json"] = res
	//u.ServeJSON()
}

// @Title Facebook
// @Description login facebook
// @Success 307 redirect to login facebook URL
// @router /facebook [get]
func (u *LoginController) Facebook() {
	urlRes := helper.LoginInWithFacebook()
	u.Ctx.Redirect(http.StatusTemporaryRedirect, urlRes)
	//u.Data["json"] = res
	//u.ServeJSON()
}

// @Title FacebookCallback
// @Description revice code from facebook and get infor user facebook
// @Param	state	query	string	false	"state from server"
// @Param	code	query	string	false	"code response from Facebook"
// @Success 200 {string} token when exists user
// @Success 307 redirect to register page
// @Failure 403 error message
// @Failure 500  error server
// @router /facebookcallback [get]
func (u *LoginController) FacebookCallback() {
	var state, code string
	_ = u.Ctx.Input.Bind(&state, "state")
	_ = u.Ctx.Input.Bind(&code, "code")
	res := make(map[string]interface{})
	resCallBack, err := helper.ProcessCallBackFacebook(state, code)
	if err != nil {
		res["code"] = "failed"
		res["message"] = err.Error()
		u.Ctx.Output.Status = http.StatusInternalServerError
	} else {
		pass := utils.RandStringBytes(16)
		input := models.User{
			BaseModel: models.BaseModel{
				//FirstName: resCallBack.FirstName,
				//LastName:  resCallBack.LastName,
				//IsActive:  true,

				//Avatar:    resCallBack.Picture["data"]["url"].(string),

				ID: resCallBack.ID,
			},
			Password: &pass,
			Email:    resCallBack.Email,
			//TypeSocial: "facebook",
		}
		isExist, token, err := helper.CheckUserSocialAndCreateToken(input)
		if err != nil {
			res["code"] = "forbidden"
			res["message"] = err.Error()
			u.Ctx.Output.Status = http.StatusForbidden
		} else {
			if isExist {
				res["code"] = "success"
				res["message"] = "login success"
				res["token"] = token
				u.Ctx.Output.Status = http.StatusOK
			} else {
				// REDIRECT TO REGISTER PAGE
				urlRes := helper.RedirectToRegisterPage(input)
				u.Ctx.Redirect(http.StatusTemporaryRedirect, urlRes)
			}
		}
	}
	u.Data["json"] = res
	u.ServeJSON()
}

// @Title GoogleCallback
// @Description revice code from google and get infor user google
// @Param	state	query	string	false	"state from server"
// @Param	code	query	string	false	"code response from google"
// @Success 200 {string} token when exists user
// @Success 307 redirect to register page
// @Failure 403 error message
// @Failure 500  error server
// @router /googlecallback [get]
func (u *LoginController) GoogleCallback() {
	var state, code string
	_ = u.Ctx.Input.Bind(&state, "state")
	_ = u.Ctx.Input.Bind(&code, "code")
	res := make(map[string]interface{})
	resCallBack, err := helper.ProcessCallBackGoogle(state, code)
	if err != nil {
		res["code"] = "failed"
		res["message"] = err.Error()
		u.Ctx.Output.Status = http.StatusInternalServerError
	} else {
		pass := utils.RandStringBytes(16)
		input := models.User{
			BaseModel: models.BaseModel{
				//FirstName: "",
				//LastName:  resCallBack.Name,
				//IsActive:  true,

				//Avatar:    resCallBack.Picture,
				ID: resCallBack.Sub,
			},
			Password: &pass,
			Email:    resCallBack.Email,
			//TypeSocial: "google",
		}
		isExist, token, err := helper.CheckUserSocialAndCreateToken(input)
		if err != nil {
			res["code"] = "forbidden"
			res["message"] = err.Error()
			u.Ctx.Output.Status = http.StatusForbidden
		} else {
			if isExist {
				res["code"] = "success"
				res["message"] = "login success"
				res["token"] = token
				u.Ctx.Output.Status = http.StatusOK
			} else {
				urlRes := helper.RedirectToRegisterPage(input)
				u.Ctx.Redirect(http.StatusTemporaryRedirect, urlRes)
			}
		}
	}
	u.Data["json"] = res
	u.ServeJSON()
}
