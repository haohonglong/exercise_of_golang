package login
import (
	"fmt"
	"os"
)
type LoginS struct {
	UserId, Key int
	UserName, UserPwd string
}

func(l *LoginS) Login() (err error) {
	fmt.Printf("userId = %d userPwd = %s\n",l.UserId,l.UserPwd)
	return nil
	
}
func(l *LoginS) Init() {
	var loop bool = true
	var key int
	for loop{
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
			loop = false
		   case 2:
			fmt.Println("注册用户")
			loop = false
		   case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		   default:
			fmt.Println("你的输入有误，请重写输入")

		}
		if !loop {
			break
		}
	}
	l.Key = key
}
