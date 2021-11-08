package model

type FollowStock struct {
    Account string `gorm:"primaryKey" json:"account"`
    StockCode string `gorm:"primaryKey" json:"stock_code"`
    Exchange string `gorm:"<-" json:"exchange"`
    Name string `gorm:"<-" json:"name"`
    Industry string `gorm:"<-" json:"industry"`
}
