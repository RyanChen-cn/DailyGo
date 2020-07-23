package main

import "fmt"

/*
声明变量
	方式：声明不赋值/声明赋值
		var v1 int --->在内存中存了0
		v2:=999    --->在内存中存了999

指针
	声明方式
		var v3 *int 新建1个空间 1个为v4本身，值为另nil
		var v4 = new(int) 新建两个空间 1个为v4本身，值为另一个空间的地址；另一个空间的值会默认存为数据默认值
							这个操作会建立一个指针，还有指针对应的值
nil:指go语言中的空值
 */

func main() {
	var v *int
	var vv *float32
	var v2 = new(int)
	fmt.Println(v)
	fmt.Println(vv)
	fmt.Println(v2)
}
