package main

import (
	"packagestudy/core"
	"packagestudy/middleware"
)

var (
	name string
	nick string
)

// https://blog.csdn.net/baidu_33256174/article/details/136520102 【go语言开发】gorm库连接和操作mysql，实现一个简单的用户注册和登录
func main() {

	// var a int = 10
	// var ip *int
	// ip = &a
	// fmt.Println("ip:%d", *ip)

	// str := "123"
	// num, err := strconv.Atoi(str)
	// if err != nil {
	// 	fmt.Println("转换错误:", err)
	// } else {
	// 	fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str, num)
	// }

	core.InitMysql()
	middleware.InitRouter()
	// test1.CusPrint()
}

func add() (int, int) {
	a, b := 1, 3
	return a, b
}

// 一个可以返回多个值的函数
func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}
