package config

import (
	"log"

	"github.com/go-ini/ini"
)

var (
	AppMode  string
	HttpPort string

	RedisAddress string
	RedisDB      string

	User     string
	Password string
	Host     string
	Port     string
	Db       string
)

func Init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		log.Fatal(err)
	}
	LoadServer(file)
	LoadMysql(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("port").String()
}

func LoadMysql(file *ini.File) {
	User = file.Section("mysql").Key("user").String()
	Password = file.Section("mysql").Key("password").String()
	Host = file.Section("mysql").Key("host").String()
	Port = file.Section("mysql").Key("port").String()
	Db = file.Section("mysql").Key("db").String()
}
