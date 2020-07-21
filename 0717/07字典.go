package main

import "fmt"

func main() {
	//1.光声明map还是不能用，需要用make初始化
	var map_var map[string]string
	map_var = make(map[string]string)
	map_var2 := make(map[string]string)
	map_var["xiaoming"] = "爱读书"
	map_var2["laoming"] = "爱吃肉"

	//2.其他初始化方法
	var scores map[string]int = map[string]int{"english": 80, "chinese": 85}
	scores2 := map[string]int{"english": 80, "chinese": 85}
	fmt.Println(scores, scores2)

	//3.查找操作
	//3.1 判断是否存在，has会返回一个布尔值
	var hobby, has = map_var["xiaoming1"]
	println(hobby, has)
	//3.2 不判断是否存在  如果不存在会返回零值
	var hobby2 = map_var["xiaoming1"]
	println(hobby2)

	//4.删除操作
	delete(scores, "math")

	//5.遍历操作
	//5.1获取key-value
	for subject, score := range scores2 {
		fmt.Printf("key: %s, value: %d\n", subject, score)
	}
	//5.2获取key
	for subject := range scores2 {
		fmt.Printf("key: %s\n", subject)
	}
	//5.3获取value
	for _, score := range scores2 {
		fmt.Printf("value: %d\n", score)
	}
}
