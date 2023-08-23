package dao

import (
	"DOUYIN-DEMO/common"
	"fmt"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL() {
	var DBError error

	cfg, err := ini.Load("config.ini")
	if err != nil {
		panic(common.ErrorGetIniFaild)
	}

	user := cfg.Section("mysql").Key("user").String()
	password := cfg.Section("mysql").Key("password").String()
	ip := cfg.Section("mysql").Key("ip").String()
	port := cfg.Section("mysql").Key("port").String()
	database := cfg.Section("mysql").Key("database").String()

	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local" // sample
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, database)

	DB, DBError = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if DBError != nil {
		panic(DBError)
	}

	// fmt.Println(database)
	fmt.Printf("link database: %v successfully.\n", database)
}
