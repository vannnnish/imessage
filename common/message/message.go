package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType     = "RegisterMes"
)

type Message struct {
	Type string // 消息类型
	Data string // 消息内容
}

type LoginMes struct {
	UserId   int    // 用户id
	UserPwd  string // 用户密码
	UserName string // 用户昵称
}

type LoginResMes struct {
	Code  int    // 状态码 500 表示用户未注册  200 表示登录成功
	Error string // 返回错误信息
}

type RegisterMes struct {
	//
}
