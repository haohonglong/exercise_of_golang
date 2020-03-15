package message

//消息类型
const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`//the content of message
}

type LoginMes struct {
	UserId int `json:"userId"`
	UserName string `json:"userName"`
	UserPwd string `json:"userPwd"`//the password of user
}

type LoginResMes struct {
	Code int `json:"code"`//状态码 
	Error string `json:"error"`// the message of error
}

type RegisterMes struct {

}

