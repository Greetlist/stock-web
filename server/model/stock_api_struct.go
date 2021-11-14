package model

// Request Struct start
type QueryStockDataRequest struct {
    StockList []string `json:"stock_list" example:['000001', '002142', ...] binding:"required"`
    QueryDataLen int `json:"query_data_len" example: 30 binding:"required" `
}
type GetRecommandStockRequest struct {
    QueryDateString string `json:"query_date" example:"2021-07-01"`
    StockCode string `json:"stock_code", example:"002142.SZ"`
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
    StockRawDatas []SingleStockData `json:"stock_datas"`
    StockPredictDatas []StockPredictItem `json:"stock_prediction_datas"`
}

type GetTotalMarketIndexDataRequest struct {
    QueryDateString string `json:"query_date" example:"2021-07-01"`
}

type GetTotalMarketIndexDataResponse struct {
    IndexRawData SinglePredictRecord `json:"index_raw_data"`
    IndexPredData SinglePredictRecord `json:"index_pred_data"`
    ShowMsg string `json:"show_msg"`
}


// Response Struct end
