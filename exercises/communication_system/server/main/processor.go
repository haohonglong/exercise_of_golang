package main
import (
	"io"
	"net"
	"fmt"
	"exercises/communication_system/common/message"
	"exercises/communication_system/server/process"
	"exercises/communication_system/server/utils"
	"console"
)
type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcessMes 函数,这里就相当于一个router
//功能：根据客户端发送消息种类不同， 决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
		case message.LoginMesType:
			up := &process.UserProcess{
				Conn : this.Conn,
			}
			err = up.ServerProcessLogin(mes)
		case message.RegisterMesType:
		
		default:
			fmt.Printf("%s The type of message is not exist\n",console.Log())
		
	}
	return
}

func (this *Processor) process2() (err error) {
	for {
		//这里我们将读取数据包，直接封装成一个函数readPkg(),返回Message, Err
		tf := &utils.Transfer{
			Conn : this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%s 客户端退出了\n", console.Log())
				return err
			}else {
				fmt.Printf("%s readPkg err\n", console.Log(), err)
				return err
			}
			return err
			
		}
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
		
	}
}