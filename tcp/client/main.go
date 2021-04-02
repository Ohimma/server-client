package main

import (
	"bufio"
	"fmt"
	"net"
)

// 定义用户名和密码
var userID int
var userPwd string
var userName string

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "aaa GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("xxx GET / HTTP/1.0 %v \r\n\r\n", status)
	// data := "hello world"
	// _, err = conn.Write([]byte(data))
	// if err != nil {
	// 	fmt.Println("客户端 发送消息本身err = ", err)
	// 	return
	// }
	// fmt.Printf("客户端 发送消息成功 长度 = %d  内容 = %v \n", len(data), string(data))
}
