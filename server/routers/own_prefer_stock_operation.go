package routers

import (
    "greetlist/stock-web/server/api"
    "github.com/gin-gonic/gin"
)

func InitOwnPreferStockOperationApiRouter(RouterGroup *gin.RouterGroup) {
    RouterGroup.POST("/user/add_prefer_stock", api.AddPreferStock)
    //RouterGroup.POST("/user/delete_prefer_stock", api.DeletePreferStock)
    //RouterGroup.POST("/user/get_prefer_stock", api.GetPreferStock)
}
