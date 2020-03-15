package main

import (
	"fmt"
	myLogin "exercises/communication_system/login/login"

)

func main(){
	
	loginS := &myLogin.LoginS{}
	loginS.Init()
	key := loginS.Key

	switch key {
		case 1:
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n",&loginS.UserId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n",&loginS.UserPwd)
			login(loginS.UserId, loginS.UserPwd)
			// if err != nil {
			// 	fmt.Println("登录失败")
			// }else {
			// 	fmt.Println("success")
			// }
		case 2:
		
		default:


	}
	

}