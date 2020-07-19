package main

import "fmt"

func main() {
	//1.整形
	//int(看操作系统位数) int8 int16 int 32 int 64
	//uint(看操作系统位数) uint8 uint16 uint32 uint64
	//进制表示
		//2进制 0b 0B开头
		//8进制 0o 0O开头
		//16进制 0x 0X开头
	var num01 int = 0b1100
	//print(num01)
	var num02 int = 0XC
	//println(num02)
	//支持格式化输出
	fmt.Printf("2进制数 %b 表示的是: %d \n", num01, num01)
	fmt.Printf("16进制数 %x 表示的是: %d \n", num02, num02)
	//%b    表示为二进制
	//%c    该值对应的unicode码值
	//%d    表示为十进制
	//%o    表示为八进制
	//%q    该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	//%x    表示为十六进制，使用a-f
	//%X    表示为十六进制，使用A-F
	//%U    表示为Unicode格式：U+1234，等价于"U+%04X"
	//%E    用科学计数法表示
	//%f    用浮点数表示

	//2.浮点型
	//go语言中浮点数类型只能使用10进制表示，可以使用E或e来进行科学表示
	//float32第一位符号位 8位指数位 32位尾数
	//float64第一位符号位 11位指数位 52位尾数
	var num03 float32 = 3.7e10
	fmt.Println(num03)
	var num04 float32 = 3.7E-5
	fmt.Println(num04)
	//3.复数
	var num05 complex64 = 3.7e-20
	var num06 complex128 = -2e-100
	fmt.Println(num05,num06)
	//4.布尔型
	var boolvar bool = true
	fmt.Println(boolvar)
	//go语言的类型审查非常严格 bool不能跟0/1比较
	//fmt.Println(boolvar==1)
}
