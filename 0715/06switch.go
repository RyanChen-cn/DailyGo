package main

import (
	"fmt"
	"time"
)

func main() {
	i:=2
	fmt.Println(i)
	switch i {
	case 1:
		fmt.Println('1')
	case 2:
		fmt.Println("2")
	}
	fmt.Println(time.Now().Weekday())


	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	var a int
	print(a)
}

