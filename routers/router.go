package routers

import (
	"github.com/mtgterao/derby-mas-memo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
