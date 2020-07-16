package main

import "fmt"

func main() {
	//go语言必须进行强制类型转换
	var i int = 43
	var f float32 = float32(i)
	fmt.Println(f)
}
