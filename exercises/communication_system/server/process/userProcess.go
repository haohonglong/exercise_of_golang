package process
import (
	"net"
	"fmt"
	"encoding/json"
	"exercises/communication_system/common/message"
	"exercises/communication_system/server/utils"
	"console"
)
type UserProcess struct {
	Conn net.Conn
}

//用来处理登录的请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//1.先从mes 中取出 mes.Data,并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	//1.
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//2.
	var loginResMes message.LoginResMes
	


	//如果用户id=100，密码=123456，任务合法，否则不合法
	if loginMes.UserId == 100 &&  loginMes.UserPwd == "123456" {//合法
		loginResMes.Code = 200

	}else {
		loginResMes.Code = 500 //500 状态码，表示该用户不存在
		loginResMes.Error = "该用户不存在，请先注册再使用..."
	}

	//3.将loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Printf("%s jsonMarshal fail\n", console.Log(), err)
		return
	}

	//4.将data 赋值给resMes
	resMes.Data = string(data)

	//5. 对resMes 序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Printf("%s jsonMarshal fail\n", console.Log(), err)
		return
	}
	//6.发送data, 将其封装到writePkg函数中
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	

	return

}