class Solution(object):
    def combinationSum(self, candidates, target):
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
            self.dfs(candidates, target - candidates[i], i, value_list + [candidates[i]])

if __name__ == '__main__':
    print(Solution().combinationSum([8, 7, 4, 3], 11))