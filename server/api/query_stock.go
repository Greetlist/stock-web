package api

import (
    "net/http"
    "io/ioutil"
    "strings"
    "fmt"
    "time"
    _ "os"
    "path"
    "errors"
    "github.com/gin-gonic/gin"
    "greetlist/stock-web/server/model"
    "greetlist/stock-web/server/conf"
    "greetlist/stock-web/server/util"
)

// QueryStockData godoc
// @Summary Query Sepcific Stock Computed Data
// @Description Return Strategy Sepcific Computed Data
// @ID queryStockData
// @Accept json
// @Produce json
// @Param request_json body model.QueryStockDataRequest true "Query Stock List"
// @Success 200 {object} model.QueryStockDataResponse
// @Router /api/stock/queryStockData [post]
func QueryStockData(context *gin.Context) {
    request := model.QueryStockDataRequest{}
    if err := context.BindJSON(&request); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    var response model.QueryStockDataResponse
    todayDateStr := time.Now().Format("2006-01-02")
    for _, stockCode := range(request.StockList) {
        if data, err := extractStockRawData(stockCode, todayDateStr, request.QueryDataLen); err == nil {
            response.StockTotalDatas = append(response.StockTotalDatas, data)
        }
    }
    context.JSON(http.StatusOK, response)
}

func extractStockRawData(stockCode, endDate string, qLen int) (model.SingleStockData, error) {
    var singleStockData model.SingleStockData
    stockInfoList := util.ReadStockBasicInfo(stockCode)
    if len(stockInfoList) == 0 {
        fmt.Printf("Get Specific stock info error.\n")
        return singleStockData, errors.New("Get Specific stock info error")
    }
    singleStockData.StockInfo = stockInfoList[0]
    periodList := []string{"day", "week", "month"}
    for _, period := range(periodList) {
        var curPeriodData model.StockSinglePeriodData
        curPeriodData.MarketData = util.ReadMarketRawData(stockCode, period, endDate, qLen)
        curPeriodData.FactorData = util.ReadFactorRawData(stockCode, period, endDate, qLen)
        curPeriodData.PeriodType = period
        singleStockData.StockRawDatas = append(singleStockData.StockRawDatas, curPeriodData)
    }
    return singleStockData, nil
}


// GetAllStockCode godoc
// @Summary Query All Stock Code
// @Description Return All Stock Code
// @ID getAllStockCode
// @Produce json
// @Success 200 {object} model.GetAllStockCodeResponse
// @Router /api/stock/getAllStockCode [get]
func GetAllStockCode(context *gin.Context) {
    stockInfoList := util.ReadStockBasicInfo("all")
    response := model.GetAllStockCodeResponse{StockInfoList: stockInfoList}
    context.JSON(http.StatusOK, response)
}

// GetAllStockCode godoc
// @Summary Query Daily Calc Stock Data
// @Description Return All Daily Calc Stock Data
// @ID getDailyCalcStockData
// @Produce json
// @Success 200 {object} model.MarketRawData
// @Router /api/stock/getDailyCalcStockData [get]
func GetDailyCalcStockData(context *gin.Context) {
    //todayStr := strings.ReplaceAll(util.GetTodayStr(), "-", "/")
    //dailyFilePath := path.Join(conf.DailyResultBaseDir, todayStr, "/ma_statistic_result.csv")

    //type result struct {
        //StockCode string `csv:"StockCode"`
        //CheckFlag bool `csv:"CheckFlag"`
    //}
    //f, err := os.OpenFile(dailyFilePath, os.O_RDWR, os.ModePerm)
    //if err != nil {
        //fmt.Printf("Open File %s Error : %s.\n", dailyFilePath, err)
    //}
    //defer f.Close()
    //var resultList []*result
    //if err := gocsv.UnmarshalFile(f, &resultList); err != nil {
        //fmt.Printf("Unmarshal csv file error : %s.\n", err)
	//}
    //var dailyStockCodeList []string
    //for _, value := range(resultList) {
        //if (*value).CheckFlag {
            //dailyStockCodeList = append(dailyStockCodeList, (*value).StockCode)
        //}
    //}

    //var stockDatas []model.StockDataItem
    //for _, stockCode := range(dailyStockCodeList) {
        //var curItem model.StockDataItem
        //curItem.StockCode = stockCode
        //stockCsvFilePath := path.Join(conf.StockResultBaseDir, stockCode, "/result_ma.csv")
        //curItem.Records = extractSingleStockDataFromCsv(stockCsvFilePath, 0)
        //stockDatas = append(stockDatas, curItem)
    //}

    //response := model.GetDailyCalcStockDataResponse{stockDatas}
    //context.JSON(http.StatusOK, response)
}

