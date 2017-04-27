package main

//o(n) 方法  直接顺序遍历
func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target <= nums[i] {
			return i
		}
	}
	return len(nums)
}

//o(lgn) 方法 采用二分法
func searchInsert1(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	j := len(nums) - 1
	for i <= j {
		mid := int((i + j) / 2)
		if nums[mid] > target {
			j = mid - 1
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			return mid
		}
	}
	return i
}

func main() {
	a := []int{3, 4, 6}
	println(searchInsert(a, 0))
	println(searchInsert(a, 3))
	println(searchInsert(a, 4))
	println(searchInsert(a, 5))
	println(searchInsert(a, 6))
	println(searchInsert(a, 7))

	println()

	println(searchInsert1(a, 0))
	println(searchInsert1(a, 3))
	println(searchInsert1(a, 4))
	println(searchInsert1(a, 5))
	println(searchInsert1(a, 6))
	println(searchInsert1(a, 7))
}
