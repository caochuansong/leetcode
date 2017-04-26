package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		//去重，因为上一个数i已经把i开头的所有可能都算过了，所有下一个数如果相同，没必要再计算了
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		r := twoSums(nums[i], nums[i+1:], 0-nums[i])
		if len(r) > 0 {
			result = append(result, r...)
		}
	}
	return result
}

func twoSums(v int, nums []int, target int) [][]int {
	i := 0
	j := len(nums) - 1
	var r [][]int
	for i < j {
		if nums[i]+nums[j] == target {
			r = append(r, []int{v, nums[i], nums[j]})
			i++
			j--
			for nums[i] == nums[i-1] && i < j {
				i++
			}
			for nums[j] == nums[j+1] && i < j {
				j--
			}
		} else if nums[i]+nums[j] < target {
			i++
			for nums[i] == nums[i-1] && i < j {
				i++
			}
		} else {
			j--
			for nums[j] == nums[j+1] && i < j {
				j--
			}
		}

	}
	return r
}

func main() {
	a := []int{-1, -1, 0, 1, 2}
	r := threeSum(a)
	for _, m := range r {
		for _, n := range m {
			println(n)
		}
	}
}
