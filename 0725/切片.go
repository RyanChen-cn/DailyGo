package main

import "fmt"

func main() {
	/*
		切片有点类似动态数组

			type slice struct {
				array unsafe.Pointer 指针：指向数据地址
				len int 当前长度
				cap int 容量
			}
		在向切片中插入数据时容量不够用时，内部会自动扩容为当前的2倍  如果len超过1024，扩容只增加1/4


	*/
	/*
		创建方式

		//1.只声明
		var nums []int
		//2.声明并赋值
		var data = []int{11,22,33}
		//3.通过make初始化 (类型，长度，容量) (类型，长度=容量)
		//make只可以初始化 切片、字典、channel使用
		var users = make([]int,2,5)
		var users = make([]int,2)

			//4.new创建切片指针
		var ptr = new([]int)
	*/

	/*
		自动扩容
		v2 := append(v1,"99")
		注意：1.这段代码的本质是创建了一个新的v2，v2的len比原来大1，数据存储的地址指向的是同一个，v1，v2对[0]具有共同的控制权，
				v1因为len的限制不能访问[1]
			 2.如果在append过程中cap如果需要扩充也会自动扩充，会生成一个新的数据存储地址，不会再进行共用了
			v1 := make([]string,1,3)
			fmt.Println(v1,len(v1),cap(v1))
			v2 := append(v1,"99")
			fmt.Println(v2,len(v2),cap(v2))
			v2[0] = "66"
			fmt.Println(v1[1],v2)
	*/

	/*
		常见操作
		1.len、cap
		2.索引 ---》索引最大值是看len的，不是看cap的
		3.切片 ---》v1[1:3] [1:] [:3] 返回值数据是同一块内存地址，可能指针指向的不是头了往后指了几位
		4.追加：v1 = append(v1,55,66,77) 原理见上面
		       v1 = append(v1,v1...) 两个加一块 原理见上面
		5.删除：v1 = append(v1[:2],v1[3:]...) 删除[2]
		6.插入：跟删除类似，利用切片思路
		7.嵌套：切片里面套切片：v2 := [][]int{[]int{1,2},[]int{2,3,4}}
		       切片里面放数组：v3 := [][2]int{[2]int{1,2},{1,2}}
	*/

	v2 := [][]int{[]int{1, 2}, []int{2, 3, 4}}
	v3 := [][2]int{[2]int{1, 2}, {1, 2}}
	fmt.Println(v2, v3)
}
