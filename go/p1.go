package main

import "fmt"

// 利用map来实现  复杂度o(n)
func twoSum(nums []int, target int) []int {
	nMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if v, ok := nMap[target-nums[i]]; ok {
			return []int{v, i}
		}
		nMap[nums[i]] = i
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
