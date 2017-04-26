package main

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	minDiff := result - target
	if minDiff < 0 {
		minDiff = -minDiff
	}
	for i := 0; i < len(nums)-2; i++ {
		//去重，因为上一个数i已经把i开头的所有可能都算过了，所有下一个数如果相同，没必要再计算了
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		begin := i + 1
		end := len(nums) - 1
		for begin < end {
			sum := nums[begin] + nums[end] + nums[i]
			if sum == target {
				return target
			} else if sum < target {
				diff := target - sum
				if diff < minDiff {
					minDiff = diff
					result = sum
				}
				begin++
				for nums[begin] == nums[begin-1] && begin < end {
					begin++
				}
			} else {
				diff := sum - target
				if diff < minDiff {
					minDiff = diff
					result = sum
				}
				end--
				for nums[end] == nums[end+1] && begin < end {
					end--
				}
			}

		}
	}
	return result
}
func main() {
	println(threeSumClosest([]int{0, 2, 1, -3}, 1))
}
