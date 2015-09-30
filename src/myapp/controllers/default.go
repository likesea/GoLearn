package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["True"] = true
	c.Data["html"] = "<div>hello</div>"
	c.Data["test"] = "<div>hesafsafllo</div>"
	c.TplNames = "index.tpl"
}
