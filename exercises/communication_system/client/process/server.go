package process

import (
	"fmt"
	"os"
	"exercises/communication_system/client/utils"
	"console"
	"net"
)
type Menu struct {

}
//显示登录成功后的界面
func (this *Menu) Show() {
	var key int
	fmt.Println("-------恭喜用户登录成功-------")
	A:
	fmt.Println("-------1.显示在线用户列表")
	fmt.Println("-------2.发送信息")
	fmt.Println("-------3.信息列表")
	fmt.Println("-------4.退出系统")
	fmt.Println("请选择(1-4)")
	fmt.Scanf("%d\n",&key)
	switch key {
		case 1 :
			fmt.Println("显示在线用户列表")
		case 2 :
			fmt.Println("发送信息")
		case 3 :
			fmt.Println("信息列表")
		case 4 :
			fmt.Println("退出系统")
			os.Exit(0)
		default :
			fmt.Println("请选择一个正确的选项！")
			goto A

	}
}

func serverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn : conn,
	}
	for {
		fmt.Printf("%s waiting the client read  message  that  send from server\n", console.Log())
		mes,  err := tf.ReadPkg()
		if err != nil {
			fmt.Printf("%s tf.ReadPkg() err= \n", console.Log(), err)
			return 
		}

		//如果读取到了
		fmt.Printf("%s mes=%v\n", console.Log(), mes)

	}
}