package main

func main() {
	//常量是一个简单的标识符，在程序的运行时不会修改，
	//常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
	const NAME = "GO"
	const AGE int= 18
	//用作枚举
	const (
		Unknown = 0
		Female = 1
		Male = 2
	)
	//iota是一种特殊常量，iota，特殊常量，可以认为是一个可以被编译器修改的常量。
	//iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
	//iota 可以被用作枚举值：
	const (
		q = iota
		w
		r
	)
	println(q,w,r)
	const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
	)

}