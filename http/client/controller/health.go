package controller

import (
	"log"
	"time"

	"github.com/Ohimma/server-client/http/client/config"
	"github.com/Ohimma/server-client/http/client/utils"
)

type HealthRequest struct {
	Host string `json:"host" form:"host"`
}

func Health1() {
	t := time.NewTicker(1000 * time.Millisecond)
	for t1 := range t.C {
		time := t1.Format("2006-01-02 15:04:05")
		log.Println("Tick at", time) //每取一次，这里隔一秒会执行一次
		SendHealth2()
	}
	t.Stop()
	log.Println("停止该定时器")
}

func SendHealth() {
	url := "http://" + config.Conf.Server.Host + "/test/health"

	log.Println(url)

	resp := utils.Get(url)
	log.Println(resp)

}
func SendHealth2() {
	url := "http://" + config.Conf.Server.Host + "/test/health"

	log.Println(url)

	data := HealthRequest{
		Host: "127.0.0.2",
	}
	contentType := "application/json;charset=UTF-8"

	resp := utils.Post(url, data, contentType)
	log.Println(resp)
}
