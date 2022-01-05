package api

import (
    "net/http"
    "io/ioutil"
    "strings"
    "fmt"
    "time"
    "os"
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

    if request.StockCode == "" {
        predictionDir := path.Join(conf.StockPredictionBaseDir, strings.ReplaceAll(request.QueryDateString, "-", "/"))
        recommendListFile := path.Join(predictionDir, "summary", "main.txt")
        recommendStockList := []string{"000001.SH", "399001.SZ", "399006.SZ"}
        recommendStockList = append(recommendStockList, util.ReadPredictionFileInOrder(recommendListFile)...)
        for _, stockCode := range(recommendStockList) {
            var curPredictItem model.StockPredictItem
            curPredictItem.StockInfo = util.ReadStockBasicInfo(stockCode)[0]
            financeType := util.GetCodeFinanceType(stockCode)
            curPredictItem.PredictionRecord = util.ReadPredictionData(path.Join(predictionDir, financeType + "_Kline", stockCode + ".csv"), true)
            curPredictItem.ShowMsg = util.ReadPredictionMsg(path.Join(predictionDir, financeType + "_Kline", stockCode + ".txt"))
            if curRawData, err := extractStockRawData(stockCode, request.QueryDateString, request.QueryDataLen); err == nil {
                response.StockRawDatas = append(response.StockRawDatas, curRawData)
                response.StockPredictDatas = append(response.StockPredictDatas, curPredictItem)
            }
        }
    } else {
        todayDateStr, lastTradingDateStr := util.LastTradingDayDirStr()
        fmt.Printf("%s %s\n", todayDateStr, lastTradingDateStr)
        if _, err := os.Stat(path.Join(conf.StockPredictionBaseDir, todayDateStr)); err == nil {
            lastTradingDateStr = todayDateStr
        }
        fmt.Printf("%s\n", path.Join(conf.StockPredictionBaseDir, lastTradingDateStr))
        fmt.Printf("%s %s\n", todayDateStr, lastTradingDateStr)
        singleStockPredFilePath := path.Join(conf.StockPredictionBaseDir, lastTradingDateStr, "stock_Kline", request.StockCode+".csv")
        var curPredictItem model.StockPredictItem
        curPredictItem.StockInfo = util.ReadStockBasicInfo(request.StockCode)[0]
        curPredictItem.PredictionRecord = util.ReadPredictionData(singleStockPredFilePath, true)
        curPredictItem.ShowMsg = util.ReadPredictionMsg(path.Join(strings.ReplaceAll(singleStockPredFilePath, "csv", "txt")))
        rawStockData, _ := extractStockRawData(request.StockCode, request.QueryDateString, request.QueryDataLen)
        response.StockRawDatas = append(response.StockRawDatas, rawStockData)
        response.StockPredictDatas = append(response.StockPredictDatas, curPredictItem)
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
    fileList, _ := ioutil.ReadDir(indexFileDir)
    for _, fileItem := range(fileList) {
        if !strings.HasSuffix(fileItem.Name(), "csv") {
            continue
        }
        indexFileName := fileItem.Name()
        indexMsgFileName := strings.ReplaceAll(indexFileName, "csv", "txt")
        var singleIndexData model.IndexData
        singleIndexData.IndexName = strings.ReplaceAll(indexFileName, ".csv", "")
        singleIndexData.IndexRawData = util.ReadPredictionData(path.Join(indexFileDir, indexFileName), false)
        singleIndexData.IndexPredData = util.ReadPredictionData(path.Join(indexFileDir, indexFileName), true)
        singleIndexData.ShowMsg = util.ReadPredictionMsg(path.Join(indexFileDir, indexMsgFileName))
        response.IndexDataList = append(response.IndexDataList, singleIndexData)
    }
    for index, value := range(response.IndexDataList) {
        if value.IndexName == "all" {
            var curIndexData = response.IndexDataList[0]
            response.IndexDataList[0] = response.IndexDataList[index]
            response.IndexDataList[index] = curIndexData
        }
    }
    context.JSON(http.StatusOK, response)
}

// GetTotalMarketIndexData godoc
// @Summary Query Market Distribution
// @Description Return Market Distribution
// @ID getMarketDistribution
// @Accept json
// @Produce json
// @Param request_json body model.GetMarketDistributionRequest true "Query Index "
// @Success 200 {object} model.GetMarketDistributionResponse
// @Router /api/stock/getMarketDistribution [post]
func GetMarketDistribution(context *gin.Context) {
    request := model.GetMarketDistributionRequest{}
    if err := context.BindJSON(&request); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    var response model.GetMarketDistributionResponse
    filePath := path.Join(conf.StockPredictionBaseDir, strings.ReplaceAll(request.QueryDateString, "-", "/"), "summary", "main.csv")
    response.DistributionDataList = util.ReadDistributionData(filePath)
    context.JSON(http.StatusOK, response)
}
