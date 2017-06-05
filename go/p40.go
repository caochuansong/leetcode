package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	resultList := make([][]int, 0)
	dfs2(candidates, target, 0, make([]int, 0), &resultList)
	return resultList
}

func dfs2(candidates []int, target int, start int, valueList []int, resultList *[][]int) {
	length := len(candidates)
	if target == 0 {
		tmpCurComb := make([]int, len(valueList))
		copy(tmpCurComb, valueList)
		*resultList = append(*resultList, tmpCurComb)
		return
	}
	for i := start; i < length; i++ {
		if target < candidates[i] {
			break
		}
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		dfs2(candidates, target-candidates[i], i+1, append(valueList, candidates[i]), resultList)
	}
}

func main() {
	resultList := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	for _, s := range resultList {
		for _, v := range s {
			print(v)
		}
		println()
	}
}
