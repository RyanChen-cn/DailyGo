package main

import "fmt"

func main() {
	//goto这东西怎么还秽土转生了。。。。。。
	var numbers = 10
flag:
	if numbers > 5 {
		fmt.Println(numbers)
		numbers--
		goto flag
	}
	//goto语句跳转之间不可以有变量声明
	var numbers2 = 10
flag2:
	if numbers > 5 {
		fmt.Println(numbers2)
		numbers--
		i := 10
		fmt.Println(i)
		goto flag2
	}

}
