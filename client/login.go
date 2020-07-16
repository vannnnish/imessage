package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"imessage/common/message"
	"net"
)

func login(userId int, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("连接失败:", err.Error())
		return
	}
	// 延时关闭
	defer conn.Close()
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	// 将loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json marshal 失败:", err)
		return
	}
	fmt.Println("loginMes:", loginMes)
	mes.Data = string(data)
	// 将message序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json marshal 失败:", err)
		return
	}
	// 先把data长度 发送给服务器，

	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	n, err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write fail", err)
		return
	}
	fmt.Printf("客户端，消息长度发送成功 = %d\n", len(data))
	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write fail", err)
		return
	}
	mes1, err := readPkg(conn)
	if err != nil {
		fmt.Println("readPkg err", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes1.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("成功")
	} else {
		fmt.Println(loginResMes.Error)
	}
	// 休眠
	return
}
