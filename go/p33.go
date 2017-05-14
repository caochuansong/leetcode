package main
//如果我们不寻找那个元素先，而是直接相等的时候也向一个方向继续夹逼，如果向右夹逼，
// 最后就会停在右边界，而向左夹逼则会停在左边界，如此用停下来的两个边界就可以知道结果了，只需要两次二分查找
func search(nums []int, target int) int {
	var l, h = 0, len(nums)-1
	for l <= h {
		m := l + ((h - l) >> 1)
		if nums[m] == target {
			return m
		}
		if nums[m] < nums[h] {
			if nums[m] < target && target <= nums[h] {
				l = m + 1
			} else {
				h = m - 1
			}
		} else {
			if nums[l] <= target && target < nums[m] {
				h = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return -1
}
