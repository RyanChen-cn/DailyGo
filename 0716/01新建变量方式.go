package main

import "fmt"

func GetData() (int, int) {
	return 100, 200
}
func main() {
	//1.var <name> <type> 给了准确的类型就给了初始化比如 string 类型就初始化为空字符串
	//int 类型就初始化为0，float 就初始化为 0.0，bool类型就初始化为false，指针类型就初始化为 nil。
	var number int
	var str string = "hello go"

	fmt.Println(number)
	fmt.Println(str)

	//2.同时定义多个变量
	var (
		name   string
		age    int
		gender string
	)
	fmt.Println(name, age, gender)

	//3.:=推断声明初始化 := 结构不能在函数外使用
	name1 := "hello"
	// 等价于
	var name2 string = "go"
	//4.由初始化赋值推断类型
	var name3 = "!"
	fmt.Println(name1, name2, name3)
	//5.推断声明多个变量
	myname, myage := "ryan", 18
	fmt.Println(myname, myage)

	//6.new函数声明指针变量
	var ptr = &myname
	fmt.Println(ptr)
	//*指针变量 获取指向的值
	fmt.Println(*ptr)
	//新建一个int类型的匿名变量初始化规则同上，ptr2保留这个变量的地址
	ptr2 := new(int)
	fmt.Println(ptr2)
	//*指针变量 获取指向的值
	fmt.Println(*ptr2)
	//可以通过指针改变变量值
	*ptr2 = 10
	fmt.Println(ptr2)
	fmt.Println(*ptr2)

	//7.匿名变量/占位符：不分配内存、可重复声明
	a, _ := GetData()
	_, b := GetData()
	fmt.Println(a, b)
}
