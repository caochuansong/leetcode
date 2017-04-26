package main

func maxArea(height []int) int {
	var area int = 0
	i := 0
	j := len(height) - 1
	for i < j {
		area = max(area, min(height[i], height[j])*(j-i))
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	println(maxArea([]int{1, 2, 5, 6}))
}
