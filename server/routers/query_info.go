package routers

import (
    "greetlist/stock-web/server/api"
    "github.com/gin-gonic/gin"
)

func InitQueryInfoRouter(RouterGroup *gin.RouterGroup) {
    StockRouterGroup := RouterGroup.Group("query")
    StockRouterGroup.GET("/getAllAlgoName", api.GetAllAlgoName)
}
