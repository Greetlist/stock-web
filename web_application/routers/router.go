package routers

import (
    "github.com/beego/beego/v2/server/web"
    "greetlist/stock-web/web_application/controllers"
)

func init() {
    web.Router("/", &controllers.MainController{}, "*:ShowHomePage")
    web.Router("/test", &controllers.MainController{}, "*:Test")
}
