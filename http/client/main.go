package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ohimma/server-client/http/client/controller"
	// "github.com/prometheus/client_golang/prometheus/promhttp"
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

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

func main() {
	// 指定定时器的时间间隔是 1s

	go controller.Health1()

	// 阻塞方式一
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan
	close(stopChan)
	time.Sleep(5000 * time.Millisecond)

	// 阻塞进程方式二
	// var wg sync.WaitGroup
	// wg.Add(1)
	// wg.Wait()

	// flag.Parse()
	// http.Handle("/metrics", promhttp.Handler())
	// log.Fatal(http.ListenAndServe(*addr, nil))

	fmt.Println("stop server")

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
