package controllers

import (
	"demo/models"
	"net/http"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

func (c *UserController) URLMapping() {
	c.Mapping("GetUser", c.GetUser)
}

// GetUser ...
// @Title GetUser
// @Description get Customer by id
// @Param	Authorization	header	string	true	"authen key"
// @Success 200 {object} models.Customer
// @Failure 403  forbidden
// @Failure 500  error server
// @Failure 401  Unauthorized
// @router / [get]
func (u *UserController) GetUser() {
	res := make(map[string]interface{})
	tokenAuth, err := models.ExtractTokenMetadata(u.Ctx.Request)
	if err != nil {
		res["code"] = "failed"
		res["message"] = "Wrong Email or Password"
		u.Data["json"] = res
		u.Ctx.Output.Status = http.StatusForbidden
		return
	}
	//userId, err := models.FetchAuth(tokenAuth)
	userId := tokenAuth.UserId
	cus, err := models.GetUserFromID(userId)
	if err != nil {

		return
	}
	u.Data["json"] = cus
	u.ServeJSON()
}
