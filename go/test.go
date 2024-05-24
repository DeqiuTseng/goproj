package main

import (
	"fmt"
	"packagestudy/test1"
	"strconv"
)

var (
	name string
	nick string
)

func main() {
	var a int = 10
	var ip *int
	ip = &a
	fmt.Println("ip:%d", *ip)

	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str, num)
	}

	test1.CusPrint()
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
