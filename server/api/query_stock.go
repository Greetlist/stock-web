package api

import (
    "net/http"
    "io/ioutil"
    "strings"
    "fmt"
    "os"
    "path"
    "github.com/gin-gonic/gin"
    "github.com/gocarina/gocsv"
    "greetlist/stock-web/server/model"
    "greetlist/stock-web/server/conf"
    "greetlist/stock-web/server/util"
)

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
    if err := context.BindJSON(&request); err != nil {
	    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    var response model.GetQueryStockDataResponse
    for _, stockCode := range(request.StockList) {
        var stockDataItem model.StockDataItem
        stockDataItem.StockCode = stockCode
        stockFilePath := path.Join(conf.StockMaBaseDir, "day", selectExchange(stockCode), stockCode + ".ma.csv")
        stockDataItem.Records = extractSingleStockDataFromCsv(stockFilePath, request.QueryDataLen)
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
    var stockInfoList []model.StockBasicInfo
    fileList := []string{"total_stock.csv", "total_index.csv"}
    for _, fileName := range(fileList) {
        var curStockList []*model.StockBasicInfo
        f, _ := os.OpenFile(path.Join(conf.StockRawBaseDir, fileName), os.O_RDWR, os.ModePerm)
        if err := gocsv.UnmarshalFile(f, &curStockList); err != nil {
            fmt.Printf("Unmarshal csv file error : %s.\n", err)
        }
        for _, stockItemPtr := range(curStockList) {
            stockInfoList = append(stockInfoList, *stockItemPtr)
        }
    }
    response := model.GetAllStockCodeResponse{StockInfoList: stockInfoList}
    context.JSON(http.StatusOK, response)
}

// GetAllStockCode godoc
// @Summary Query Daily Calc Stock Data
// @Description Return All Daily Calc Stock Data
// @ID getDailyCalcStockData
// @Produce json
// @Success 200 {object} model.GetDailyCalcStockDataResponse
// @Router /api/stock/getDailyCalcStockData [get]
func GetDailyCalcStockData(context *gin.Context) {
    todayStr := strings.ReplaceAll(util.GetTodayStr(), "-", "/")
    dailyFilePath := path.Join(conf.DailyResultBaseDir, todayStr, "/ma_statistic_result.csv")

    type result struct {
        StockCode string `csv:"StockCode"`
        CheckFlag bool `csv:"CheckFlag"`
    }
    f, err := os.OpenFile(dailyFilePath, os.O_RDWR, os.ModePerm)
    if err != nil {
        fmt.Printf("Open File %s Error : %s.\n", dailyFilePath, err)
    }
    defer f.Close()
    var resultList []*result
    if err := gocsv.UnmarshalFile(f, &resultList); err != nil {
        fmt.Printf("Unmarshal csv file error : %s.\n", err)
	}
    var dailyStockCodeList []string
    for _, value := range(resultList) {
        if (*value).CheckFlag {
            dailyStockCodeList = append(dailyStockCodeList, (*value).StockCode)
        }
    }

    var stockDatas []model.StockDataItem
    for _, stockCode := range(dailyStockCodeList) {
        var curItem model.StockDataItem
        curItem.StockCode = stockCode
        stockCsvFilePath := path.Join(conf.StockResultBaseDir, stockCode, "/result_ma.csv")
        curItem.Records = extractSingleStockDataFromCsv(stockCsvFilePath, 0)
        stockDatas = append(stockDatas, curItem)
    }

    response := model.GetDailyCalcStockDataResponse{stockDatas}
    context.JSON(http.StatusOK, response)
}

func extractSingleStockDataFromCsv(stockFilePath string, queryDataLen int) []model.SingleRecord {
    dataLenGap := 30
    if queryDataLen != 0 {
        dataLenGap = queryDataLen
    }
    var res []model.SingleRecord
    f, err := os.OpenFile(stockFilePath, os.O_RDWR, os.ModePerm)
    defer f.Close()
    if err != nil {
        fmt.Printf("Open File %s Error : %s.\n", stockFilePath, err)
        return res
    }

    var curStockRecords []*model.SingleRecord
    if err := gocsv.UnmarshalFile(f, &curStockRecords); err != nil {
        fmt.Printf("Unmarshal csv file error : %s.\n", err)
    }
    dataLen := len(curStockRecords)
    startIndex := 0
    if dataLen > dataLenGap {
        startIndex = dataLen - dataLenGap
    }
    curStockRecords = curStockRecords[startIndex:]
    for _, value := range(curStockRecords) {
        res = append(res, *value)
    }
    return res
}

func extractPredictionDataFromCsv(stockFilePath string) []model.SinglePredictRecord {
    var res []model.SinglePredictRecord
    f, err := os.OpenFile(stockFilePath, os.O_RDWR, os.ModePerm)
    defer f.Close()
    if err != nil {
        fmt.Printf("Open File %s Error : %s.\n", stockFilePath, err)
        return res
    }
    var curPredictionRecords []*model.SinglePredictRecord
    if err := gocsv.UnmarshalFile(f, &curPredictionRecords); err != nil {
        fmt.Printf("Unmarshal csv file error : %s.\n", err)
    }
    for _, value := range(curPredictionRecords) {
        res = append(res, *value)
    }
    return res
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
        if !strings.HasSuffix(fileItem.Name(), "compared.csv") {
            var curItem model.StockDataItem
            curItem.StockCode = strings.ReplaceAll(fileItem.Name(), ".csv", "")
            maCsvFileName := strings.ReplaceAll(fileItem.Name(), ".csv", ".ma.csv")
            maCsvFilePath := path.Join(conf.StockMaBaseDir, "day", selectExchange(fileItem.Name()), maCsvFileName)
            curItem.Records = extractSingleStockDataFromCsv(maCsvFilePath, request.QueryDataLen)
            var curPredictItem model.StockPredictItem
            curPredictItem.StockCode = curItem.StockCode
            curPredictItem.PredictionRecords = extractPredictionDataFromCsv(path.Join(predictionDir, fileItem.Name()))
            response.StockRawDatas = append(response.StockRawDatas, curItem)
            response.StockPredictDatas = append(response.StockPredictDatas, curPredictItem)
        }
    }
    context.JSON(http.StatusOK, response)
}

func selectExchange(stockCode string) string {
    if strings.Contains(stockCode, "SH") {
        return "SSE"
    } else if strings.Contains(stockCode, "SZ") {
        return "SZSE"
    }
    return ""
}
