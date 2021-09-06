package model

// Request Struct start
type GetQueryStockDataRequest struct {
    StockList []string `json:"stock_list" example:['000001', '002142', ...] binding:"required"`
    QueryDataLen int `json:"query_data_len" example: 30 binding:"required" `
}
type GetRecommandStockRequest struct {
    QueryDateString string `json:"query_date" example:"2021-07-01"`
    QueryDataLen int `json:"query_data_len" example: 30 binding:"required" `
}
// Request Struct end

// Response Struct start
type GetQueryStockDataResponse struct {
    StockDatas []StockDataItem `json:"stock_datas"`
}

type GetAllStockCodeResponse struct {
    StockInfoList []StockBasicInfo `json:"stock_info_list"`
}

type GetDailyCalcStockDataResponse struct {
    StockDatas []StockDataItem `json:"stock_datas"`
}

type GetRecommandStockResponse struct {
    StockRawDatas []StockDataItem `json:"stock_datas"`
    StockPredictDatas []StockPredictItem `json:"stock_prediction_datas"`
}
// Response Struct end

type SingleRecord struct {
    Date string `json:"date" csv:"trade_date"`
    Open float64 `json:"open" csv:"open"`
    Close float64 `json:"close" csv:"close"`
    High float64 `json:"high" csv:"high"`
    Low float64 `json:"low" csv:"low"`
    Volume float64 `json:"volume" csv:"vol"`
    Money float64 `json:"money" csv:"amount"`
    Ma13 float64 `json:"ma13" csv:"MA13"`
    Ma34 float64 `json:"ma34" csv:"MA34"`
    Ma55 float64 `json:"ma55" csv:"MA55"`
}

type SinglePredictRecord struct {
    Date string `json:"date" csv:"day_to_prediction"`
    Open float64 `json:"open" csv:"open"`
    Close float64 `json:"close" csv:"close"`
    High float64 `json:"high" csv:"high"`
    Low float64 `json:"low" csv:"low"`
}

type StockDataItem struct {
    StockCode string `json:"stock_code" example:"002142.SZ"`
    Records []SingleRecord `json:"records"`
}

type StockPredictItem struct {
    StockCode string `json:"stock_code" example:"002142.SZ"`
    PredictionRecords []SinglePredictRecord `json:"prediction_records"`
}

type StockBasicInfo struct {
    StockCode string `json:"stock_code" example:"002142.SZ" csv:"ts_code"`
    StockName string `json:"stock_name" example:"宁波银行" csv:"name"`
}
