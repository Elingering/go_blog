package Services

import (
	"bolg/config"
	"github.com/go-redis/redis/v7"
)

var TX *redis.Client

func init() {
	//db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/example?charset=utf8")
	var err error
	TX = redis.NewClient(&redis.Options{
		Addr:     config.MasterTxConfig.Addr,
		Password: config.MasterTxConfig.Password, // no password set
		DB:       config.MasterTxConfig.DB,       // use default DB
	})
	_, err = TX.Ping().Result()
	//检查redis是否连接成功
	if err != nil {
		panic(err)
		return
	}
}
