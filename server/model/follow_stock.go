package model

type FollowStock struct {
    Account string `gorm:"primaryKey"`
    InstId string `gorm:"primaryKey"`
    ExchangeId string `gorm:"<-"`
}
