package main

import (
	"fmt"
	"os"
)

/*
go的main函数不支持传入参数，也没有返回值，程序通过os.Args 获取命令行参数
*/

func main() {
	fmt.Println(os.Args)
}
