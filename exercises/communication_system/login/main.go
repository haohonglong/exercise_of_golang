package main
import (
	"fmt"
	"exercises/communication_system/login/login"
)

func main() {
	login := &login.LoginS{}
	fmt.Println("helll")
	login.Init()
	key := login.Key

	switch key {
		case 1:
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n",&login.UserId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n",&login.UserPwd)
			err := login.Login()
			if err != nil {
				fmt.Println("fail")
			}else {
				fmt.Println("success")
			}
		case 2:
		
		default:


	}


}