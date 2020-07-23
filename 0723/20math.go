package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Abs(-19))
	fmt.Println(math.Floor(3.14)) //下取整
	fmt.Println(math.Ceil(3.14))	//上取整
	fmt.Println(math.Round(3.49)) //四舍五入
	fmt.Println(math.Round(3.1415926*100)/100) //保留两位小数的四舍五入
	fmt.Println(math.Mod(10,3)) //模运算
	fmt.Println(math.Pow(2,5))
	fmt.Println(math.Pow10(2))
	fmt.Println(math.Max(2,3))
	fmt.Println(math.Min(2,3))


}
