package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ohimma/server-client/http/client/controller"
)

func main() {

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

	fmt.Println("stop server")

}
