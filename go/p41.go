package main

func firstMissingPositive(nums []int) int {
	i := 0
	n := len(nums)
	for i < n {
		if nums[i] != i+1 && nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}

	for i = 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func main() {
	println(firstMissingPositive([]int{1, 2, 4, 4, 5}))
}
