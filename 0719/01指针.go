package main

import "fmt"

func main() {
	//1新建指针
	//1.1指向已有变量的地址
	var num int = 12
	var ptr  = &num
	fmt.Println(*ptr)
	//可以通过指针修改指向的值
	*ptr = 22
	fmt.Println(*ptr)
	//1.2创建指针，分配内存，然后再把值放到指向的地址
	var ptr1 = new(int)
	*ptr1 = 66
	//1.3声明变量，之后再找个变量的地址赋过来
	var ptr2 *int
	ptr2 = &num
	fmt.Println(*ptr2)
	//第一种第三种建出来的指针是一样的
	fmt.Println(ptr==ptr2)
	//2指针的类型
	astr := "hello"
	aint := 1
	abool := false
	arune := '☞'


	fmt.Printf("astr 指针类型是：%T\n", &astr)
	fmt.Printf("aint 指针类型是：%T\n", &aint)
	fmt.Printf("abool 指针类型是：%T\n", &abool)
	fmt.Printf("arune 指针类型是：%T\n", &arune)
	//3.指针的零值 为nil
	//指针判空
	if ptr == nil{

	}
	//4.指针在函数传递时的用法，注意切片也是用的指针
	var array = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	modify(array[:])
	fmt.Println(array)
	modify2(&array)
	fmt.Println(array)
}
//4.1使用切片
func modify(sls []int) {
	sls[0] = 90
}
//4.2 使用指针
func modify2(arr *[10]int) {
	(*arr)[0] = 66
}