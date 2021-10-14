package model
import (
    "time"
)

type Order struct {
    Account string `gorm:"primaryKey"`
    InstId string `gorm:"primaryKey"`
    ExchangeId string `gorm:"<-"`
    FinanceType string `gorm:"<-"`
    Direction string `gorm:"<-"`
    Price float64 `gorm:"<-"`
    Volume int64 `gorm:"<-"`
    OrderDate time.Time `gorm:"<-"`
}
