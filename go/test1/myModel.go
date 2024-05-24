package test1

import (
	"fmt"
)

func init() {
	fmt.Println("导入resources包时自动执行")
}

func CusPrint() {
	fmt.Println("my custom print")
}
