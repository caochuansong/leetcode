class Solution(object):
    def combinationSum2(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        candidates = sorted(candidates)
        Solution.ret = []
        self.dfs(candidates, target, 0, [])
        return Solution.ret

    def dfs(self, candidates, target, start, value_list):
        length = len(candidates)
        if target == 0:
            Solution.ret.append(value_list)
        for i in range(start, length):
            if target < candidates[i]:
                return
            if i > start and candidates[i] == candidates[i - 1]:
                continue
            self.dfs(candidates, target - candidates[i], i + 1, value_list + [candidates[i]])

if __name__ == '__main__':
    print(Solution().combinationSum2([10, 1, 2, 7, 6, 1, 5], 8))