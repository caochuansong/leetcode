package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	currentIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[currentIndex] {
			nums[currentIndex + 1] = nums[i]
			currentIndex++
		}
	}
	return currentIndex + 1
}

func main() {
	println(removeDuplicates([]int{10, 9, 2, 1, 1}))
}
