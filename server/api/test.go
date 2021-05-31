package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// Test godoc
// @Summary Show a account
// @Description get string by ID
// @ID test
// @Produce  json
// @Success 200
// @Router /test [get]
func Test(c *gin.Context) {
    var msg struct {
        Name string `json:"name"`
        Age int `json:"age"`
        Msg string `json:"msg"`
    }
    msg.Name = "Greetlist"
    msg.Age = 26
    msg.Msg = "Fuck University"
    c.JSON(http.StatusOK, msg)
}
