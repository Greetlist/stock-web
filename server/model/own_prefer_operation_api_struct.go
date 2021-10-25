package model

type AddPreferStockRequest struct {
    Account string `json:"account"`
    StockCode string `json:"stock_code"`
}

type AddPreferStockResponse struct {
    IsSuccess bool `json:"is_success"`
}
