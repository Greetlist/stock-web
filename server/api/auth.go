package api

import (
    "time"
    "net/http"
    "strconv"
    "crypto/sha256"
    "github.com/gin-gonic/gin"
    "github.com/gomodule/redigo/redis"

    "greetlist/stock-web/server/model"
    "greetlist/stock-web/server/util"
    "greetlist/stock-web/server/database"
    redisMod "greetlist/stock-web/server/redis"
)

// Login godoc
// @Summary User Login
// @Description User Login
// @ID login
// @Accept json
// @Produce json
// @Param request_json body model.LoginRequest true "login"
// @Success 200 {object} model.LoginResponse
// @Router /api/auth/login [post]
func Login(context *gin.Context) {
    request := model.LoginRequest{}
    if err := context.BindJSON(&request); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var response model.LoginResponse
    if util.ArgsContainIllegueWord(request) || !verifyAccountPasswd(request.Account, request.Password) {
        response.LoginSucc = false
        context.JSON(http.StatusOK, response)
        return
    }

    cookieValue := genRandomCookieValue(request.Account, request.Password)
    if !setCookieToRedis(cookieValue) {
        response.LoginSucc = false
    } else {
        response.LoginSucc = true
        cookie := &http.Cookie{
            Name: "session_id",
            Value: cookieValue,
            Path: "/",
            HttpOnly: true,
        }
        http.SetCookie(context.Writer, cookie)
    }
    context.JSON(http.StatusOK, response)
}

func verifyAccountPasswd(account, passwd string) bool {
    dbConn, _ := <-database.DBConnPool
    if dbConn == nil {
        return false
    }
    var fundAccount model.FundAccount
    dbConn.Where("account = ?", account).First(&fundAccount)
    database.DBConnPool <- dbConn
    if fundAccount.Password != passwd {
        return false
    }
    return true
}

func genRandomCookieValue(account, passwd string) string {
    now := time.Now().Unix()
    var lowMask byte = 15
    shaBytes := sha256.Sum256([]byte(strconv.FormatInt(now, 10) + account + passwd))
    var newBytesArr [len(shaBytes)*2]byte
    for i := 0; i < len(shaBytes); i++ {
        var high byte = (shaBytes[i] >> 4)
        var low byte = (shaBytes[i] & lowMask)
        high += offset(high)
        low += offset(low)
        newBytesArr[i*2] = high
        newBytesArr[i*2+1] = low
    }
    res := string(newBytesArr[:])
    return res
}

func offset(b byte) byte {
    if b > 9 {
        return 87
    }
    return 48
}

func setCookieToRedis(cookie string) bool {
    redisConn, _ := <-redisMod.RedisConnPool
    if redisConn == nil {
        return false
    }
    res, _ := redis.String(redisConn.Do("SET", cookie, string(time.Now().Unix())))
    if res != "OK" {
        return false
    }
    return true
}

// Logout godoc
// @Summary User Logout
// @Description User Logout
// @ID logout
// @Accept json
// @Produce json
// @Param request_json body model.LogoutRequest true "Logout"
// @Success 200 {object} model.LogoutResponse
// @Router /api/auth/logout [post]
func Logout(context *gin.Context) {
    request := model.LogoutRequest{}
    if err := context.BindJSON(&request); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    var response model.LogoutResponse
    context.JSON(http.StatusOK, response)
}
