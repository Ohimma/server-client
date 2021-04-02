package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Ohimma/server-client/http/client/config"
)

// func main() {

// 	// 引入配置
// 	fmt.Print("conf = ", config.Conf.Server.Host)
// 	// 原始 url
// 	// http.HandleFunc("/", Hello)
// 	// 引用router ，gin使用的包
// 	router := httprouter.New()
// 	router.GET("/", Index)
// 	router.GET("/hello/:name", Hello)

// 	fmt.Println("服务器开始启动")
// 	log.Fatal(http.ListenAndServe(":8080", router))
// 	// err := http.ListenAndServe("0.0.0.0:8080", nil)
// 	// if err != nil {
// 	// 	fmt.Println("启动服务器错误")
// 	// 	return
// 	// }

// }

func Get(url string) {
	// resp, err := http.Get(url)
	// 要管理客户端的头部等，则建一个client实例，否则可直接get
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp)
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

func runServer() {
	fmt.Println("prometheus runServer")
}

func main() {
	i := 1
	for {
		url := "http://" + config.Conf.Server.Host + "/test/health"
		log.Info(url)
		fmt.Println(i, "xxxx", url)
		Get(url)

		time.Sleep(1 * time.Second)
		i++
	}
}

// func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	fmt.Println("r.Method = ", r.Method)
// 	fmt.Println("r.URL = ", r.URL)
// 	fmt.Println("r.Header = ", r.Header)
// 	fmt.Println("r.Body = ", r.Body)
// 	fmt.Fprint(w, "Welcome!\n")
// }

// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
// }
