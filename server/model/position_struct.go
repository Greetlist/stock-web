package model

type Position struct {
    Account string `gorm:"primaryKey"`
    InstId string `gorm:"primaryKey"`
    ExchangeId string `gorm:"<-"`
    FinanceType string `gorm:"<-"`
    Direction string `gorm:"<-"`
    PosValue float64 `gorm:"<-"`
    CurrentPosition int64 `gorm:"<-"`
    TotalBuyPosition int64 `gorm:"<-"`
    TotalSellPosition int64 `gorm:"<-"`
    TotalPnl float64 `gorm:"<-"`
    AvgCost float64 `gorm:"<-"`
}
