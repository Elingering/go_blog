package config

type DbConf struct {
	Host   string
	Port   string
	User   string
	Pwd    string
	DbName string
}

var MasterDbConfig = DbConf{
	Host:   "192.168.10.10",
	Port:   "3306",
	User:   "homestead",
	Pwd:    "secret",
	DbName: "go_blog",
}