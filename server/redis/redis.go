package redisMod

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
    "greetlist/stock-web/server/conf"
)

var RedisConnPool chan redis.Conn = make(chan redis.Conn, conf.RedisConnectionNum)
func InitRedisConnectionPool() {
    for i := 0; i < conf.RedisConnectionNum; i++ {
        conn, err := redis.Dial(conf.RedisProtocol, conf.RedisAddr)
        if err != nil {
            fmt.Printf("Connect to Redis Server Error: %s.\n", err)
            continue
        }
        RedisConnPool <- conn
    }
}
