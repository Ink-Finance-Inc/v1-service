package conf

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// const MySQLDSN = "root:9988123@(118.195.152.119:3306)/xjh_testdb3?charset=utf8mb4&parseTime=True&loc=Local"
var MySQLDSN = "root:9988123@(118.195.152.119:3306)/xjh_test_local?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func ConnectDB(dsn string) *gorm.DB {

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %s", err))
	}

	return db
}
