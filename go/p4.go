package main

// 转化为求第k最小数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	k := int((m + n) / 2)
	if (m+n)&0x01 == 1 {
		return float64(findKthElement(nums1, m, nums2, n, k+1))
	} else {
		return (float64(findKthElement(nums1, m, nums2, n, k)) + float64(findKthElement(nums1, m, nums2, n, k+1))) / 2
	}
}

func findKthElement(nums1 []int, m int, nums2 []int, n int, k int) int {
	if m > n {
		return findKthElement(nums2, n, nums1, m, k)
	}
	if m == 0 {
		return nums2[k-1]
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}

	pa := min(int(k/2), m)
	pb := k - pa
	if nums1[pa-1] < nums2[pb-1] {
		return findKthElement(nums1[pa:m], m-pa, nums2, n, k-pa)
	} else if nums1[pa-1] > nums2[pb-1] {
		return findKthElement(nums1, m, nums2[pb:n], n-pb, k-pb)
	} else {
		return nums1[pa-1]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	result := findMedianSortedArrays([]int{1, 2, 3, 10}, []int{4, 5, 6})
	println(result)
}
