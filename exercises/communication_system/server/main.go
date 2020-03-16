package main

import (
	"io"
	"log"
	"net"
	"fmt"
	"encoding/json"
	"encoding/binary"
	"exercises/communication_system/common/message"
	"console"
	_"errors"
)
func readPkg(c net.Conn) (mes message.Message, err error) {
	buf := make([]byte,8096)
	fmt.Printf("%s 服务器在等待客户端%s 发送信息\n",console.Log(), c.RemoteAddr().String())
	//conn.Read 在conn没有被关闭情况下，才会被堵塞

	_, err = c.Read(buf[:4])
	if err != nil {
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	//根据 packLen 读取消息内容
	n, err := c.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		// err = errors.New("conn.Read fail err")
		return
	}
	//把pkgLen 反序列化成 -> message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Printf("%s json.Unmarshal err=\n", console.Log(), err)
		return
	}

	return
}
func writePkg(conn net.Conn, data []byte) (err error) {
	//先发送一个长度
	var pkgLen uint32
	//强制转换uuint32类型
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//现在发送长度给对方，看看是否发送的成功，如果成功就发送真正的数据内容
	n, err := conn.Write(buf[:4])
	if n !=4 || err != nil {
		fmt.Printf("%s conn.Write(bytes) error:%s\n", console.Log(), err)
		return
	}
	//注意：这里是发送是真正的data数据
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Printf("%s conn.Write(bytes) error:%s\n", console.Log(), err)
		return
	}

	return

}

//用来处理登录的请求
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
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
	err = writePkg(conn,data)
	

	return

}

//编写一个ServerProcessMes 函数,这里就相当于一个router
//功能：根据客户端发送消息种类不同， 决定调用哪个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
		case message.LoginMesType:
			err = serverProcessLogin(conn, mes)
		case message.RegisterMesType:
		
		default:
			fmt.Printf("%s The type of message is not exist\n",console.Log())
		
	}
	return
}

func handleConnection(c net.Conn) {
	// Echo all incoming data.
	// io.Copy(c, c)
	// Shut down the connection.
	defer c.Close()
	

	for {
		//这里我们将读取数据包，直接封装成一个函数readPkg(),返回Message, Err
		mes, err := readPkg(c)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%s 客户端退出了\n", console.Log())
			}else {
				fmt.Printf("%s readPkg err\n", console.Log(), err)
			}
			return
			
		}
		err = serverProcessMes(c,&mes)
		if err != nil {
			return
		}
		
	}
}
func main() {
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", "0.0.0.0:2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}else{
			fmt.Printf("%s the ip of client is:%v\n", console.Log(), conn.RemoteAddr())
		}
		
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go handleConnection(conn)
	}
}
