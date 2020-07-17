package main

import "fmt"

func main() {
	//数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。
	//1.确定大小类型
	var array [10] int
	array[0] = 1
	fmt.Println(array)
	//指定了array[4]=1其他的是初始化的值
	array2 := [...]int{4:1}
	fmt.Println(array2)
	//2.大小类型推导
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	//切片
	var balance2 = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	var balance3 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance,balance2,balance3[4])
	//3.不定长,切片 也有不同点，=切片本身是引用类型，它更像是 Python 中的 list ，我们可以对它 append 进行元素的添加。
	var balance4 []int
	balance4=append(balance4, 10)
	fmt.Println(balance4,cap(balance4))
	//4.多维数组
	var threedim [5][10][4]int
	fmt.Println(threedim)
	//5.切片
	var slice = balance[0:3]
	fmt.Println(slice)
	//make( []Type, size, cap )
	a := make([]int, 2)
	b := make([]int, 2, 10)
	fmt.Println(a, b)
	fmt.Println(len(a), len(b))
	fmt.Println(cap(a), cap(b))


	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[0:2]
	fmt.Printf("myslice为 %d, 其长度为: %d cap %d\n", myslice, len(myslice),cap(myslice))
	//因为切片是对数组的引用，所以会改变原来的值
	myslice[0] =99
	fmt.Println(myslice[0],numbers4[0])


}