package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	resultList := make([][]int, 0)
	dfs(candidates, target, 0, make([]int, 0), &resultList)
	return resultList
}

func dfs(candidates []int, target int, start int, valueList []int, resultList *[][]int) {
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
		dfs(candidates, target-candidates[i], i, append(valueList, candidates[i]), resultList)
	}
}

func main() {
	resultList := combinationSum([]int{2, 3, 7}, 18)
	for _, s := range resultList {
		for _, v := range s {
			print(v)
		}
		println()
	}
}
