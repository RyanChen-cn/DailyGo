package main

func main() {
/*
	//1.光声明map还是不能用，需要用make初始化
	var map_var map[string]string
	map_var = make(map[string]string)
	map_var2 := make(map[string]string,10) //避免后续扩张时性能问题,可以提前申请内存
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
 */
/*
	套娃的map

	//make初始化第一层map，分配内存
	stuMap := make(map[string]map[string]string)
	//第二层map初始化
	stuMap["stu01"] = make(map[string]string)
	stuMap["stu01"]["名字"] = "大狗子"
	stuMap["stu01"]["年纪"] = "18"
	//切记，map必须make后方可使用
	stuMap["stu02"] = make(map[string]string)
	stuMap["stu02"]["名字"] = "二麻子"
	stuMap["stu02"]["年纪"] = "17"
	fmt.Println(stuMap)
	//取出所有学生的信息
	for k, v := range stuMap {
		fmt.Printf("k值是学生：%v  v值是学生信息：%v\n", k, v)
		//k1是键，v1是值
		for k1, v1 := range v {
			fmt.Printf("\tk1：%v v1：%v\n", k1, v1)
		}
		fmt.Println()
	}

 */
/*
	map 切片？ 或者叫map的动态数组

	//声明map切片，
	// 默认值        [map[] map[] map[] map[] map[]]
	sliceMap := make([]map[string]string, 5)
	for i := 0; i < 5; i++ {
		//map必须初始化再用,遍历初始化
		sliceMap[i] = make(map[string]string)
	}
	sliceMap[0]["名字"] = "张飞"
	sliceMap[1]["性别"] = "不男不女"
	sliceMap[2]["体重"] = "三百斤"
	fmt.Println(sliceMap)
	fmt.Printf("容量：%v，长度：%v\n", cap(sliceMap), len(sliceMap))
	//动态扩容map切片，用append函数 触发扩容 cap翻倍
	newSliceMap := map[string]string{
		"姓名": "狗子",
		"爱好": "吃包子",
	}
	//todo 测试扩容机制
	//append函数进行切片扩容，动态增加
	sliceMap = append(sliceMap, newSliceMap)
	fmt.Println(sliceMap)
	fmt.Printf("容量：%v，长度：%v\n", cap(sliceMap), len(sliceMap))
*/

/*
	1.map的排序
	map本身是无序的，想要对map排序输出，就是对key值进行排序 然后按key排序索引
	排序时调用sort包 sort.Strings(slice[:]) 无返回值
	2.map是传引用调用 以map作为参数的函数会修改字典
	3.map自动扩容 机制 //todo
	4.map的value也可以是struct类型，适合更复杂的数据
 */

}

