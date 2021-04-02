package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	fmt.Println("xxxxx", conn)
	// 创建总控
	// processor := &Processor{
	// 	Conn: conn,
	// }
	// err := processor.process2()
	// if err != nil {
	// 	fmt.Println("客户端和服务端通讯协程错误 err = ", err)
	// 	return
	// }
}

// 对 userDao 实例初始化
// func initUserDao() {
// 	model.MyUserDao = model.NewUserDao(pool) // 这里的pool已经在redis.go中全局定义了
// }

func main() {
	// initPool("111.206.251.155:6379", 10, 0, 300*time.Second)
	// initUserDao()

	// 1. 启动监听 tcp 8888 端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	defer listen.Close()

	if err != nil {
		fmt.Println("服务端 listen err = ", err)
		return
	} else {
		fmt.Println("服务器 开始监听 8888 端口.....")
	}

	// 2. 服务端 循环等待 客户端连接
	for {
		fmt.Println("服务端 循环等待客户端的连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("服务端 listen.Accept err = ", err)
			return
		}

		// 3. 一旦连接成功，服务端开启协程 循环保持和客户端通讯
		go process(conn)

	}

}
