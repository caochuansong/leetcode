package main

func searchRange(nums []int, target int) []int {
	var res []int = []int{-1, -1}
	ll, lr := 0, len(nums)-1
	for ll <= lr {
		m := int((ll + lr) / 2)
		if nums[m] < target {
			ll = m + 1
		} else {
			lr = m - 1
		}
	}
	rl, rr := 0, len(nums)-1
	for rl <= rr {
		m := int((rl + rr) / 2)
		if nums[m] <= target {
			rl = m + 1
		} else {
			rr = m - 1
		}
	}
	if ll <= rr {
		res[0] = ll
		res[1] = rr
	}
	return res
}

func main() {
	nums := []int{2, 7, 11, 15}
	searchRange(nums, 7)
}
