package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDB はデータベース接続を行い、*gorm.DBを返します
func ConnectDB() (*gorm.DB, error) {
	dsn := "root@tcp(mysql:3306)/RecoDb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
    return db, nil
}
