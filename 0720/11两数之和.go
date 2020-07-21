package main

import "fmt"

func main() {
	nums := []int{3, 3}
	target := 6
	fmt.Println(twoSumMap(nums, target))
}
func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}

			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
func twoSumMap(nums []int, target int) []int {
	m := make(map[int]int)
	//之所以不把字典建完了再扫描，是因为可能出现重复的key，所以就先判断再放进去
	for pos, k := range nums {
		if pos2, has := m[target-k]; has {
			return []int{pos, pos2}
		}

		m[k] = pos
	}

	return []int{}
}
