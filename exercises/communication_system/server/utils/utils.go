package utils
import (
	"net"
	"fmt"
	"encoding/json"
	"encoding/binary"
	"exercises/communication_system/common/message"
	"console"
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