package api

import (
    "io/ioutil"
    "fmt"
    "net/http"
    "strings"
    _ "time"
    _ "os"
    _ "errors"
    "github.com/gin-gonic/gin"
    "greetlist/stock-web/server/model"
    "greetlist/stock-web/server/conf"
    _ "greetlist/stock-web/server/util"
)

// GetAllAlgoName godoc
// @Summary Query All Algo Name
// @Description Return All Algo Name
// @ID getAllAlgoName
// @Produce json
// @Success 200 {object} model.AllAlgoNameResponse
// @Router /api/query/getAllAlgoName [get]
func GetAllAlgoName(context *gin.Context) {
    files, err := ioutil.ReadDir(conf.StockPredictionBaseDir)
    var response model.AllAlgoNameResponse
    if err != nil {
        fmt.Printf("Read Dir Error: %s", err)
        return
    }
    for _, file := range files {
        if strings.HasPrefix(file.Name(), "predictions") {
            //algoName := strings.Split(file.Name(), "_")[1]
            response.AlgoNameList = append(response.AlgoNameList, file.Name())
        }
    }
    context.JSON(http.StatusOK, response)
}
