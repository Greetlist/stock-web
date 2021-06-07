package routers

import (
    "greetlist/stock-web/server/api"
    "github.com/gin-gonic/gin"
)

func InitStockApiRouter(RouterGroup *gin.RouterGroup) {
    StockRouterGroup := RouterGroup.Group("stock")
    StockRouterGroup.GET("/getDailyStockData", api.GetDailyStockData)
    StockRouterGroup.POST("/getQueryStockData", api.GetQueryStockData)
}
