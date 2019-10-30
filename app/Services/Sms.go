package Services

import (
	"bolg/app/Helper"
	"bolg/config"
	"fmt"
	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
	"net/url"
	"time"
)

const PREFIX = "code_"

func GetCode(phone string) {
	//生成6位随机数
	code := Helper.Random6()
	// 发送短信
	server := ypclnt.New(config.MasterYpConfig.ApiKey)
	param := ypclnt.NewParam(3)
	param[ypclnt.MOBILE] = phone
	param[ypclnt.TPL_ID] = config.MasterYpConfig.Tpl
	param[ypclnt.TPL_VALUE] = url.Values{"#name#": {"用户"}, "#code#": {code}, "#hour#": {"5"}}.Encode()
	r := server.Sms().TplSingleSend(param)
	fmt.Println(r.Msg)
	//账户:clnt.User() 签名:clnt.Sign() 模版:clnt.Tpl()//3265430 短信:clnt.Sms() 语音:clnt.Voice() 流量:clnt.Flow()
	//存入Redis 5分钟失效
	TX.SetNX(PREFIX+phone, code, 5*60*time.Second)
}
