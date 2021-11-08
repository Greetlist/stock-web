package model

type AddPreferStockRequest struct {
    Account string `json:"account"`
    StockCode string `json:"stock_code"`
}

type AddPreferStockResponse struct {
    IsSuccess bool `json:"is_success"`
    ErrMsg string `json:"error_msg"`
}

type DeletePreferStockRequest struct {
    Account string `json:"account"`
    StockCode string `json:"stock_code"`
}

type DeletePreferStockResponse struct {
    IsSuccess bool `json:"is_success"`
    ErrMsg string `json:"error_msg"`
}

type GetPreferStockRequest struct {
    Account string `json:"account"`
}

type GetPreferStockResponse struct {
    IsSuccess bool `json:"is_success"`
    ErrMsg string `json:"error_msg"`
    FollowStockList []FollowStock `json:"follow_stock_list"`
}
