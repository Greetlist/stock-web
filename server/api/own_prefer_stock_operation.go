package api

import (
    "net/http"
    "github.com/gin-gonic/gin"

    "greetlist/stock-web/server/model"
    "greetlist/stock-web/server/database"
)

// AddPreferStock godoc
// @Summary User AddPreferStock
// @Description User AddPreferStock
// @ID addPreferStock
// @Accept json
// @Produce json
// @Param request_json body model.AddPreferStockRequest true "login"
// @Success 200 {object} model.AddPreferStockResponse
// @Router /user/add_prefer_stock [post]
func AddPreferStock(context *gin.Context) {
    request := model.AddPreferStockRequest{}
    if err := context.BindJSON(&request); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var response model.AddPreferStockResponse
    context.JSON(http.StatusOK, response)
}

func doAddPreferStock() bool {
    dbConn, _ := <-database.DBConnPool
    if dbConn == nil {
        return false
    }
    database.DBConnPool <- dbConn
    return true
}
