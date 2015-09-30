package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["IsHome"] = true
	beego.Trace("asfd")
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplNames = "home.html"
}
