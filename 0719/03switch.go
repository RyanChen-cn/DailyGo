package main

import "fmt"

func main() {
	//go语言中默认满足一个case就会退出 都不满足就会执行default
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90: grade = "A"
	case 80: grade = "B"
	//可以case多个值
	case 50,60,70 : grade = "C"
	default: grade = "D"
	}

	switch {
	case grade == "A" :
		fmt.Printf("优秀!\n" )
	case grade == "B", grade == "C" :
		fmt.Printf("良好\n" )
	case grade == "D" :
		fmt.Printf("及格\n" )
	case grade == "F":
		fmt.Printf("不及格\n" )
	default:
		fmt.Printf("差\n" );
	}
	fmt.Printf("你的等级是 %s\n", grade );
	//开启c语言那种穿透执行
	//switch中可以空着
	s := "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		//只能穿透一层
		fallthrough
	case s != "world":
		fmt.Println("world")
	}
}