package server

import (
    "github.com/gin-gonic/gin"
    "greetlist/stock-web/server/api"
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"

    _ "greetlist/stock-web/server/docs"
)

func RunServer(bindAddr string, bindPort int) {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    r.GET("/test", api.Test)
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.Run(":8080")
}
