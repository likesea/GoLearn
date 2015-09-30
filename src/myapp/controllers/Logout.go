package controllers

import (
	"github.com/astaxie/beego"
)

type LogoutController struct {
	beego.Controller
}

func (this *LogoutController) Get() {
	this.Ctx.SetCookie("uname", "", -1, "/")
	this.Ctx.SetCookie("pwd", "", -1, "/")
	this.Redirect("/", 301)
	return
}
