package main

import "fmt"

func main() {
	//字符串的索引还是按照占用字节数来进行取用的 截取方式与python类似
	var name string = "DailyGO你好"
	fmt.Println(name[0:10])

	//这种方式是按照字节来输出的
	for i := 0;i<len(name);i++{
		fmt.Println(i,name[i])
	}

	//这种方式是按照字符串长度输出的,能够识别最后两个是占三个字节的中文
	for index,item := range(name){
		fmt.Println(index,item)
	}
	fmt.Println("--------------------整体转为Unicode编码----------------------------")
	//转换编码后 两种做法就等价了
	dataList :=[]rune(name)
	//这种方式是按照字节来输出的
	for i := 0;i<len(dataList);i++{
		fmt.Println(i,dataList[i])
	}

	//这种方式是按照字符串长度输出的,能够识别最后两个是占三个字节的中文
	for index,item := range(dataList){
		fmt.Println(index,item)
	}
}
