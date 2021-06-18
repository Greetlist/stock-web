package api

import (
    "net/http"
    "io/ioutil"
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
// @Param request_json body model.GetQueryStockDataRequest true "Query Stock List"
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
        stockFilePath := conf.StockResultBaseDir + stockCode + "/result_ma.csv"
        csvReader := util.CsvReader{}
        csvReader.ReadCsv(stockFilePath)
        dataLen := len(csvReader.Data)
        startIndex := dataLen - dataLen / 7
        stockDataItem.Records = csvReader.Data[startIndex:]
        response.StockDatas = append(response.StockDatas, stockDataItem)
    }
    context.JSON(http.StatusOK, response)
}

// GetAllStockCode godoc
// @Summary Query All Stock Code
// @Description Return All Stock Code
// @ID getAllStockCode
// @Produce json
// @Success 200 {object} model.GetAllStockCodeResponse
// @Router /api/stock/getAllStockCode [get]
func GetAllStockCode(context *gin.Context) {
    var stockList []string
    fileInfoList, _ := ioutil.ReadDir(conf.StockStorageBaseDir)
    for _, fileItem := range(fileInfoList) {
        if fileItem.IsDir() {
            stockList = append(stockList, fileItem.Name())
        }
    }
    response := model.GetAllStockCodeResponse{StockList: stockList}
    context.JSON(http.StatusOK, response)
}
