package server

import (
    initModel "greetlist/stock-web/server/init"
    redisMod "greetlist/stock-web/server/redis"
    "greetlist/stock-web/server/database"
    "strconv"
)

func RunServer(bindAddr string, bindPort int64) {
    router := initModel.InitRouterAndMiddleware()
    database.InitDataBase()
    redisMod.InitRedisConnectionPool()
    router.Run(bindAddr + ":" + strconv.FormatInt(bindPort, 10))
}
