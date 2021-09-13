package model

// Request Struct start
type QueryStockDataRequest struct {
    StockList []string `json:"stock_list" example:['000001', '002142', ...] binding:"required"`
    QueryDataLen int `json:"query_data_len" example: 30 binding:"required" `
}
type GetRecommandStockRequest struct {
    QueryDateString string `json:"query_date" example:"2021-07-01"`
    QueryDataLen int `json:"query_data_len" example: 30 binding:"required" `
}
// Request Struct end

// Response Struct start
type QueryStockDataResponse struct {
    StockTotalDatas []SingleStockData `json:"stock_datas"`
}

type GetAllStockCodeResponse struct {
    StockInfoList []StockBasicInfo `json:"stock_info_list"`
}

type GetRecommandStockResponse struct {
    StockRawData []SingleStockData `json:"stock_data"`
    StockPredictDatas []StockPredictItem `json:"stock_prediction_datas"`
}
// Response Struct end
