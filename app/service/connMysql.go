package service

import (
	"bolg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	//db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/example?charset=utf8")
	var err error
	DB, _ = gorm.Open("mysql", config.MasterDbConfig.User+":"+config.MasterDbConfig.Pwd+
		"@tcp("+config.MasterDbConfig.Host+":"+config.MasterDbConfig.Port+")/"+config.MasterDbConfig.DbName+"?parseTime=true&charset=utf8mb4")
	//检查数据库是否连接成功
	if DB.Error != nil {
		panic(err)
		return
	}
}
