package init

import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

    "greetlist/stock-web/server/routers"
    _ "greetlist/stock-web/server/docs"
    midware "greetlist/stock-web/server/middleware"
)

func InitRouterAndMiddleware() *gin.Engine {
    //全局性设置
    Router := gin.New()
    Router.Use(gin.Logger())
    Router.Use(gin.Recovery())
    //跨域请求
    Router.Use(midware.SetCORSHeader())
    //API Doc
    Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    //前端展示路由,需要跟vue集成
    //WebInterfaceGroup := Router.Group("")

    /*
      功能API路由设置
      API 相关组
    */
    ApiRouterGroup := Router.Group("api")
    routers.InitStockApiRouter(ApiRouterGroup)

    //登陆路由
    AuthRouterGroup := Router.Group("auth")
    //AuthRouterGroup.Use(midware.CheckLoginStatus())
    routers.InitAuthApiRouter(AuthRouterGroup)

    //用户绑定操作，需要验证登陆Cookie


    return Router
}
