package main

import (
	"fmt"
	"imessage/server/process"
	"net"
)

func main() {
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.listen err=", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("监听失败err:", err.Error())
			continue
		}
		go dealConn(conn)

	}
}

func dealConn(conn net.Conn) {
	defer conn.Close()
	// 读取客户端的消息
	p := process.Processor{Conn: conn}
	p.Process2()
}

/*func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("等待读取客户端发送的数据:")
	// conn.Read 在 conn没有被关闭的情况下，才会注册，如果客户端关闭了conn, 那么就不会阻塞了， 因此，当客户端关闭连接口，服务端也应该关闭这个链接
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("read fail ", err)
		return
	}
	// 根据将buf转成uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	// 根据pkgLen读取消息
	_, err = conn.Read(buf[:pkgLen])
	if err != nil {
		fmt.Println("conn.Read fail err", err)
		return
	}
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("conn.Real fail", err)
		return
	}
	return
}*/

/*func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		err = serverProcessLogin(conn, mes)
		// 处理登录
	case message.RegisterMesType:
	// 处理注册
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}*/

/*func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
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
	err = writePkg(conn, data)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	return
}*/
/*
func writePkg(conn net.Conn, data []byte) (err error) {
	// 先发送一个长度给客户端
	// 根据将buf转成uint32类型
	var pkgLen uint32
	var buf [4]byte
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 根据pkgLen读取消息
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Read fail err", err)
		return
	}
	// 发送data本身
	n, err = conn.Write(data)
	if n != 4 || err != nil {
		fmt.Println("conn.Read fail err", err)
		return
	}
	return
}
*/
