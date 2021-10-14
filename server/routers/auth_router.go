package routers

import (
    "greetlist/stock-web/server/api"
    "github.com/gin-gonic/gin"
)

func InitAuthApiRouter(RouterGroup *gin.RouterGroup) {
    StockRouterGroup := RouterGroup.Group("stock")
    StockRouterGroup.POST("/login", api.Login)
    StockRouterGroup.POST("/logout", api.Logout)
}
