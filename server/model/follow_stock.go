package model

type FollowStock struct {
    Account string `gorm:"primaryKey"`
    StockCode string `gorm:"primaryKey"`
    Exchange string `gorm:"<-"`
    Name string `gorm:"<-"`
    Industry string `gorm:"<-"`
}
