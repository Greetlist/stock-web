package model

type GetQueryStockDataRequest struct {
    StockList []string `json:"stock_list" example:['000001', '002142', ...] binding:"required"`
}

type SingleRecord struct {
    Date string `json:"date"`
    Open float64 `json:"open"`
    Close float64 `json:"close"`
    High float64 `json:"high"`
    Low float64 `json:"low"`
    Volume float64 `json:"volume"`
    Money float64 `json:"money"`
    Ma13 float64 `json:"ma13"`
    Ma34 float64 `json:"ma34"`
    Ma55 float64 `json:"ma55"`
}

type StockDataItem struct {
    StockCode string `json:"stock_code" example:"002142"`
    Records []SingleRecord `json:"records"`
}

type GetQueryStockDataResponse struct {
    StockDatas []StockDataItem `json:"stock_datas"`
}

type GetAllStockCodeResponse struct {
    StockList []string `json:"stock_list" example:"['000001', '002142', ...]"`
}
