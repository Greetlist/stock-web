package routers

import (
    "greetlist/stock-web/server/api"
    "github.com/gin-gonic/gin"
)

func InitAuthApiRouter(RouterGroup *gin.RouterGroup) {
    RouterGroup.POST("/login", api.Login)
    RouterGroup.POST("/logout", api.Logout)
}
