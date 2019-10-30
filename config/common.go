package config

type YpConf struct {
	ApiKey string
	Tpl    string
}

//短信服务
var MasterYpConfig = YpConf{
	ApiKey: "9a0c********************",
	Tpl:    "32*****",
}
