package model

type GetQueryStockDataRequest struct {
    StockList []string `json:"stock_list" example:"['000001', '002142', ...]" binding:"required"`
}

type SingleRecord struct {
    Date string `json:"date"`
    Open float64 `json:"open"`
    Close float64 `json:"close"`
    High float64 `json:"high"`
    Low float64 `json:"low"`
    Volume float64 `json:"volume"`
    Money float64 `json:"money"`
}

type StockDataItem struct {
    StockCode string `json:"stock_code" example:"002142"`
    Records []SingleRecord `json:"records"`
}

type GetQueryStockDataResponse struct {
    StockDatas []StockDataItem `json:"stock_datas"`
}
