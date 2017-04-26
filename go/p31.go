package main

func nextPermutation(nums []int) {
	length := len(nums)
	i := length - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := length - 1
		for j > i && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums, i+1, length-1)
}

func reverse(nums []int, left int, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func main() {
	nums := []int{1, 1, 5}
	nextPermutation(nums)
	for _, v := range nums {
		println(v)
	}
}
