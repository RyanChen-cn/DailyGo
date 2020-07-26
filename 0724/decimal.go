package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

/*
优秀的第三方包，用来更精准的处理浮点数

如何安装第三方包？
1. go get github.com/shopspring/decimal 会下载到 GOPATH/src 还有pkg
使用第三方包

 */
func main() {
	var v1 = decimal.NewFromFloat(0.000001)
	var v2 = decimal.NewFromFloat(0.286)
	fmt.Println(v1.Add(v2))
	fmt.Println(v1.Sub(v2))
	fmt.Println(v1.Mul(v2))
	fmt.Println(v1.Div(v2))
	fmt.Println(v2.Round(2)) //四舍五入
	fmt.Println(v2.Truncate(2)) //直接舍弃


}