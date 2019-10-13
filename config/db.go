package config

type DbConf struct {
	Host   string
	Port   string
	User   string
	Pwd    string
	DbName string
}

var MasterDbConfig = DbConf{
	//Host:   "192.168.10.10",
	Host: "127.0.0.1",
	Port: "3306",
	//User:   "homestead",
	User: "admin",
	//Pwd:    "secret",
	Pwd:    "secret",
	DbName: "go_blog",
}
