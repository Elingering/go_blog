package Services

import (
	"bolg/app/Helper"
	"fmt"
	"github.com/mojocn/base64Captcha"
)

const PREFIX_C = "chapter_"

func GetVerificationCode() (string, string) {
	//数字验证码配置
	var configD = base64Captcha.ConfigDigit{
		Height:     80,
		Width:      240,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 5,
	}
	//创建数字验证码
	idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	//以base64编码
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)
	//存入Redis
	num := Helper.Random6()
	key := PREFIX_C + num
	TX.Set(key, idKeyD, 0)
	fmt.Println(idKeyD)
	return key, base64stringD
}

func VerificationCode(idkey, verifyValue string) bool {
	verifyResult := base64Captcha.VerifyCaptcha(idkey, verifyValue)
	println(idkey)
	println(verifyValue)
	println(verifyResult)
	if verifyResult {
		return true
	} else {
		return false
	}
}
