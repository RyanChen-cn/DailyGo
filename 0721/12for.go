package main

import "fmt"

func main() {
	//1.for当while用
	var num = 3
	for num > 1 {
		fmt.Println(num)
		num--
	}
	//out : 3 2

	//2.for当for用
	for i := 1; i < 3; i++ {
		fmt.Println(i)
	}
	//out: 1 2

	//3.for死循环
	for {
		fmt.Println(num)
		break
	}
	for {
		fmt.Println(num)
		break
	}

	//4.for-range遍历可迭代对象 在map那里用过
	var strarry []string = []string{"hello", "go"}
	for key, str := range strarry {
		fmt.Println(key, str)
	}
	//out 0 hello
	//	  1 go

	//注：如果只接受一个变量，接受的是索引变量
	for key := range strarry {
		fmt.Println(key)
	}
	//out 0 1

}
