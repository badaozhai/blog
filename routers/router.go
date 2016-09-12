package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/search", &controllers.MainController{}, "*:Search")
	beego.Router("/json-diff", &controllers.JsonDiffController{})
}
