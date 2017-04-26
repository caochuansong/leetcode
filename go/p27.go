package main

func removeElement(nums []int, val int) int {
	index := 0
	for _, v := range nums {
		if v != val {
			nums[index] = v
			index++
		}
	}
	return index
}
