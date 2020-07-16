package main

import "fmt"

var (
	userId int
	pwd    string
)

func main() {
	var key int
	var loop = true
	for loop {
		fmt.Println("--欢迎登陆多人聊天系统--")
		fmt.Println("1. 登陆聊天系统")
		fmt.Println("2. 注册")
		fmt.Println("3. 退出系统")
		fmt.Println("3. 请选择1-3")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出聊天室")
			loop = false
		default:
			fmt.Println("输入的数字有误")

		}
	}
	if key == 1 {
		// 说明要登录
		fmt.Println("请输入id:")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入密码:")
		fmt.Scanf("%s\n", &pwd)
		// 登录
		err := login(userId, pwd)
		if err != nil {
			fmt.Println("登录失败")
		} else {
			fmt.Println("登录成功")
		}
	} else if key == 2 {
		fmt.Println("进行用户注册")
	}
}
