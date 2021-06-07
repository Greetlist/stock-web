package middlemare

import (
    "github.com/gin-gonic/gin"
)

func SetCORSHeader() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Next()
    }
}
