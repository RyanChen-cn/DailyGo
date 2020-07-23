package main

import "fmt"
/*
 fmt.Sprint:字符串格式化组合，比用加号加起来更好
 */
func main() {
	var name,address,action string
	fmt.Scanln(&name, &address,&action)
	result := fmt.Sprintf("我叫%s，在%s正在干%s",name, address, action)
	fmt.Println(result)

}
