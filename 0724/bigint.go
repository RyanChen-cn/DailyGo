package main

import (
	"fmt"
	"math/big"
)
/*
创建
	v1 v3都对初值进行了初始化
	v2 没有初始化时nil
写入
	v1.SetInt64(19996)
	第二个参数是输入的进制
	v1.SetString("89465123463212132",10)
输出
	v1.String() string输出
	v1.Int64() int64输出
操作
	result = new(big.Int)
	result.Add(n1,n2)
	Sub
	Mul
	Div 只返回商
	DivMod(v1,v2,minder) minder存余数
 */
/*
new(返回指针)和make（返回类型）
new和make就是为引用类型的变量分配内存空间的，但是对于本来就是引用类型的slice，chan和map来说，创建后它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了，此时就用make来创建

二者异同
二者都是内存的分配（堆上），但是make只用于slice、map以及channel的初始化（非零值）；而new用于类型的内存分配，并且内存置为零。所以在我们编写程序的时候，就可以根据自己的需要很好的选择了。

make返回的还是这三个引用类型本身；而new返回的是指向类型的指针。
 */
func main() {

	//var v1 big.Int
	//var v2 *big.Int
	var v3 = new(big.Int)
	v3.SetString("123546",10)
	fmt.Println(v3)
	//v1.SetInt64(19996)
	//v1.SetString("89465123463212132",10)


	//简化
	//var v4 = big.NewInt(66)
}
