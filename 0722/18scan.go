package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
fmt.Scan(地址,地址2。。。。) return 输入的数量、错误信息(不符合数据格式) 必须输入指定数量的值，否则会等待
fmt.Scanln() 跟上面一样，但是读到回车就会直接终止，不会等待输入足够变量
fmt.Scanf() 模板输入 需要按照模板输入，目前不太好用  有相同的return
问题：如果输入字符串带空格怎么处理：使用系统输入
 */

func main() {
	var name string
	var age int
	count,err:=fmt.Scan(&name,&age)
	fmt.Println(name,age,count,err)
	//fmt.Scanf("我叫%s 今年 %d 岁",&name,&age)
	//fmt.Println(name, age)

	//解决空格输入,reader默认一次读取4096个字节，返回的line是字节流 isPrefix：4096个字节一次能读完返回false，否则为trud（这样就可以继续读）
	reader:=bufio.NewReader(os.Stdin)
	line,isPrefix,err:= reader.ReadLine()
	fmt.Println(string(line),isPrefix,err)

}
