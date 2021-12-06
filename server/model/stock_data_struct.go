package model

type StockSinglePeriodData struct {
    PeriodType string `json:"period_type"`
    MarketData MarketRawData `json:"market_raw_data"`
    FactorData FactorRawData `json:"factor_raw_data"`
}

type MarketRawData struct {
    Date []string `json:"date" csv:"trade_date"`
    Open []float64 `json:"open" csv:"open"`
    Close []float64 `json:"close" csv:"close"`
    High []float64 `json:"high" csv:"high"`
    Low []float64 `json:"low" csv:"low"`
    Volume []float64 `json:"volume" csv:"vol"`
    Money []float64 `json:"money" csv:"amount"`
}

type FactorRawData struct {
    MA13 []float64 `json:"ma13" csv:"MA13"`
    MA34 []float64 `json:"ma34" csv:"MA34"`
    MA55 []float64 `json:"ma55" csv:"MA55"`
    EMA13 []float64 `json:"ema13" csv:"EMA13"`
    EMA34 []float64 `json:"ema34" csv:"EMA34"`
    EMA55 []float64 `json:"ema55" csv:"EMA55"`
}

type SinglePredictRecord struct {
    Date []string `json:"date" csv:"day_to_prediction"`
    Open []float64 `json:"open" csv:"open"`
    Close []float64 `json:"close" csv:"close"`
    High []float64 `json:"high" csv:"high"`
    Low []float64 `json:"low" csv:"low"`
}

type StockPredictItem struct {
    StockInfo StockBasicInfo `json:"stock_info"`
    PredictionRecord SinglePredictRecord `json:"prediction_record"`
    ShowMsg string `json:"show_msg"`
}

type StockBasicInfo struct {
    StockCode string `json:"stock_code" example:"002142.SZ" csv:"ts_code"`
    StockName string `json:"stock_name" example:"宁波银行" csv:"name"`
    StockIndustry string `json:"stock_industry" example:"石油" csv:"industry"`
    Exchange string `json:"exchange" example:"BJSE/SZSE/SSE"`
}

type SingleStockData struct {
    StockInfo StockBasicInfo `json:"stock_info"`
    StockRawDatas []StockSinglePeriodData `json:"stock_raw_datas"`
}

type IndexData struct {
    IndexName string `json:"index_name"`
    IndexRawData SinglePredictRecord `json:"index_raw_data"`
    IndexPredData SinglePredictRecord `json:"index_pred_data"`
    ShowMsg string `json:"show_msg"`
}
