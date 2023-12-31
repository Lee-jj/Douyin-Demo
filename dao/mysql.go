package dao

import (
	"DOUYIN-DEMO/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL() {
	var DBError error

	// using viper import yaml
	conf := config.GetConfig()
	user := conf.MySQL.User
	password := conf.MySQL.Password
	ip := conf.MySQL.Ip
	port := conf.MySQL.Port
	database := conf.MySQL.Database

	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local" // sample
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, database)

	DB, DBError = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if DBError != nil {
		log.Fatal(DBError)
	}

	// fmt.Println(database)
	fmt.Printf("link database: %v successfully.\n", database)
}
