package routers

import (
    "github.com/beego/beego/v2/server/web"
    "greetlist/stock-web/web_application/controllers"
)

func init() {
    web.Router("/", &controllers.MainController{}, "*:ShowHomePage")
    web.Router("/test", &controllers.MainController{}, "*:Test")

    web.Router("/get_stock_raw_data", &controllers.StockController{}, "post:GetStockRawData")
}
