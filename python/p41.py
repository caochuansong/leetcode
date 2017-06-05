class Solution(object):
    def firstMissingPositive(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        i = 0
        n = len(nums)
        while i < n:
            if nums[i] != i + 1 and 0 < nums[i] <= n and nums[i] != nums[nums[i] - 1]:
                nums[nums[i] - 1], nums[i] = nums[i], nums[nums[i] - 1]
            else:
                i += 1

        for i in range(n):
            if nums[i] != i + 1:
                return i + 1

        return n + 1


if __name__ == '__main__':
    print(Solution().firstMissingPositive([2, 1]))
