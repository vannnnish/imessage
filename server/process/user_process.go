package process

import (
	"encoding/json"
	"fmt"
	"imessage/common/message"
	"imessage/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (process *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("反序列化失败:", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		// 合法
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在"
		// 不合法
	}
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json marshal 失败", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//
	tf := &utils.Transfer{
		Conn: process.Conn,
		Buf:  [8096]byte{},
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	return
}
