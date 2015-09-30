package routers

import (
	"github.com/astaxie/beego"
	"myapp/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/category", &controllers.CategoryController{})
}
