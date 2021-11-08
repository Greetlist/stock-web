package routers

import (
    "greetlist/stock-web/server/api"
    "github.com/gin-gonic/gin"
)

func InitOwnPreferStockOperationApiRouter(RouterGroup *gin.RouterGroup) {
    UserRouterGroup := RouterGroup.Group("user")
    UserRouterGroup.POST("/addPreferStock", api.AddPreferStock)
    UserRouterGroup.POST("/deletePreferStock", api.DeletePreferStock)
    UserRouterGroup.POST("/getPreferStock", api.GetPreferStock)
}
