package main

import "fmt"

func main() {
	//defer延迟调用
	//defer后面接一个函数，这个函数的调用会推迟到当前函数执行完再执行
	//是一个比较特殊的关键字
	//多个defer越先写的越后调用
	//这个方法常用于最后释放资源用
	defer fmt.Println("main over")
	fmt.Println("main working")

}
