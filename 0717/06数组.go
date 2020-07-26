package main

import "fmt"

func main() {
	//数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。
	//1.确定大小类型
	var array [10]int
	array[0] = 1
	fmt.Println(array)
	//2.指定了array[4]=1其他的是初始化的值
	array2 := [...]int{4: 1}
	fmt.Println(array2)
	//3.大小推导
	var balance3 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance3)
	//4.多维数组
	var threedim [5][10][4]int
	fmt.Println(threedim)

	/*
	声明指针类型的数组不会开辟初始化数组的只，numbers = nil
	var numbers *[3]int
	//初始化
	numbers = new([3]int)
	 */

	/*
	数组的内存管理
		1.数组的内存是连续的
		2.数组的内存地址实际上就是数组第一个元素的内存地址
		3.int类型是直接存在列表中的，占多少自己看是int多少
		4.字符串列表不存储字符串本身，存储字符串指针还有字符 串的len（字节长度） 这个指针和长度要用16个字节来存储
			注意：字符串在go中存储的方式就是用一个指针还有一个长度来进行存储

	 */

	/*
	可变和拷贝
	可变：数组的元素可以被更改，数组的长度、类型不可变
	拷贝:变量赋值时会重新拷贝一份，修改name2不会影响到name1
		name1 := [2]string{"daily","go"}
		name2 := name1

	 */
	/*
	获取数组长度
		len(array)
	切片
		data := array[0:2]
	 */
}
