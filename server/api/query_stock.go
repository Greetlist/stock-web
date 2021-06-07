package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "greetlist/stock-web/server/model"
    "greetlist/stock-web/server/conf"
    "greetlist/stock-web/server/util"
)

// GetDailyStockData godoc
// @Summary Query Daily Stock Computed Data
// @Description Return Strategy Daily Computed Data
// @ID queryDailyStockData
// @Accept json
// @Produce json
// @Success 200
// @Router /api/stock/getDailyStockData [get]
func GetDailyStockData(context *gin.Context) {
}

// GetQueryStockData godoc
// @Summary Query Sepcific Stock Computed Data
// @Description Return Strategy Sepcific Computed Data
// @ID getQueryStockData
// @Accept json
// @Produce json
// @Success 200 {object} model.GetQueryStockDataResponse
// @Router /api/stock/getQueryStockData [post]
func GetQueryStockData(context *gin.Context) {
    request := model.GetQueryStockDataRequest{}
    //request := make(map[string]interface{})
    if err := context.BindJSON(&request); err != nil {
	    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    var response model.GetQueryStockDataResponse
    for _, stockCode := range(request.StockList) {
        var stockDataItem model.StockDataItem
        stockDataItem.StockCode = stockCode
        exchangeId := ".XSHE"
        if stockCode[0] == '6' {
            exchangeId = ".XSHG"
        }
        stockFilePath := conf.StockResultBaseDir + stockCode + exchangeId + "/result_ma.csv"
        csvReader := util.CsvReader{}
        csvReader.ReadCsv(stockFilePath)
        dataLen := len(csvReader.Data)
        startIndex := dataLen / 7
        stockDataItem.Records = csvReader.Data[startIndex:]
        response.StockDatas = append(response.StockDatas, stockDataItem)
    }
    context.JSON(http.StatusOK, response)
}
