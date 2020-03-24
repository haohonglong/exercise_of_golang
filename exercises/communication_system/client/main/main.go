package main

import (
	"fmt"
	"os"
	"exercises/communication_system/client/process"

)

func main(){
	
	var key int
	user := &process.UserProcess{}

	A:
	fmt.Println("\t      ==========welcome to login========")
	fmt.Println()
	fmt.Println("\t\t\t 1 登录聊天室")
	fmt.Println("\t\t\t 2 注册用户")
	fmt.Println("\t\t\t 3 退出系统")
	fmt.Println("\t\t\t 请选择（1-3):")

	fmt.Scanf("%d\n",&key)
	switch key {
	   case 1:
		fmt.Println("登录聊天室")
		fmt.Println("请输入用户的id")
		fmt.Scanf("%d\n",&user.UserId)
		fmt.Println("请输入用户的密码")
		fmt.Scanf("%s\n",&user.UserPwd)
		user.Login()
	   case 2:
		fmt.Println("注册用户")
	   case 3:
		fmt.Println("退出系统")
		os.Exit(0)
	   default:
		fmt.Println("你的输入有误，请重写输入")
		goto A

	}


	

}