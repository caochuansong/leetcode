package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var result [][]int
	length := len(nums)
	for i := 0; i < length-3; i++ {
		//去重，因为上一个数i已经把i开头的所有可能都算过了，所有下一个数如果相同，没必要再计算了
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < length-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			sum := target - nums[i] - nums[j]
			left, right := j+1, length-1
			for left < right {
				if nums[left]+nums[right] == sum {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
					for nums[left] == nums[left-1] && left < right {
						left++
					}
					for nums[right] == nums[right+1] && left < right {
						right--
					}
				} else if nums[left]+nums[right] < sum {
					left++
					for nums[left] == nums[left-1] && left < right {
						left++
					}
				} else {
					right--
					for nums[right] == nums[right+1] && left < right {
						right--
					}
				}
			}
		}
	}
	return result
}

func main() {
	fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
}
