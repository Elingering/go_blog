package service

import (
	"bolg/config"
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	//db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/example?charset=utf8")
	var err error
	DB, _ := gorm.Open("mysql", config.MasterDbConfig.User+":"+config.MasterDbConfig.Pwd+
		"@tcp("+config.MasterDbConfig.Host+":"+config.MasterDbConfig.Port+")/"+config.MasterDbConfig.DbName+"?parseTime=true&charset=utf8mb4")
	//检查数据库类型是否符合
	if DB.Error != nil {
		fmt.Println("错误", err)
		return
	} else {
		fmt.Println("连接成功")
	}
	//检查是否连接成功
	//c := DB.()
	//if c != nil {
	//	fmt.Println("错误", err)
	//	return
	//} else {
	//	fmt.Println("连接成功")
	//}
}
