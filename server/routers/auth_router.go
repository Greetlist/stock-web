package routers

import (
    "greetlist/stock-web/server/api"
    "github.com/gin-gonic/gin"
)

func InitAuthApiRouter(RouterGroup *gin.RouterGroup) {
    AuthRouterGroup := RouterGroup.Group("auth")
    AuthRouterGroup.POST("/login", api.Login)
    AuthRouterGroup.POST("/logout", api.Logout)
}
