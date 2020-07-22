package main

import "fmt"

func main() {
	//go语言必须进行强制类型转换
	var i int = 43
	var f float32 = float32(i)
	fmt.Println(f)
	/*
	在函数外面声明的是全局变量
	var country string = "china"
	或可使用 var country = "china"
	不可使用 country := "china"

	类似函数的情况，首字母小写的全局变量只能在本包内全局有效，首字母大写的全局变量代表共有 是全局有效的
	 */
}
