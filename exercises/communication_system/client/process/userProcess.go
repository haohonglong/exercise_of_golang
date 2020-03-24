package process
import (
	"net"
	"fmt"
	"exercises/communication_system/common/message"
	"exercises/communication_system/client/utils"
	"encoding/json"
	"encoding/binary"
	"console"
	_"errors"
)
type UserProcess struct {
	UserId int
	UserPwd string
}

func (this *UserProcess) Login() (err error) {
	//1.连接到服务器
	conn, err := net.Dial("tcp", "0.0.0.0:2000")
	
	if err != nil {
		fmt.Printf("%s client dial error=\n", console.Log(), err)
		return
	}
	defer conn.Close()

	//2.准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType
	
	//3.创建一个LoginMes
	var loginMes message.LoginMes

	loginMes.UserId = this.UserId
	loginMes.UserPwd = this.UserPwd

	//4.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Printf("%s json.Marshal err=\n", console.Log(), err)
		return
	}

	//5.把data赋给mes.Data字段
	mes.Data = string(data)
	//6.将mes进行序列化


	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Printf("%s json.Marshal err=\n", console.Log(), err)
		return
	}

	//7.到了这里data就是我们要发送的消息
	//7.1 先把data的长度发送给服务器
	//先获取到data 的长度->转成一个表示长度的byte切片
	var pkgLen uint32
	//强制转换uuint32类型
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//现在发送长度
	n, err := conn.Write(buf[:4])
	if n !=4 || err != nil {
		fmt.Printf("%s conn.Write(bytes) error:%s\n", console.Log(), err)
		return
	}

	fmt.Printf("%s 客户端，发送消息的长度=%d 内容=%s\n", console.Log(), len(data), string(data))
	_, err = conn.Write(data)
	if err != nil {
		fmt.Printf("%s conn.Write(bytes) error:%s\n", console.Log(), err)
		return
	}
	// fmt.Println("休息一下哦。。。")
	// time.Sleep(10 * time.Second)


	// fmt.Printf(`客服端发送消息长度:%d\n
	// 内容是：%s\n
	
	// `, pkgLen, data)
	tf := &utils.Transfer{
		Conn : conn,
	}
	mes, err = tf.ReadPkg()
	
	
	if err != nil {
		fmt.Printf("%s readPkg(conn) err=", console.Log(), err)
		return
	}
	//将mes的Data部分反序列化成 LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	
	if loginResMes.Code == 200 {
		//这里还要在客户端启动一个协程
		//该协程和服务器端的通讯，
		go serverProcessMes(conn)
		menu := &Menu{}
		menu.Show()
	}else if loginResMes.Code == 500 {
		fmt.Printf("%s err:%s\n", console.Log(),loginResMes.Error)
		return
	}

	return


}
