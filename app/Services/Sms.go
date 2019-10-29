package Services

import (
	"bolg/app/Helper"
	"time"
)

const PREFIX = "code_"

func GetCode(phone string) string {
	//生成6位随机数
	code := Helper.Random6()
	// TODO 发送短信
	//存入Redis 5分钟失效
	TX.SetNX(PREFIX+phone, code, 5*60*time.Second)
	return code
}