// GetRecommandStockPrediction godoc
// @Summary Query Recommand Stock Prediction Data
// @Description Return Recommand Stock Prediction Data
// @ID getRecommandStockPrediction
// @Accept json
// @Produce json
// @Param request_json body model.GetRecommandStockRequest true "Query Date"
// @Success 200 {object} model.GetRecommandStockResponse
// @Router /api/stock/getRecommandStockPrediction [post]
func GetRecommandStockPrediction(context *gin.Context) {
    request := model.GetRecommandStockRequest{}
    if err := context.BindJSON(&request); err != nil {
	    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    var response model.GetRecommandStockResponse

    predictionDir := path.Join(conf.StockPredictionBaseDir, strings.ReplaceAll(request.QueryDateString, "-", "/"))
    fileInfoList, _ := ioutil.ReadDir(predictionDir)
    for _, fileItem := range(fileInfoList) {
        if !strings.HasSuffix(fileItem.Name(), "compared.csv") && strings.HasSuffix(fileItem.Name(), ".csv") {
            stockCode := strings.ReplaceAll(fileItem.Name(), ".csv", "")
            var curPredictItem model.StockPredictItem
            fmt.Printf("%s\n", path.Join(predictionDir, fileItem.Name()))
            curPredictItem.StockInfo = util.ReadStockBasicInfo(stockCode)[0]
            curPredictItem.PredictionRecord = util.ReadPredictionData(path.Join(predictionDir, fileItem.Name()), true)
            curPredictItem.ShowMsg = util.ReadPredictionMsg(path.Join(predictionDir, "stock_Kline", strings.ReplaceAll(fileItem.Name(), "csv", "txt")))
            if curRawData, err := extractStockRawData(stockCode, request.QueryDateString, request.QueryDataLen); err == nil {
                response.StockRawData = append(response.StockRawData, curRawData)
                response.StockPredictDatas = append(response.StockPredictDatas, curPredictItem)
            }
        }
    }
    context.JSON(http.StatusOK, response)
}

// GetTotalMarketIndexData godoc
// @Summary Query Sepcific All Index Computed Data
// @Description Return Strategy Sepcific Computed Data
// @ID getTotalMarketIndexData
// @Accept json
// @Produce json
// @Param request_json body model.GetTotalMarketIndexDataRequest true "Query Index "
// @Success 200 {object} model.GetTotalMarketIndexDataResponse
// @Router /api/stock/getTotalMarketIndexData [post]
func GetTotalMarketIndexData(context *gin.Context) {
    request := model.GetTotalMarketIndexDataRequest{}
    if err := context.BindJSON(&request); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    var response model.GetTotalMarketIndexDataResponse
    indexFileDir := path.Join(conf.StockPredictionBaseDir, strings.ReplaceAll(request.QueryDateString, "-", "/"), "my_index_Kline")
    fmt.Printf("%s\n", path.Join(indexFileDir, "all.csv"))
    response.IndexRawData = util.ReadPredictionData(path.Join(indexFileDir, "all.csv"), false)
    response.IndexPredData = util.ReadPredictionData(path.Join(indexFileDir, "all.csv"), true)
    response.ShowMsg = util.ReadPredictionMsg(path.Join(indexFileDir, "all.txt"))
    context.JSON(http.StatusOK, response)
}
