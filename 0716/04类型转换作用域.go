package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//go语言必须进行强制类型转换
	var i int = 43
	var f float32 = float32(i)
	fmt.Println(f)
	/*
	整形与字符串互转
	strconv.Itoa(v) 里面只能放int，不是int可以转一下
	strconv.Atoi(result) 因为可能转换失败，所以会返回两个值 一个str 一个错误信息,如果转化顺利 err=nil，本质调用了strconv.ParseInt

	拓展：查看变量类型
	reflect.TypeOf(result)
	 */
	v :=19
	result := strconv.Itoa(v)
	fmt.Println(result,reflect.TypeOf(result))
	strtoint, err :=strconv.Atoi(result)
	fmt.Println(strtoint,err,reflect.TypeOf(strtoint))

	/*
	进制转化
	十进制转到其他
		strconv.FormatInt(int64(v1),16) 返回的是字符串
	其他进制字符转int
		r2,err2:=strconv.ParseInt(r1,字符串表示的进制,存的是int多少的约束) 填的8就是超过int8就报错 如果存的下返回的结果 永远 是int64
	 */
	v1 := 17
	r1 := strconv.FormatInt(int64(v1),2)
	fmt.Println(r1,reflect.TypeOf(r1))

	r2,err2:=strconv.ParseInt(r1,2,8)
	fmt.Println(r2, err2)

	/*
	字符串与布尔类型转换
	string->bool
		result,err := strconv.ParseBool("true")
		true:"1","t","T","true","TRUE","True"
		false:"0","f","F","false","FALSE","False"
		如果不是上面的几个关键字，则会返回false，同时err也会出现错误信息
	bool->string
		bo = strconv.FormatBool(false)

	*/


	/*
	在函数外面声明的是全局变量
	var country string = "china"
	或可使用 var country = "china"
	不可使用 country := "china"

	类似函数的情况，首字母小写的全局变量只能在本包内全局有效，首字母大写的全局变量代表共有 是全局有效的
	 */
}
