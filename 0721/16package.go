package main
// go包管理，明确项目中文件间组织关系
// 一个文件夹可以称为一个包 包名可以与文件夹名不同
// 在文件夹（包）中可以创建多个文件
// 在同一个包中的文件必须指定相同的包名称
// 包分为main包和非main包，main包里面必须要写一个main函数作为项目的入口
// 调本包直接写函数名  掉其他包需要 包名.函数名 函数名大写认为是外部包可调用的方法  函数名小写是包内才可以用的

import "fmt"

func main() {
	fmt.Println(add(10,12))
}
