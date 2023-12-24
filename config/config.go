package config

import (
	"log"
	"gopkg.in/ini.v1"
	"github.com/aihou/todo/utils"
)

type ConfigList struct {
	Port string
	SQLDriver string
	DbName string
	LogFile string
	Static string
}

var Config ConfigList

//init関数はmain関数より先に実行される
func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port: cfg.Section("web").Key("port").String(),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName: cfg.Section("db").Key("name").String(),
		LogFile: cfg.Section("web").Key("logfile").String(),
		Static: cfg.Section("web").Key("static").String(),
	}
}