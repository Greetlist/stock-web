package model
import (
    "time"
)

type Profit struct {
    Account string `gorm:"primaryKey"`
    InstId string `gorm:"primaryKey"`
    ExchangeId string `gorm:"<-"`
    Pnl float64 `gorm:"<-"`
    Date time.Time `gorm:"<-"`
}
