package routers

import (
	"github.com/kooksee/kood/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 首页
	beego.Router("/", &controllers.MainController{})
	beego.Router("/tag", &controllers.MainController{})
	beego.Router("/category", &controllers.MainController{})
	beego.Router("/metadata", &controllers.MainController{})
	beego.Router("/object", &controllers.MainController{})
	beego.Router("/fs", &controllers.MainController{})
	beego.Router("/admin", &controllers.MainController{})
}
