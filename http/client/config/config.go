package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type All struct {
	Server `yaml:"server"`
}

type Server struct {
	Host string `json:"host" yaml:"host"`
	Mode string `json:"mode" yaml:"mode"`
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
