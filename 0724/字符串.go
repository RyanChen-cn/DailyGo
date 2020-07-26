package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
字符串的本质：存的utf-8 英文占一个字节，中文占三个
1.转换
	string -> []byte  byteSet := []byte(name)
	[]byte -> string  str := string(byteSet)

	[]rune <-> string 与上面一样
	rune存的是Unicode版本 都是占的4个字节

	string(65) -->A 里面接收数字会查utf8码表
	utf8.DecodeRuneInString("A") --> "65" 如果超过两个就只转换第一个


2.长度
	len() -->字节长度
	utf8.RuneCountInString(name) -->字符串长度

3.是否以xxx开头
	res := strings.HasPrefix(name,"g")
4.是否以xxx结尾
	res := strings.HasSuffix(name,"g")
5.是否包含xxx
	res := strings.Contains(name,"o")
6.大小写，不会改变输入的字符串
	res := strings.ToUpper(name)
	res := strings.ToLower(name)
7.去除左右边
	去除左边：res := strings.TrimLeft(name,"go")
	去除右边: res := strings.TrimRight(name,"go")
	去除两边: res := strings.Trim(name,"go")
8.替换
	res := strings.Replace(name,"go","GOOO",2) 从左到右找前两个go替换为GOOO，-1为替换所有
9.分割
	res := strings.Split(name,"的")
10.拼接
	1.性能差	res := "go" + "daily"
	2.性能好一些 	dataList:=[]string{name,lang}
				result := strings.Join(dataList,"")

	3.性能最好 	var builder strings.Builder
				builder.WriteString(name)
				builder.WriteString(lang)
				fmt.Println(builder.String())

 */

func main() {
	var name string ="go语言"
	var lang string = "!"
	var builder strings.Builder
	builder.WriteString(name)
	builder.WriteString(lang)
	fmt.Println(builder.String())
	fmt.Println(utf8.DecodeRuneInString("AAb"))
	fmt.Println(string(27496))
}
