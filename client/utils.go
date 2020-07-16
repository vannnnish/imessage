package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"imessage/common/message"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
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
}

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
