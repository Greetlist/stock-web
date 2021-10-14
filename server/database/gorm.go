package database

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "fmt"
    "os"
    "greetlist/stock-web/server/conf"
    "greetlist/stock-web/server/model"
)

var DBConnPool chan *gorm.DB = make(chan *gorm.DB, conf.DBConnectionNum)
func InitDataBase() {
    for i := 0; i < conf.DBConnectionNum; i++ {
        db, err := gorm.Open(mysql.Open(conf.DBConnectionStr), &gorm.Config{})
        if err != nil {
            fmt.Printf("Create DB Error: %s. Exit.\n", err);
            os.Exit(1)
        }
        DBConnPool <- db
    }
    migrateAllTable()
}

func migrateAllTable() {
    db, _ := <-DBConnPool
    db.AutoMigrate(
        &model.FundAccount{},
        &model.FollowStock{},
        &model.Position{},
    )
    DBConnPool <- db
}
