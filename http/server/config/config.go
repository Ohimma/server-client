package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type All struct {
	Server `yaml:"server"`
	Mysql  `json:"mysql" yaml:"mysql"`
}

type Server struct {
	Port     int16  `json:"port" yaml:"port"`
	Mode     string `json:"mode" yaml:"mode"`
	UseRedis bool   `json:"useRedis" yaml:"useRedis"`
}

type Mysql struct {
	Host        string `json:"host" yaml:"host"`
	Db          string `json:"db" yaml:"db"`
	User        string `json:"user" yaml:"user"`
	Password    string `json:"password" yaml:"password"`
	Config      string `json:"config" yaml:"config"`
	AutoMigrate bool   `json:"autoMigrate" yaml:"autoMigrate"`
}

var Conf *All

func init() {

	viper.BindEnv("ENV")
	env := viper.Get("ENV")
	log.Print("ENV = ", env)
	if env == "prod" {
		viper.SetConfigName("config_prod")
	} else {
		viper.SetConfigName("config_dev")
	}
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	viper.AddConfigPath(dir)
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		log.Print("读取config配置错误", err)
	}

	var c All
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Print("config unmarshal struct faild", err)
	}
	Conf = &c
}
