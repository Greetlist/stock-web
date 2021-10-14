package middlemare

import (
    "github.com/gin-gonic/gin"
    redisMod "greetlist/stock-web/server/redis"
    "greetlist/stock-web/server/conf"
    "github.com/gomodule/redigo/redis"
    "net/http"
    "time"
)

func CheckLoginStatus() gin.HandlerFunc {
    return func(c *gin.Context) {
        redisConn, _ := <-redisMod.RedisConnPool
        cookie, err := c.Request.Cookie("session_id")
        if err != nil {
            c.Abort()
            return
        }
        if redisConn != nil {
            exists, _ := redis.Bool(redisConn.Do("EXISTS", cookie.Value))
            if !exists {
                c.JSON(http.StatusUnauthorized, gin.H{
                    "error": "Unauthorized",
                })
                c.Abort()
                return
            } else {
                loginTime, _ := redis.Int64(redisConn.Do("GET", cookie.Value))
                now := time.Now().Unix()
                if now - loginTime > conf.SessionExpireTime {
                    c.JSON(http.StatusUnauthorized, gin.H{
                        "error": "Session Expire",
                    })
                    c.Abort()
                    return
                } else {
                    c.Next()
                    return
                }
            }
            redisMod.RedisConnPool <- redisConn
        } else {
            c.Abort()
            return
        }
    }
}
