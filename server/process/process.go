package process

import (
	"fmt"
	"imessage/common/message"
	"imessage/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (p *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		up := &UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessLogin(mes)
		// 处理登录
	case message.RegisterMesType:
	// 处理注册
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (p *Processor) Process2() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: p.Conn,
			Buf:  [8096]byte{},
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出了，服务端也应该退出")
				return err
			}
			fmt.Println("readPkg err=", err)
			continue
		}
		err = p.serverProcessMes(&mes)
		if err != nil {
			fmt.Println("err:", err)
			return err
		}
	}
}
