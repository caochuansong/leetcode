class Solution(object):
    # 利用 map来存每个数出现的下标
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        nMap = {}
        for i, num in enumerate(nums):
            v = nMap.get(target - num, -1)
            if v >= 0:
                return [v, i]
            nMap[num] = i
        return []


if __name__ == "__main__":
    s = Solution()
    print(s.twoSum([1, 2, 3, 4], 3))
