package controllers

import (
    "github.com/astaxie/beego"
)

type BaseController struct {
    beego.Controller
}

type BaseResponse struct {

}

func (this *BaseController) SetResponse(data interface{}) {
    this.Data["Data"] = data
}
