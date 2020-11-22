package controllers

import (
	"demo/helper"
	"encoding/json"
	"log"
	"net/http"

	"github.com/astaxie/beego"
)

// OtpController operations for Otp
type OtpController struct {
	beego.Controller
}

// URLMapping ...
func (c *OtpController) URLMapping() {
	c.Mapping("GenerateOtp", c.GenerateOtp)
	c.Mapping("CheckOtp", c.CheckOtp)
}

// Generate Otp ...
// @Title Generate Otp from phone number
// @Description Generater Otp
// @Param	body	body 	phone_number	true		"Phone number"
// @Success 200 {object} Check OTP in your phone
// @Failure 403
// @Failure 400
// @router /generate-otp [post]
func (c *OtpController) GenerateOtp() {
	res := helper.ResponseDefault{}
	var in helper.OtpInput
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &in)
	log.Println(in)
	err, statusCode := helper.GenerateOtp(in.PhoneNumber)
	if err == nil {
		res = helper.ResponseDefault{
			Code:    "success",
			Message: "Check OTP in your phone",
		}
	} else {
		c.Ctx.Output.Status = statusCode
		res = helper.ResponseDefault{
			Code:    "error",
			Message: err.Error(),
		}
	}
	c.Data["json"] = res
	c.ServeJSON()
}

// GetAll ...
// @Title Check Otp
// @Description get Meta
// @Param	body	body 	helper.OtpInput	true		"Phone Number and Otp"
// @Success 200 {object} []models.Otp
// @Failure 403
// @Failure 403
// @router /check-otp [post]

func (c *OtpController) CheckOtp() {
	res := helper.ResponseDefault{}
	var in helper.OtpInput
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &in)
	data, err, statusCode := helper.CheckOtp(in)
	c.Ctx.Output.Status = statusCode
	if err == nil {
		if statusCode == http.StatusForbidden {
			res = helper.ResponseDefault{
				Code:    "failed",
				Message: "Something when wrong with your OTP",
				Data:    data,
			}
		} else {
			res = helper.ResponseDefault{
				Code:    "success",
				Message: "",
				Data:    data,
			}
		}
	} else {
		res = helper.ResponseDefault{
			Code:    "error",
			Message: err.Error(),
			Data:    nil,
		}
	}
	c.Data["json"] = res
	c.ServeJSON()
}
