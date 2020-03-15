package main

import (
	"net"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	conn, err := net.Dial("tcp", "0.0.0.0:2000")
	
	if err != nil {
		fmt.Println("client dial error=",err)
		return
	}
	// fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	
	for {

		line, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("bufio.NewReader err:",err)
			return
		}

		n, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
		fmt.Println("客户端发送了 %d 字节的数据",n)
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("程序退出！")
			break
		}
		

	}
	
	

}