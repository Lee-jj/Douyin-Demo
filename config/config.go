package config

import (
	"DOUYIN-DEMO/common"

	"github.com/spf13/viper"
)

type mysqlConfig struct {
	User     string
	Password string
	Ip       string
	Port     string
	Database string
}

type minioConfig struct {
	Endpoint        string
	AccessKsyID     string
	SecretAccessKey string
	VideoBucket     string
	ImageBucket     string
}

type Configs struct {
	MySQL mysqlConfig
	Minio minioConfig
}

var Config Configs

func GetConfig() Configs {
	return Config
}

func LoadConfig() {
	viper.SetConfigFile("./config.yaml")
	viper.ReadInConfig()
	err := viper.ReadInConfig()
	if err != nil {
		panic(common.ErrorGetConfigFaild)
	}

	mysql := mysqlConfig{
		User:     viper.GetString("mysql.user"),
		Password: viper.GetString("mysql.password"),
		Ip:       viper.GetString("mysql.ip"),
		Port:     viper.GetString("mysql.port"),
		Database: viper.GetString("mysql.database"),
	}

	minio := minioConfig{
		Endpoint:        viper.GetString("minio.endpoint"),
		AccessKsyID:     viper.GetString("minio.accessKeyID"),
		SecretAccessKey: viper.GetString("minio.secretAccessKey"),
		VideoBucket:     viper.GetString("minio.videoBucket"),
		ImageBucket:     viper.GetString("minio.imageBucket"),
	}

	Config = Configs{
		MySQL: mysql,
		Minio: minio,
	}
}
