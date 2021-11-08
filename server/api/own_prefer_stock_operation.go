package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "greetlist/stock-web/server/model"
    "greetlist/stock-web/server/errcode"
    "greetlist/stock-web/server/util"
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
// @Router /api/user/addPreferStock [post]
func AddPreferStock(context *gin.Context) {
    request := model.AddPreferStockRequest{}
    if err := context.BindJSON(&request); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var response model.AddPreferStockResponse
    response.IsSuccess, response.ErrMsg = doAddPreferStock(request.Account, request.StockCode)
    context.JSON(http.StatusOK, response)
}

func doAddPreferStock(account, stockCode string) (bool, string) {
    dbConn, _ := <-database.DBConnPool
    defer func(dbConn *gorm.DB) {database.DBConnPool <- dbConn}(dbConn)
    if dbConn == nil {
        return false, errcode.DBConnError
    }
    stockInfoList := util.ReadStockBasicInfo(stockCode)
    if len(stockInfoList) == 0 {
        return false, errcode.StockNotFound
    }
    stockInfo := stockInfoList[0]
    var followStock model.FollowStock
    followStock.Account = account
    followStock.StockCode = stockCode
    followStock.Exchange = stockInfo.Exchange
    followStock.Name = stockInfo.StockName
    followStock.Industry = stockInfo.StockIndustry
    if res := dbConn.Create(&followStock); res.Error != nil {
        return false, errcode.DBExecuteError
    }
    return true, ""
}

// DeletePreferStock godoc
// @Summary User DeletePreferStock
// @Description User DeletePreferStock
// @ID deletePreferStock
// @Accept json
// @Produce json
// @Param request_json body model.DeletePreferStockRequest true "login"
// @Success 200 {object} model.DeletePreferStockResponse
// @Router /api/user/deletePreferStock [post]
func DeletePreferStock(context *gin.Context) {
    request := model.DeletePreferStockRequest{}
    if err := context.BindJSON(&request); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var response model.DeletePreferStockResponse
    response.IsSuccess, response.ErrMsg = doDeletePreferStock(request.Account, request.StockCode)
    context.JSON(http.StatusOK, response)
}

func doDeletePreferStock(account, stockCode string) (bool, string) {
    dbConn, _ := <-database.DBConnPool
    defer func(dbConn *gorm.DB) {database.DBConnPool <- dbConn}(dbConn)
    if dbConn == nil {
        return false, errcode.DBConnError
    }
    var followStock model.FollowStock
    followStock.Account = account
    followStock.StockCode = stockCode
    if res := dbConn.Delete(&followStock); res.Error != nil {
        return false, errcode.DBExecuteError
    }
    return true, ""
}

// GetPreferStock godoc
// @Summary User GetPreferStock
// @Description User GetPreferStock
// @ID getPreferStock
// @Accept json
// @Produce json
// @Param request_json body model.GetPreferStockRequest true "login"
// @Success 200 {object} model.GetPreferStockResponse
// @Router /api/user/getPreferStock [post]
func GetPreferStock(context *gin.Context) {
    request := model.GetPreferStockRequest{}
    if err := context.BindJSON(&request); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var response model.GetPreferStockResponse
    response.IsSuccess, response.ErrMsg, response.FollowStockList = doGetPreferStock(request.Account)
    context.JSON(http.StatusOK, response)
}

func doGetPreferStock(account string) (bool, string, []model.FollowStock) {
    dbConn, _ := <-database.DBConnPool
    var followStockList []model.FollowStock
    defer func(dbConn *gorm.DB) {database.DBConnPool <- dbConn}(dbConn)
    if dbConn == nil {
        return false, errcode.DBConnError, followStockList
    }
    if res := dbConn.Where("account = ?", account).Find(&followStockList); res.Error != nil {
        return false, errcode.DBExecuteError, followStockList
    }
    return true, "", followStockList
}
