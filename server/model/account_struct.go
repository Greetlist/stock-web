package model

type FundAccount struct {
    Account string `gorm:"primaryKey"`
    Password string `gorm:"<-"`
    EnableBalance float64 `gorm:"<-"`
    FrozenBalance float64 `gorm:"<-"`
    WithdrawBalance float64 `gorm:"<-"`
    TotalAssets float64 `gorm:"<-"`
}
