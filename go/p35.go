package main

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target <= nums[i] {
			return i
		}
	}
	return len(nums)
}

func main() {
	a := []int{3, 4, 6}
	println(searchInsert(a, 0))
	println(searchInsert(a, 3))
	println(searchInsert(a, 4))
	println(searchInsert(a, 5))
	println(searchInsert(a, 6))
	println(searchInsert(a, 7))
}
